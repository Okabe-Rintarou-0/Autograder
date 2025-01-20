// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"autograder/pkg/model/dbm"
)

func newAppRunTask(db *gorm.DB, opts ...gen.DOOption) appRunTask {
	_appRunTask := appRunTask{}

	_appRunTask.appRunTaskDo.UseDB(db, opts...)
	_appRunTask.appRunTaskDo.UseModel(&dbm.AppRunTask{})

	tableName := _appRunTask.appRunTaskDo.TableName()
	_appRunTask.ALL = field.NewAsterisk(tableName)
	_appRunTask.ID = field.NewUint(tableName, "id")
	_appRunTask.CreatedAt = field.NewTime(tableName, "created_at")
	_appRunTask.UpdatedAt = field.NewTime(tableName, "updated_at")
	_appRunTask.DeletedAt = field.NewField(tableName, "deleted_at")
	_appRunTask.UUID = field.NewString(tableName, "uuid")
	_appRunTask.UserID = field.NewUint(tableName, "user_id")
	_appRunTask.Status = field.NewInt32(tableName, "status")
	_appRunTask.Pass = field.NewInt32(tableName, "pass")
	_appRunTask.Total = field.NewInt32(tableName, "total")
	_appRunTask.TestResult = field.NewString(tableName, "test_result")

	_appRunTask.fillFieldMap()

	return _appRunTask
}

type appRunTask struct {
	appRunTaskDo appRunTaskDo

	ALL        field.Asterisk
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	UUID       field.String
	UserID     field.Uint
	Status     field.Int32
	Pass       field.Int32
	Total      field.Int32
	TestResult field.String

	fieldMap map[string]field.Expr
}

func (a appRunTask) Table(newTableName string) *appRunTask {
	a.appRunTaskDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appRunTask) As(alias string) *appRunTask {
	a.appRunTaskDo.DO = *(a.appRunTaskDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appRunTask) updateTableName(table string) *appRunTask {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.UUID = field.NewString(table, "uuid")
	a.UserID = field.NewUint(table, "user_id")
	a.Status = field.NewInt32(table, "status")
	a.Pass = field.NewInt32(table, "pass")
	a.Total = field.NewInt32(table, "total")
	a.TestResult = field.NewString(table, "test_result")

	a.fillFieldMap()

	return a
}

func (a *appRunTask) WithContext(ctx context.Context) IAppRunTaskDo {
	return a.appRunTaskDo.WithContext(ctx)
}

func (a appRunTask) TableName() string { return a.appRunTaskDo.TableName() }

func (a appRunTask) Alias() string { return a.appRunTaskDo.Alias() }

func (a appRunTask) Columns(cols ...field.Expr) gen.Columns { return a.appRunTaskDo.Columns(cols...) }

func (a *appRunTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appRunTask) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["uuid"] = a.UUID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["status"] = a.Status
	a.fieldMap["pass"] = a.Pass
	a.fieldMap["total"] = a.Total
	a.fieldMap["test_result"] = a.TestResult
}

func (a appRunTask) clone(db *gorm.DB) appRunTask {
	a.appRunTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appRunTask) replaceDB(db *gorm.DB) appRunTask {
	a.appRunTaskDo.ReplaceDB(db)
	return a
}

type appRunTaskDo struct{ gen.DO }

