package user

import (
	"context"

	"autograder/pkg/model/dbm"
	"autograder/pkg/repository/query"

	"gorm.io/gorm"
)

type DaoImpl struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) *DaoImpl {
	return &DaoImpl{db}
}

func (d *DaoImpl) FindById(ctx context.Context, id uint) (*dbm.User, error) {
	u := query.Use(d.db).User
	return u.WithContext(ctx).Where(u.ID.Eq(id)).Take()
}

func (d *DaoImpl) FindByUsernameOrEmail(ctx context.Context, username, email string) (*dbm.User, error) {
	u := query.Use(d.db).User
	return u.WithContext(ctx).Where(u.Username.Eq(username)).Or(u.Email.Eq(email)).Take()
}

func (d *DaoImpl) Save(ctx context.Context, users ...*dbm.User) error {
	u := query.Use(d.db).User
	return u.WithContext(ctx).Save(users...)
}

func (d *DaoImpl) ListByPage(ctx context.Context, page *dbm.Page) (*dbm.ModelPage[*dbm.User], error) {
	u := query.Use(d.db).User
	offset := (page.PageNo - 1) * page.PageSize
	var total int64
	models, total, err := u.WithContext(ctx).
		Order(u.ID.Desc()).
		FindByPage(offset, page.PageSize)
	if err != nil {
		return nil, err
	}
	return &dbm.ModelPage[*dbm.User]{
		Total: total,
		Items: models,
	}, err
}
