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
	ListUserTasksByPage(ctx context.Context, userID uint, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTask], error)
	ListTasksByPage(ctx context.Context, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTask], error)
}
