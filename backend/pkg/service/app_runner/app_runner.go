package apprunner

import (
	"context"

	"github.com/sirupsen/logrus"

	"autograder/pkg/dao"
	"autograder/pkg/entity"
)

type serviceImpl struct {
	groupDAO *dao.GroupDAO
}

func NewService(groupDAO *dao.GroupDAO) *serviceImpl {
	return &serviceImpl{groupDAO}
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

	if removeFn != nil {
		logrus.Info("[AppRunnerService][RunApp] post running, remove the container")
		defer removeFn()
	}

	err = s.groupDAO.HurlDAO.RunAllTests(ctx)
	if err != nil {
		logrus.Errorf("[AppRunnerService][RunApp] call HurlDAO.RunAllTests error %+v", err)
		return err
	}

	return nil
}
