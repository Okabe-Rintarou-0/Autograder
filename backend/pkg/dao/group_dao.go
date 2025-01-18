package dao

import (
	"autograder/pkg/dao/docker"
	"autograder/pkg/dao/file"
	"autograder/pkg/dao/hurl"
)

type GroupDAO struct {
	DockerDAO docker.DAO
	FileDAO   file.DAO
	HurlDAO   hurl.DAO
}

func NewGroupDAO() *GroupDAO {
	return &GroupDAO{
		DockerDAO: docker.NewDAO(),
		FileDAO:   file.NewDao(),
		HurlDAO:   hurl.NewDAO(),
	}
}
