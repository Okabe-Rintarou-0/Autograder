package task

import (
	"context"
	"io"
	"path"
	"strconv"
	"sync"

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

	rawResults, testResults, err := s.groupDAO.HurlDAO.RunAllTests(ctx, info)
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
