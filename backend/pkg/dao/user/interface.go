package user

import (
	"context"

	"autograder/pkg/model/dbm"
)

type DAO interface {
	Find(ctx context.Context, filter *dbm.UserFilter) (*dbm.User, error)
	Save(ctx context.Context, users ...*dbm.User) error
	ListByPage(ctx context.Context, filter *dbm.UserFilter, page *dbm.Page) (*dbm.ModelPage[*dbm.User], error)
}
