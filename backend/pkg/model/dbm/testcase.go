package dbm

import "gorm.io/gorm"

const (
	Active   int32 = 1
	Inactive int32 = 2
)

type Testcase struct {
	gorm.Model

	Name    string `gorm:"type:varchar(255);uniqueIndex"`
	Path    string `gorm:"type:varchar(1024)"`
	Status  int32  `gorm:"default:1"`
	Content string `gorm:"type:text"`
}

type TestcaseFilter struct {
	Names  []string
	Paths  []string
	Status *int32
}
