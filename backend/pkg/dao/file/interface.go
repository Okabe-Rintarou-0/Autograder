package file

import (
	"context"

	"autograder/pkg/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	Unzip(ctx context.Context, info *entity.AppInfo) error
	// RemoveDir(ctx context.Context) error
}
