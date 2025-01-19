package task

import (
	"context"

	"autograder/pkg/model/dbm"
	"autograder/pkg/repository/query"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type daoImpl struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) *daoImpl {
	return &daoImpl{db}
}

func (d *daoImpl) FindById(ctx context.Context, id uint) (*dbm.AppRunTask, error) {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

func (d *daoImpl) FindByUUID(ctx context.Context, UUID string) (*dbm.AppRunTask, error) {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Where(t.UUID.Eq(UUID)).Take()
}

func (d *daoImpl) Save(ctx context.Context, tasks ...*dbm.AppRunTask) error {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Save(tasks...)
}

func (d *daoImpl) SaveIfNotExist(ctx context.Context, tasks ...*dbm.AppRunTask) error {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Clauses(&clause.OnConflict{
		DoNothing: true,
	}).Save(tasks...)
}

func (d *daoImpl) ListUserTasksByPage(ctx context.Context, userID uint, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTask], error) {
	t := query.Use(d.db).AppRunTask
	offset := (page.PageNo - 1) * page.PageSize
	var total int64
	var models []*dbm.AppRunTask
	err := d.db.WithContext(ctx).
		Where(t.UserID.Eq(userID)).
		Offset(offset).
		Limit(page.PageSize).
		Order(t.ID.Asc()).
		Find(&models).
		Limit(-1).
		Count(&total).
		Error
	if err != nil {
		return nil, err
	}
	return &dbm.ModelPage[*dbm.AppRunTask]{
		Total: total,
		Items: models,
	}, err
}
