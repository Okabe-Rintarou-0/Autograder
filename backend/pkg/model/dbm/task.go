package dbm

import (
	"gorm.io/gorm"
)

const (
	AppRunTaskStatusWaiting int32 = 1
	AppRunTaskStatusRunning int32 = 2
	AppRunTaskStatusSucceed int32 = 3
	AppRunTaskStatusFail    int32 = 4
	AppRunTaskStatusError   int32 = 5
)

type AppRunTask struct {
	gorm.Model

	UUID string `gorm:"type:varchar(36);uniqueIndex"`
	// Actual bound user
	UserID uint `gorm:"not null"`
	// Operator
	OperatorID uint   `gorm:"not null"`
	Status     int32  `gorm:"not null"`
	Error      string `gorm:"type:text"`
	Pass       int32  `gorm:"not null"`
	Total      int32  `gorm:"not null"`

	TestResults *string `gorm:"type:json"`
}

type UserProfile struct {
	ID       uint
	Username string
	RealName string
	Email    string
	Role     int32
}

type AppRunTaskWithUser struct {
	gorm.Model

	UUID        string
	Error       string
	User        *UserProfile
	Operator    *UserProfile
	Status      int32
	Pass        int32
	Total       int32
	TestResults *string
}

type TaskFilter struct {
	UserID     *uint
	OperatorID *uint
}
