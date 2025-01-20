package task

import (
	"context"

	"autograder/pkg/model/dbm"
	"autograder/pkg/model/entity"
	"autograder/pkg/utils"

	"github.com/sirupsen/logrus"
)

func (s *ServiceImpl) dealWithApp(ctx context.Context, info *entity.AppInfo) {
	var err error
	if info == nil {
		return
	}

	model := info.ToDBM(dbm.AppRunTaskStatusRunning)

	defer func() {
		task, dbErr := s.groupDAO.TaskDAO.FindByUUID(ctx, info.UUID)
		if dbErr != nil || task == nil {
			return
		}
		if err != nil {
			task.Status = dbm.AppRunTaskStatusFail
			_ = s.groupDAO.TaskDAO.Save(ctx, task)
		}
		s.removeUserTask(task.UserID)
	}()

	err = s.groupDAO.TaskDAO.Save(ctx, model)
	if err != nil {
		logrus.Infof("[App Runner Service][worker] saving task failed, error %+v", err)
		return
	}
	logrus.Infof("[App Runner Service][worker] run app with info %s", utils.FormatJsonString(info))
	err = s.RunApp(ctx, info)
	if err != nil {
		logrus.Errorf("[App Runner Service][worker] run app error %+v", err)
	}
}

func (s *ServiceImpl) worker(ctx context.Context) {
	logrus.Infof("[App Runner Service][worker] start running")
	for info := range s.workerCh {
		s.dealWithApp(ctx, info)
	}
}
