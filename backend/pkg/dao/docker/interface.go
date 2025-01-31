package docker

import (
	"context"
	"io"

	"autograder/pkg/model/entity"
)

type DAO interface {
	CompileAndRun(ctx context.Context, info *entity.AppInfo, stdoutWriter, stderrWriter io.WriteCloser) (string, error)
	RemoveContainer(ctx context.Context, containerID string) error
}
