package docker

import (
	"context"
	"io"

	"autograder/pkg/model/entity"
)

type ContainerRemoveFn func() error

type DAO interface {
	CompileAndRun(ctx context.Context, info *entity.AppInfo, stdoutWriter, stderrWriter io.Writer) (ContainerRemoveFn, error)
}
