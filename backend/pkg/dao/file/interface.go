package file

import (
	"context"
	"io"

	"autograder/pkg/model/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	Unzip(ctx context.Context, info *entity.AppInfo) error
	PrepareLogFile(ctx context.Context, info *entity.AppInfo) (io.WriteCloser, io.WriteCloser, error)
	Cleanup(ctx context.Context, info *entity.AppInfo) error
}
