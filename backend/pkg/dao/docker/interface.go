package docker

import (
	"context"

	"autograder/pkg/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	CompileAndRun(ctx context.Context, info *entity.AppInfo) (ContainerRemoveFn, error)
}
