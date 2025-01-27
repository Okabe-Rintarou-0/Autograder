package testcase

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"autograder/pkg/model/dbm"
	"autograder/pkg/repository/query"
)

type DaoImpl struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) *DaoImpl {
	return &DaoImpl{db}
}

func (d *DaoImpl) FindById(ctx context.Context, id uint) (*dbm.Testcase, error) {
	t := query.Use(d.db).Testcase
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

func (d *DaoImpl) getConditions(filter *dbm.TestcaseFilter) []gen.Condition {
	t := query.Use(d.db).Testcase
	var conditions []gen.Condition
	if filter == nil {
		return conditions
	}
	if len(filter.Names) > 0 {
		conditions = append(conditions, t.Name.In(filter.Names...))
	}
	if filter.Status != nil {
		conditions = append(conditions, t.Status.Eq(*filter.Status))
	}
	return conditions
}

func (d *DaoImpl) FindAll(ctx context.Context, filter *dbm.TestcaseFilter) ([]*dbm.Testcase, error) {
	t := query.Use(d.db).Testcase
	return t.WithContext(ctx).Where(d.getConditions(filter)...).Find()
}

func (d *DaoImpl) SaveIfNotExist(ctx context.Context, testcases ...*dbm.Testcase) error {
	t := query.Use(d.db).Testcase
	return t.WithContext(ctx).Clauses(&clause.OnConflict{
		DoNothing: true,
	}).Create(testcases...)
}

func (d *DaoImpl) Save(ctx context.Context, testcases ...*dbm.Testcase) error {
	t := query.Use(d.db).Testcase
	return t.WithContext(ctx).Save(testcases...)
}

func (d *DaoImpl) DeleteAll(ctx context.Context, filter *dbm.TestcaseFilter) error {
	t := query.Use(d.db).Testcase
	_, err := t.WithContext(ctx).Where(d.getConditions(filter)...).Delete()
	return err
}
