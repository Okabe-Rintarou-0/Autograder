package task

import (
	"context"

	"autograder/pkg/model/dbm"
)

type DAO interface {
	FindById(ctx context.Context, id uint) (*dbm.AppRunTask, error)
	FindByUUID(ctx context.Context, UUID string) (*dbm.AppRunTask, error)
	Save(ctx context.Context, tasks ...*dbm.AppRunTask) error
	SaveIfNotExist(ctx context.Context, tasks ...*dbm.AppRunTask) error
	ListByPage(ctx context.Context, filter *dbm.TaskFilter, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTaskWithUser], error)
}
