package service

import (
	"autograder/pkg/dao"
	"autograder/pkg/service/canvas"
	"autograder/pkg/service/task"
	"autograder/pkg/service/testcase"
	"autograder/pkg/service/user"
)

type GroupService struct {
	TaskSvc     task.Service
	UserSvc     user.Service
	CanvasSvc   canvas.Service
	TestcaseSvc testcase.Service
}

func NewGroupService(dao *dao.GroupDAO) *GroupService {
	return &GroupService{
		TaskSvc:     task.NewService(dao),
		UserSvc:     user.NewService(dao),
		CanvasSvc:   canvas.NewService(dao),
		TestcaseSvc: testcase.NewService(dao),
	}
}
