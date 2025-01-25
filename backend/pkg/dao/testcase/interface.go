package testcase

import (
	"context"

	"autograder/pkg/model/dbm"
)

type DAO interface {
	FindById(ctx context.Context, id uint) (*dbm.Testcase, error)
	SaveIfNotExist(ctx context.Context, testcases ...*dbm.Testcase) error
	Save(ctx context.Context, testcases ...*dbm.Testcase) error
	FindAll(ctx context.Context, filter *dbm.TestcaseFilter) ([]*dbm.Testcase, error)
	DeleteAll(ctx context.Context, filter *dbm.TestcaseFilter) error
}
