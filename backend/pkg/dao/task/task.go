package task

import (
	"context"

	"autograder/pkg/model/dbm"
	"autograder/pkg/repository/query"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DaoImpl struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) *DaoImpl {
	return &DaoImpl{db}
}

func (d *DaoImpl) FindById(ctx context.Context, id uint) (*dbm.AppRunTask, error) {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

func (d *DaoImpl) FindByUUID(ctx context.Context, UUID string) (*dbm.AppRunTask, error) {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Where(t.UUID.Eq(UUID)).Take()
}

func (d *DaoImpl) Save(ctx context.Context, tasks ...*dbm.AppRunTask) error {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Clauses(&clause.OnConflict{UpdateAll: true}).Save(tasks...)
}

func (d *DaoImpl) SaveIfNotExist(ctx context.Context, tasks ...*dbm.AppRunTask) error {
	t := query.Use(d.db).AppRunTask
	return t.WithContext(ctx).Clauses(&clause.OnConflict{
		DoNothing: true,
	}).Save(tasks...)
}

func (d *DaoImpl) ListUserTasksByPage(ctx context.Context, userID uint, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTask], error) {
	t := query.Use(d.db).AppRunTask
	offset := (page.PageNo - 1) * page.PageSize
	var total int64
	models, total, err := t.WithContext(ctx).
		Where(t.UserID.Eq(userID)).
		Order(t.ID.Desc()).
		FindByPage(offset, page.PageSize)
	if err != nil {
		return nil, err
	}
	return &dbm.ModelPage[*dbm.AppRunTask]{
		Total: total,
		Items: models,
	}, err
}

func (d *DaoImpl) ListTasksByPage(ctx context.Context, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTask], error) {
	t := query.Use(d.db).AppRunTask
	offset := (page.PageNo - 1) * page.PageSize
	var total int64
	models, total, err := t.WithContext(ctx).
		Order(t.ID.Desc()).
		FindByPage(offset, page.PageSize)
	if err != nil {
		return nil, err
	}
	return &dbm.ModelPage[*dbm.AppRunTask]{
		Total: total,
		Items: models,
	}, err
}
