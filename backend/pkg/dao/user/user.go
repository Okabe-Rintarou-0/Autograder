package user

import (
	"context"

	"gorm.io/gen"

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

func (d *DaoImpl) getConditions(filter *dbm.UserFilter) []gen.Condition {
	u := query.Use(d.db).User
	var conditions []gen.Condition
	if filter.RealName != nil {
		conditions = append(conditions, u.RealName.Like(*filter.RealName+"%"))
	}
	if filter.Username != nil {
		conditions = append(conditions, u.Username.Like(*filter.Username+"%"))
	}
	if filter.Email != nil {
		conditions = append(conditions, u.Email.Like(*filter.Email+"%"))
	}
	return conditions
}

func (d *DaoImpl) getDAO(ctx context.Context, filter *dbm.UserFilter) query.IUserDo {
	u := query.Use(d.db).User
	dao := u.WithContext(ctx)
	if filter == nil {
		return dao
	}
	conditions := d.getConditions(filter)
	if filter.Or {
		for i, condition := range conditions {
			if i == 0 {
				dao = dao.Where(condition)
			} else {
				dao = dao.Or(condition)
			}
		}
		return dao
	}
	return dao.Where(conditions...)
}

func (d *DaoImpl) ListByPage(ctx context.Context, filter *dbm.UserFilter, page *dbm.Page) (*dbm.ModelPage[*dbm.User], error) {
	u := query.Use(d.db).User
	offset := (page.PageNo - 1) * page.PageSize
	var total int64

	dao := d.getDAO(ctx, filter)
	models, total, err := dao.
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
