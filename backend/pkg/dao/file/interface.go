package file

import (
	"context"
	"io"

	"autograder/pkg/model/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	Unzip(ctx context.Context, info *entity.AppInfo) error
	PrepareLogFile(ctx context.Context, info *entity.AppInfo) (io.Writer, io.Writer, error)
	Cleanup(ctx context.Context, info *entity.AppInfo) error
}
