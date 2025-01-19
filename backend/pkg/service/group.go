package service

import (
	"autograder/pkg/dao"
	apprunner "autograder/pkg/service/app_runner"
	"autograder/pkg/service/user"
)

type GroupService struct {
	AppRunnerSvc apprunner.Service
	UserSvc      user.Service
}

func NewGroupService(dao *dao.GroupDAO) *GroupService {
	return &GroupService{
		AppRunnerSvc: apprunner.NewService(dao),
		UserSvc:      user.NewService(dao),
	}
}
