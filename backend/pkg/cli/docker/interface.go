package docker

import (
	"context"

	"autograder/pkg/entity"
)

type Client interface {
	PullImage(ctx context.Context, imageName string) error
	RunContainer(ctx context.Context, config *entity.DockerCreateConfig) (string, error)
	RemoveContainer(ctx context.Context, containerID string) error
	ExecuteContainer(ctx context.Context, containerID string, commands []string) error
}
