package hurl

import (
	"autograder/pkg/model/entity"
	"context"
)

type ContainerRemoveFn func() error

type DAO interface {
	RunAllTests(ctx context.Context, info *entity.AppInfo) ([]*entity.HurlTestResult, error)
}
