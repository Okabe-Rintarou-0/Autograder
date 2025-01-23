package dbm

import (
	"gorm.io/gorm"
)

const (
	AppRunTaskStatusWaiting int32 = 1
	AppRunTaskStatusRunning int32 = 2
	AppRunTaskStatusSucceed int32 = 3
	AppRunTaskStatusFail    int32 = 4
)

type AppRunTask struct {
	gorm.Model

	UUID   string `gorm:"type:varchar(36);uniqueIndex"`
	UserID uint   `gorm:"not null"`
	Status int32  `gorm:"not null"`
	Pass   int32  `gorm:"not null"`
	Total  int32  `gorm:"not null"`

	TestResults *string `gorm:"type:json"`
}

type AppRunTaskListResult struct {
	gorm.Model

	UUID        string
	UserID      uint
	Username    string
	RealName    string
	Email       string
	Status      int32
	Pass        int32
	Total       int32
	TestResults *string
}
