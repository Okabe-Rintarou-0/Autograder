package user

import (
	"context"

	"autograder/pkg/model/dbm"
)

type DAO interface {
	FindById(ctx context.Context, id uint) (*dbm.User, error)
	FindByUsernameOrEmail(ctx context.Context, username, email string) (*dbm.User, error)
	Save(ctx context.Context, users ...*dbm.User) error
	ListByPage(ctx context.Context, page *dbm.Page) (*dbm.ModelPage[*dbm.User], error)
}
