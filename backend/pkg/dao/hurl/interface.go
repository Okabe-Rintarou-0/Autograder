package hurl

import (
	"context"

	"autograder/pkg/model/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	RunAllTests(ctx context.Context, info *entity.AppInfo) (string, []*entity.HurlTestResult, error)
}
