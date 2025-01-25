package task

import (
	"context"

	"gorm.io/gen"

	"autograder/pkg/utils"

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

func (d *DaoImpl) ListByPage(ctx context.Context, filter *dbm.TaskFilter, page *dbm.Page) (*dbm.ModelPage[*dbm.AppRunTaskWithUser], error) {
	t := query.Use(d.db).AppRunTask
	offset := (page.PageNo - 1) * page.PageSize
	var total int64
	tasks, total, err := t.WithContext(ctx).
		Where(d.getConditions(filter)...).
		Order(t.ID.Desc()).
		FindByPage(offset, page.PageSize)

	if err != nil {
		return nil, err
	}

	userIDs := utils.Map(tasks, func(v *dbm.AppRunTask) uint {
		return v.UserID
	})

	u := query.Use(d.db).User
	users, err := u.WithContext(ctx).Where(u.ID.In(userIDs...)).Find()
	if err != nil {
		return nil, err
	}
	usersMap := utils.IntoMap(users, func(v *dbm.User) uint {
		return v.ID
	})

	var models []*dbm.AppRunTaskWithUser
	for _, task := range tasks {
		user, ok := usersMap[task.UserID]
		if !ok {
			continue
		}
		models = append(models, &dbm.AppRunTaskWithUser{
			Model:       task.Model,
			UUID:        task.UUID,
			UserID:      task.UserID,
			Username:    user.Username,
			RealName:    user.RealName,
			Email:       user.Email,
			Status:      task.Status,
			Pass:        task.Pass,
			Total:       task.Total,
			TestResults: task.TestResults,
		})
	}

	return &dbm.ModelPage[*dbm.AppRunTaskWithUser]{
		Total: total,
		Items: models,
	}, err
}

func (d *DaoImpl) getConditions(filter *dbm.TaskFilter) []gen.Condition {
	t := query.Use(d.db).AppRunTask
	var conditions []gen.Condition
	if filter.UserID != nil {
		conditions = append(conditions, t.UserID.Eq(*filter.UserID))
	}
	return conditions
}
