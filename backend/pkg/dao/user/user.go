package user

import (
	"context"

	"autograder/pkg/model/dbm"
	"autograder/pkg/repository/query"

	"gorm.io/gorm"
)

type daoImpl struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) *daoImpl {
	return &daoImpl{db}
}

func (d *daoImpl) FindById(ctx context.Context, id uint) (*dbm.User, error) {
	u := query.Use(d.db).User
	return u.WithContext(ctx).Where(u.ID.Eq(id)).Take()
}

func (d *daoImpl) FindByUsernameOrEmail(ctx context.Context, username, email string) (*dbm.User, error) {
	u := query.Use(d.db).User
	return u.WithContext(ctx).Where(u.Username.Eq(username)).Or(u.Email.Eq(email)).Take()
}

func (d *daoImpl) Save(ctx context.Context, users ...*dbm.User) error {
	u := query.Use(d.db).User
	return u.WithContext(ctx).Save(users...)
}
