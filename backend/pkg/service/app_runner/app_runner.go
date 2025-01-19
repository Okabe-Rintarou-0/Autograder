package apprunner

import (
	"context"
	"io"
	"path"
	"strconv"

	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
	"autograder/pkg/dao"
	"autograder/pkg/model/dbm"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"
)

type serviceImpl struct {
	groupDAO *dao.GroupDAO
	workerCh chan *entity.AppInfo
}

func NewService(groupDAO *dao.GroupDAO) *serviceImpl {
	svc := &serviceImpl{
		groupDAO: groupDAO,
		workerCh: make(chan *entity.AppInfo, 200),
	}
	go svc.worker(context.Background())
	return svc
}

func (s *serviceImpl) SubmitApp(ctx context.Context, info *entity.AppInfo) (entity.SubmitAppResult, error) {
	isFull := utils.SendIfNotFull(s.workerCh, info)
	if isFull {
		return entity.SubmitAppResultSystemBusy, nil
	}
	model := info.ToDBM(dbm.AppRunTaskStatusWaiting)
	err := s.groupDAO.TaskDAO.SaveIfNotExist(ctx, model)
	return entity.SubmitAppResultSucceed, err
}

func (s *serviceImpl) RunApp(ctx context.Context, info *entity.AppInfo) error {
	err := s.groupDAO.FileDAO.Unzip(ctx, info)
	if err != nil {
		logrus.Errorf("[AppRunnerService][RunApp] call FileDAO.Unzip error %+v", err)
		return err
	}

	stdout, stderr, err := s.groupDAO.FileDAO.PrepareLogFile(ctx, info)
	if err != nil {
		logrus.Errorf("[AppRunnerService][RunApp] call FileDAO.PrepareLogFile error %+v", err)
		return err
	}

	removeFn, err := s.groupDAO.DockerDAO.CompileAndRun(ctx, info, stdout, stderr)
	if err != nil {
		logrus.Errorf("[AppRunnerService][RunApp] call DockerDAO.CompileAndRun error %+v", err)
		return err
	}

	defer func() {
		logrus.Info("[AppRunnerService][RunApp] post running, remove the container")
		err := s.groupDAO.FileDAO.Cleanup(ctx, info)
		if err != nil {
			logrus.Warnf("[AppRunnerService][RunApp] post running, call FileDAO.Cleanup error %+v", err)
		}
		if removeFn != nil {
			if err = removeFn(); err != nil {
				logrus.Warnf("[AppRunnerService][RunApp] post running, remove container error %+v", err)
			}
		}
		if err == nil {
			logrus.Info("[AppRunnerService][RunApp] post running, clean up successfully")
		}
	}()

	err = s.groupDAO.HurlDAO.RunAllTests(ctx)
	if err != nil {
		logrus.Errorf("[AppRunnerService][RunApp] call HurlDAO.RunAllTests error %+v", err)
		return err
	}

	return nil
}

func (s *serviceImpl) ListAppTasks(ctx context.Context, userID uint, page *entity.Page) (*response.ListAppTasksResponse, error) {
	modelPage, err := s.groupDAO.TaskDAO.ListUserTasksByPage(ctx, userID, page.ToDBM())
	if err != nil {
		return nil, err
	}
	resp := &response.ListAppTasksResponse{
		Total: modelPage.Total,
		Data: utils.Map(modelPage.Items, func(m *dbm.AppRunTask) *response.AppRunTask {
			return &response.AppRunTask{
				UUID:      m.UUID,
				UserID:    m.UserID,
				Status:    m.Status,
				CreatedAt: m.CreatedAt,
			}
		}),
	}
	return resp, nil
}

func (s *serviceImpl) GetLogFile(ctx context.Context, uuid, logType string) (io.Reader, error) {
	model, err := s.groupDAO.TaskDAO.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if model == nil {
		return nil, err
	}

	userIDStr := strconv.FormatInt(int64(model.UserID), 10)
	logFile := &entity.LogFile{
		DirPath: path.Join(config.Instance.WorkDir, "logs", userIDStr),
		UUID:    model.UUID,
	}

	switch logType {
	case "stdout":
		return logFile.GetStdoutReader()
	case "stderr":
		return logFile.GetStderrReader()
	}
	return nil, nil
}
