package dao

import (
	"autograder/pkg/dao/docker"
	"autograder/pkg/dao/file"
	"autograder/pkg/dao/hurl"
	"autograder/pkg/dao/task"
	"autograder/pkg/dao/user"

	"gorm.io/gorm"
)

type GroupDAO struct {
	DockerDAO docker.DAO
	FileDAO   file.DAO
	HurlDAO   hurl.DAO
	UserDAO   user.DAO
	TaskDAO   task.DAO
}

func NewGroupDAO(systemDB, _ *gorm.DB) *GroupDAO {
	return &GroupDAO{
		DockerDAO: docker.NewDAO(),
		FileDAO:   file.NewDao(),
		HurlDAO:   hurl.NewDAO(),
		UserDAO:   user.NewDAO(systemDB),
		TaskDAO:   task.NewDAO(systemDB),
	}
}
