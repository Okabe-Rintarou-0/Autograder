package hurl

import (
	"autograder/pkg/model/dbm"
	"context"

	"autograder/pkg/model/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	RunAllTests(ctx context.Context, info *entity.AppInfo, testcases []*dbm.Testcase) (string, []*entity.HurlTestResult, error)
}
