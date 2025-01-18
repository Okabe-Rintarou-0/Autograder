package hurl

import (
	"context"
)

type ContainerRemoveFn func() error

type DAO interface {
	RunAllTests(ctx context.Context) error
}
