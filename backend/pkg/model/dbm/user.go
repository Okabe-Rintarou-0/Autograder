package dbm

import "gorm.io/gorm"

const (
	CommonUser    = 1
	Administrator = 2
)

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(32);uniqueIndex"`
	RealName string `gorm:"type:varchar(16);uniqueIndex"`
	Password string `gorm:"type:varchar(32);not null"`
	Email    string `gorm:"type:varchar(128);uniqueIndex"`
	Role     int32  `gorm:"not null;default:1"`
}
