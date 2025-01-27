package task

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"autograder/pkg/model/constants"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
	"autograder/pkg/dao"
	"autograder/pkg/dao/docker"
	"autograder/pkg/model/assembler"
	"autograder/pkg/model/dbm"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"
)

type ServiceImpl struct {
	groupDAO    *dao.GroupDAO
	workerCh    chan *entity.AppInfo
	userTaskSet *hashset.Set
	rwLock      sync.RWMutex
}

func NewService(groupDAO *dao.GroupDAO) *ServiceImpl {
	svc := &ServiceImpl{
		groupDAO:    groupDAO,
		workerCh:    make(chan *entity.AppInfo, 200),
		userTaskSet: hashset.New(),
	}
	go svc.worker(context.Background())
	return svc
}

func (s *ServiceImpl) putUserTask(userID uint) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	s.userTaskSet.Add(userID)
}

func (s *ServiceImpl) removeUserTask(userID uint) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	s.userTaskSet.Remove(userID)
}

func (s *ServiceImpl) existsUserTask(userID uint) bool {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	return s.userTaskSet.Contains(userID)
}

func (s *ServiceImpl) SubmitApp(ctx context.Context, info *entity.AppInfo) (entity.SubmitAppResult, error) {
	if s.existsUserTask(info.User.UserID) {
		return entity.SubmitAppResultSystemTaskExists, nil
	}

	isFull := utils.SendIfNotFull(s.workerCh, info)
	if isFull {
		return entity.SubmitAppResultSystemBusy, nil
	}

	s.putUserTask(info.User.UserID)

	model := info.ToDBM(dbm.AppRunTaskStatusWaiting)
	err := s.groupDAO.TaskDAO.SaveIfNotExist(ctx, model)
	return entity.SubmitAppResultSucceed, err
}

func (s *ServiceImpl) updateTaskByTestResults(model *dbm.AppRunTask, testResults []*entity.HurlTestResult) {
	model.Total = int32(len(testResults))
	pass := utils.Map(testResults, func(r *entity.HurlTestResult) bool {
		return r.Success
	})
	model.Pass = utils.Reduce(pass, func(sum int32, passed bool) int32 {
		if passed {
			return sum + 1
		}
		return sum
	})
	if model.Pass == model.Total {
		model.Status = dbm.AppRunTaskStatusSucceed
	} else {
		model.Status = dbm.AppRunTaskStatusFail
	}
}

func (s *ServiceImpl) cleanup(ctx context.Context, info *entity.AppInfo, removeFn docker.ContainerRemoveFn) {
	logrus.Info("[TaskService][RunApp] post running, remove the container")
	err := s.groupDAO.FileDAO.Cleanup(ctx, info)
	if err != nil {
		logrus.Warnf("[TaskService][RunApp] post running, call FileDAO.Cleanup error %+v", err)
	}
	if removeFn != nil {
		if err = removeFn(); err != nil {
			logrus.Warnf("[TaskService][RunApp] post running, remove container error %+v", err)
		}
	}
	if err == nil {
		logrus.Info("[TaskService][RunApp] post running, clean up successfully")
	}
}

func (s *ServiceImpl) runAllTests(ctx context.Context, info *entity.AppInfo, testcases []*dbm.Testcase) (string, []*entity.HurlTestResult, error) {
	logDir := info.GetLogDir()
	reportDir := logDir.DirPath
	reportJsonPath := filepath.Join(reportDir, "report.json")

	err := os.Remove(reportJsonPath)
	if err != nil {
		logrus.Warnf("[Hurl DAO][RunAllTests] call os.Remove error %+v", err)
	}

	writer, err := logDir.GetWriter(constants.LogTypeHurlTest)
	if err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests] call logDir.GetWriter error %+v", err)
		return "", nil, err
	}

	for _, testcase := range testcases {
		args := []string{"--report-json", reportDir, "--test"}
		command := exec.Command("hurl", args...)
		command.Stdin = strings.NewReader(testcase.Content)
		command.Stdout = writer
		command.Stderr = writer
		if err = command.Run(); err != nil {
			logrus.Errorf("[TaskService][RunAllTests] call hurl.Run err: %+v", err)
			continue
		}
		logrus.Infof("[TaskService][RunAllTests] call hurl.Run success %s", utils.FormatJsonString(testcase))
	}

	file, err := os.Open(reportJsonPath)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", nil, err
	}

	var results []*entity.HurlTestResult
	err = json.Unmarshal(bytes, &results)
	return string(bytes), results, err
}

func (s *ServiceImpl) RunApp(ctx context.Context, info *entity.AppInfo) error {
	err := s.groupDAO.FileDAO.Unzip(ctx, info)
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call FileDAO.Unzip error %+v", err)
		return err
	}

	stdout, stderr, err := s.groupDAO.FileDAO.PrepareLogFile(ctx, info)
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call FileDAO.PrepareLogFile error %+v", err)
		return err
	}

	removeFn, err := s.groupDAO.DockerDAO.CompileAndRun(ctx, info, stdout, stderr)
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call DockerDAO.CompileAndRun error %+v", err)
		return err
	}

	defer s.cleanup(ctx, info, removeFn)

	testcaseFiles, err := utils.GetAllFileNames(config.Instance.TestcasesDir, ".hurl")
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call GetAllFileNames error %+v", err)
		return err
	}

	testcases, err := s.groupDAO.TestcaseDAO.FindAll(ctx, &dbm.TestcaseFilter{
		Names:  testcaseFiles,
		Status: utils.Pointer(dbm.Active),
	})
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call TestcaseDAO.FindAll error %+v", err)
		return err
	}

	rawResults, testResults, err := s.runAllTests(ctx, info, testcases)
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call HurlDAO.RunAllTests error %+v", err)
		return err
	}

	model, err := s.groupDAO.TaskDAO.FindByUUID(ctx, info.UUID)
	if err != nil {
		return err
	}
	model.TestResults = &rawResults
	s.updateTaskByTestResults(model, testResults)
	logrus.Infof("[TaskService][RunApp] model %+v", model)
	err = s.groupDAO.TaskDAO.Save(ctx, model)
	if err != nil {
		logrus.Errorf("[TaskService][RunApp] call TaskDAO.Save error %+v", err)
	}
	return err
}

func (s *ServiceImpl) ListAppTasks(ctx context.Context, userID *uint, page *entity.Page) (*response.ListAppTasksResponse, error) {
	var (
		modelPage *dbm.ModelPage[*dbm.AppRunTaskWithUser]
		err       error
	)
	filter := &dbm.TaskFilter{
		UserID: userID,
	}
	modelPage, err = s.groupDAO.TaskDAO.ListByPage(ctx, filter, page.ToDBM())

	if err != nil {
		logrus.Errorf("[Task Service][ListAppTasks] list tasks error %+v", err)
		return nil, err
	}
	resp := &response.ListAppTasksResponse{
		Total: modelPage.Total,
		Data:  utils.Map(modelPage.Items, assembler.ConvertAppRunTaskDbmToResponse),
	}
	return resp, nil
}

func (s *ServiceImpl) GetLogFile(ctx context.Context, uuid, logType string) (io.ReadCloser, error) {
	model, err := s.groupDAO.TaskDAO.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if model == nil {
		return nil, err
	}

	userIDStr := strconv.FormatInt(int64(model.UserID), 10)
	logDir := &entity.LogDir{
		DirPath: path.Join(config.Instance.WorkDir, "logs", userIDStr),
		UUID:    model.UUID,
	}

	return logDir.GetReader(logType)
}
