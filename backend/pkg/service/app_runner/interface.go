package apprunner

import (
	"autograder/pkg/entity"
	"context"
)

type Service interface {
	RunApp(ctx context.Context, info *entity.AppInfo) error
}
