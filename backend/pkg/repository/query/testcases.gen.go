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

func newTestcase(db *gorm.DB, opts ...gen.DOOption) testcase {
	_testcase := testcase{}

	_testcase.testcaseDo.UseDB(db, opts...)
	_testcase.testcaseDo.UseModel(&dbm.Testcase{})

	tableName := _testcase.testcaseDo.TableName()
	_testcase.ALL = field.NewAsterisk(tableName)
	_testcase.ID = field.NewUint(tableName, "id")
	_testcase.CreatedAt = field.NewTime(tableName, "created_at")
	_testcase.UpdatedAt = field.NewTime(tableName, "updated_at")
	_testcase.DeletedAt = field.NewField(tableName, "deleted_at")
	_testcase.Name = field.NewString(tableName, "name")
	_testcase.Status = field.NewInt32(tableName, "status")
	_testcase.Content = field.NewString(tableName, "content")

	_testcase.fillFieldMap()

	return _testcase
}

type testcase struct {
	testcaseDo testcaseDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Status    field.Int32
	Content   field.String

	fieldMap map[string]field.Expr
}

func (t testcase) Table(newTableName string) *testcase {
	t.testcaseDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t testcase) As(alias string) *testcase {
	t.testcaseDo.DO = *(t.testcaseDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *testcase) updateTableName(table string) *testcase {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.Name = field.NewString(table, "name")
	t.Status = field.NewInt32(table, "status")
	t.Content = field.NewString(table, "content")

	t.fillFieldMap()

	return t
}

func (t *testcase) WithContext(ctx context.Context) ITestcaseDo { return t.testcaseDo.WithContext(ctx) }

func (t testcase) TableName() string { return t.testcaseDo.TableName() }

func (t testcase) Alias() string { return t.testcaseDo.Alias() }

func (t testcase) Columns(cols ...field.Expr) gen.Columns { return t.testcaseDo.Columns(cols...) }

func (t *testcase) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *testcase) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["name"] = t.Name
	t.fieldMap["status"] = t.Status
	t.fieldMap["content"] = t.Content
}

func (t testcase) clone(db *gorm.DB) testcase {
	t.testcaseDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t testcase) replaceDB(db *gorm.DB) testcase {
	t.testcaseDo.ReplaceDB(db)
	return t
}

type testcaseDo struct{ gen.DO }

type ITestcaseDo interface {
	gen.SubQuery
	Debug() ITestcaseDo
	WithContext(ctx context.Context) ITestcaseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITestcaseDo
	WriteDB() ITestcaseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITestcaseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITestcaseDo
	Not(conds ...gen.Condition) ITestcaseDo
	Or(conds ...gen.Condition) ITestcaseDo
	Select(conds ...field.Expr) ITestcaseDo
	Where(conds ...gen.Condition) ITestcaseDo
	Order(conds ...field.Expr) ITestcaseDo
	Distinct(cols ...field.Expr) ITestcaseDo
	Omit(cols ...field.Expr) ITestcaseDo
	Join(table schema.Tabler, on ...field.Expr) ITestcaseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITestcaseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITestcaseDo
	Group(cols ...field.Expr) ITestcaseDo
	Having(conds ...gen.Condition) ITestcaseDo
	Limit(limit int) ITestcaseDo
	Offset(offset int) ITestcaseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITestcaseDo
	Unscoped() ITestcaseDo
	Create(values ...*dbm.Testcase) error
	CreateInBatches(values []*dbm.Testcase, batchSize int) error
	Save(values ...*dbm.Testcase) error
	First() (*dbm.Testcase, error)
	Take() (*dbm.Testcase, error)
	Last() (*dbm.Testcase, error)
	Find() ([]*dbm.Testcase, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbm.Testcase, err error)
	FindInBatches(result *[]*dbm.Testcase, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*dbm.Testcase) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITestcaseDo
	Assign(attrs ...field.AssignExpr) ITestcaseDo
	Joins(fields ...field.RelationField) ITestcaseDo
	Preload(fields ...field.RelationField) ITestcaseDo
	FirstOrInit() (*dbm.Testcase, error)
	FirstOrCreate() (*dbm.Testcase, error)
	FindByPage(offset int, limit int) (result []*dbm.Testcase, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITestcaseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t testcaseDo) Debug() ITestcaseDo {
	return t.withDO(t.DO.Debug())
}

func (t testcaseDo) WithContext(ctx context.Context) ITestcaseDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t testcaseDo) ReadDB() ITestcaseDo {
	return t.Clauses(dbresolver.Read)
}

func (t testcaseDo) WriteDB() ITestcaseDo {
	return t.Clauses(dbresolver.Write)
}

func (t testcaseDo) Session(config *gorm.Session) ITestcaseDo {
	return t.withDO(t.DO.Session(config))
}

func (t testcaseDo) Clauses(conds ...clause.Expression) ITestcaseDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t testcaseDo) Returning(value interface{}, columns ...string) ITestcaseDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t testcaseDo) Not(conds ...gen.Condition) ITestcaseDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t testcaseDo) Or(conds ...gen.Condition) ITestcaseDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t testcaseDo) Select(conds ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t testcaseDo) Where(conds ...gen.Condition) ITestcaseDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t testcaseDo) Order(conds ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t testcaseDo) Distinct(cols ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t testcaseDo) Omit(cols ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t testcaseDo) Join(table schema.Tabler, on ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t testcaseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t testcaseDo) RightJoin(table schema.Tabler, on ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t testcaseDo) Group(cols ...field.Expr) ITestcaseDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t testcaseDo) Having(conds ...gen.Condition) ITestcaseDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t testcaseDo) Limit(limit int) ITestcaseDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t testcaseDo) Offset(offset int) ITestcaseDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t testcaseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITestcaseDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t testcaseDo) Unscoped() ITestcaseDo {
	return t.withDO(t.DO.Unscoped())
}

func (t testcaseDo) Create(values ...*dbm.Testcase) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t testcaseDo) CreateInBatches(values []*dbm.Testcase, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t testcaseDo) Save(values ...*dbm.Testcase) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t testcaseDo) First() (*dbm.Testcase, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.Testcase), nil
	}
}

func (t testcaseDo) Take() (*dbm.Testcase, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.Testcase), nil
	}
}

func (t testcaseDo) Last() (*dbm.Testcase, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.Testcase), nil
	}
}

func (t testcaseDo) Find() ([]*dbm.Testcase, error) {
	result, err := t.DO.Find()
	return result.([]*dbm.Testcase), err
}

func (t testcaseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbm.Testcase, err error) {
	buf := make([]*dbm.Testcase, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t testcaseDo) FindInBatches(result *[]*dbm.Testcase, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t testcaseDo) Attrs(attrs ...field.AssignExpr) ITestcaseDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t testcaseDo) Assign(attrs ...field.AssignExpr) ITestcaseDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t testcaseDo) Joins(fields ...field.RelationField) ITestcaseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t testcaseDo) Preload(fields ...field.RelationField) ITestcaseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t testcaseDo) FirstOrInit() (*dbm.Testcase, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.Testcase), nil
	}
}

func (t testcaseDo) FirstOrCreate() (*dbm.Testcase, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dbm.Testcase), nil
	}
}

func (t testcaseDo) FindByPage(offset int, limit int) (result []*dbm.Testcase, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t testcaseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t testcaseDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t testcaseDo) Delete(models ...*dbm.Testcase) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *testcaseDo) withDO(do gen.Dao) *testcaseDo {
	t.DO = *do.(*gen.DO)
	return t
}