type IAppRunTaskDo interface {
	gen.SubQuery
	Debug() IAppRunTaskDo
	WithContext(ctx context.Context) IAppRunTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppRunTaskDo
	WriteDB() IAppRunTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppRunTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppRunTaskDo
	Not(conds ...gen.Condition) IAppRunTaskDo
	Or(conds ...gen.Condition) IAppRunTaskDo
	Select(conds ...field.Expr) IAppRunTaskDo
	Where(conds ...gen.Condition) IAppRunTaskDo
	Order(conds ...field.Expr) IAppRunTaskDo
	Distinct(cols ...field.Expr) IAppRunTaskDo
	Omit(cols ...field.Expr) IAppRunTaskDo
	Join(table schema.Tabler, on ...field.Expr) IAppRunTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppRunTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppRunTaskDo
	Group(cols ...field.Expr) IAppRunTaskDo
	Having(conds ...gen.Condition) IAppRunTaskDo
	Limit(limit int) IAppRunTaskDo
	Offset(offset int) IAppRunTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppRunTaskDo
	Unscoped() IAppRunTaskDo
	Create(values ...*dbm.AppRunTask) error
	CreateInBatches(values []*dbm.AppRunTask, batchSize int) error
	Save(values ...*dbm.AppRunTask) error
	First() (*dbm.AppRunTask, error)
	Take() (*dbm.AppRunTask, error)
	Last() (*dbm.AppRunTask, error)
	Find() ([]*dbm.AppRunTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbm.AppRunTask, err error)
	FindInBatches(result *[]*dbm.AppRunTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*dbm.AppRunTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppRunTaskDo
	Assign(attrs ...field.AssignExpr) IAppRunTaskDo
	Joins(fields ...field.RelationField) IAppRunTaskDo
	Preload(fields ...field.RelationField) IAppRunTaskDo
	FirstOrInit() (*dbm.AppRunTask, error)
	FirstOrCreate() (*dbm.AppRunTask, error)
	FindByPage(offset int, limit int) (result []*dbm.AppRunTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppRunTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appRunTaskDo) Debug() IAppRunTaskDo {
	return a.withDO(a.DO.Debug())
}

func (a appRunTaskDo) WithContext(ctx context.Context) IAppRunTaskDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appRunTaskDo) ReadDB() IAppRunTaskDo {
	return a.Clauses(dbresolver.Read)
}

func (a appRunTaskDo) WriteDB() IAppRunTaskDo {
	return a.Clauses(dbresolver.Write)
}

func (a appRunTaskDo) Session(config *gorm.Session) IAppRunTaskDo {
	return a.withDO(a.DO.Session(config))
}

func (a appRunTaskDo) Clauses(conds ...clause.Expression) IAppRunTaskDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appRunTaskDo) Returning(value interface{}, columns ...string) IAppRunTaskDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appRunTaskDo) Not(conds ...gen.Condition) IAppRunTaskDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appRunTaskDo) Or(conds ...gen.Condition) IAppRunTaskDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appRunTaskDo) Select(conds ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appRunTaskDo) Where(conds ...gen.Condition) IAppRunTaskDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appRunTaskDo) Order(conds ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appRunTaskDo) Distinct(cols ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appRunTaskDo) Omit(cols ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appRunTaskDo) Join(table schema.Tabler, on ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appRunTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appRunTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appRunTaskDo) Group(cols ...field.Expr) IAppRunTaskDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appRunTaskDo) Having(conds ...gen.Condition) IAppRunTaskDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appRunTaskDo) Limit(limit int) IAppRunTaskDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appRunTaskDo) Offset(offset int) IAppRunTaskDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appRunTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppRunTaskDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appRunTaskDo) Unscoped() IAppRunTaskDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appRunTaskDo) Create(values ...*dbm.AppRunTask) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appRunTaskDo) CreateInBatches(values []*dbm.AppRunTask, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appRunTaskDo) Save(values ...*dbm.AppRunTask) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appRunTaskDo) First() (*dbm.AppRunTask, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.AppRunTask), nil
	}
}

func (a appRunTaskDo) Take() (*dbm.AppRunTask, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.AppRunTask), nil
	}
}

func (a appRunTaskDo) Last() (*dbm.AppRunTask, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.AppRunTask), nil
	}
}

func (a appRunTaskDo) Find() ([]*dbm.AppRunTask, error) {
	result, err := a.DO.Find()
	return result.([]*dbm.AppRunTask), err
}

func (a appRunTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbm.AppRunTask, err error) {
	buf := make([]*dbm.AppRunTask, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appRunTaskDo) FindInBatches(result *[]*dbm.AppRunTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appRunTaskDo) Attrs(attrs ...field.AssignExpr) IAppRunTaskDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appRunTaskDo) Assign(attrs ...field.AssignExpr) IAppRunTaskDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appRunTaskDo) Joins(fields ...field.RelationField) IAppRunTaskDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appRunTaskDo) Preload(fields ...field.RelationField) IAppRunTaskDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appRunTaskDo) FirstOrInit() (*dbm.AppRunTask, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.AppRunTask), nil
	}
}

func (a appRunTaskDo) FirstOrCreate() (*dbm.AppRunTask, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.AppRunTask), nil
	}
}

func (a appRunTaskDo) FindByPage(offset int, limit int) (result []*dbm.AppRunTask, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appRunTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appRunTaskDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appRunTaskDo) Delete(models ...*dbm.AppRunTask) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appRunTaskDo) withDO(do gen.Dao) *appRunTaskDo {
	a.DO = *do.(*gen.DO)
	return a
}
