// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"games/app/dal/system/model"
)

func newCategoryItemRelation(db *gorm.DB, opts ...gen.DOOption) categoryItemRelation {
	_categoryItemRelation := categoryItemRelation{}

	_categoryItemRelation.categoryItemRelationDo.UseDB(db, opts...)
	_categoryItemRelation.categoryItemRelationDo.UseModel(&model.CategoryItemRelation{})

	tableName := _categoryItemRelation.categoryItemRelationDo.TableName()
	_categoryItemRelation.ALL = field.NewAsterisk(tableName)
	_categoryItemRelation.ID = field.NewUint64(tableName, "id")
	_categoryItemRelation.CategoryID = field.NewInt(tableName, "category_id")
	_categoryItemRelation.ItemType = field.NewInt(tableName, "item_type")
	_categoryItemRelation.ItemID = field.NewInt(tableName, "item_id")
	_categoryItemRelation.CreatedAt = field.NewTime(tableName, "created_at")
	_categoryItemRelation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_categoryItemRelation.DeletedAt = field.NewField(tableName, "deleted_at")

	_categoryItemRelation.fillFieldMap()

	return _categoryItemRelation
}

// categoryItemRelation 分类-内容-关联关系
type categoryItemRelation struct {
	categoryItemRelationDo categoryItemRelationDo

	ALL        field.Asterisk
	ID         field.Uint64
	CategoryID field.Int // 分类ID: category.id
	ItemType   field.Int // 内容类型: 1=游戏，2=文章，
	ItemID     field.Int // 内容ID
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (c categoryItemRelation) Table(newTableName string) *categoryItemRelation {
	c.categoryItemRelationDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c categoryItemRelation) As(alias string) *categoryItemRelation {
	c.categoryItemRelationDo.DO = *(c.categoryItemRelationDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *categoryItemRelation) updateTableName(table string) *categoryItemRelation {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint64(table, "id")
	c.CategoryID = field.NewInt(table, "category_id")
	c.ItemType = field.NewInt(table, "item_type")
	c.ItemID = field.NewInt(table, "item_id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *categoryItemRelation) WithContext(ctx context.Context) *categoryItemRelationDo {
	return c.categoryItemRelationDo.WithContext(ctx)
}

func (c categoryItemRelation) TableName() string { return c.categoryItemRelationDo.TableName() }

func (c categoryItemRelation) Alias() string { return c.categoryItemRelationDo.Alias() }

func (c categoryItemRelation) Columns(cols ...field.Expr) gen.Columns {
	return c.categoryItemRelationDo.Columns(cols...)
}

func (c *categoryItemRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *categoryItemRelation) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["category_id"] = c.CategoryID
	c.fieldMap["item_type"] = c.ItemType
	c.fieldMap["item_id"] = c.ItemID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c categoryItemRelation) clone(db *gorm.DB) categoryItemRelation {
	c.categoryItemRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c categoryItemRelation) replaceDB(db *gorm.DB) categoryItemRelation {
	c.categoryItemRelationDo.ReplaceDB(db)
	return c
}

type categoryItemRelationDo struct{ gen.DO }

func (c categoryItemRelationDo) Debug() *categoryItemRelationDo {
	return c.withDO(c.DO.Debug())
}

func (c categoryItemRelationDo) WithContext(ctx context.Context) *categoryItemRelationDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoryItemRelationDo) ReadDB() *categoryItemRelationDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoryItemRelationDo) WriteDB() *categoryItemRelationDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoryItemRelationDo) Session(config *gorm.Session) *categoryItemRelationDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoryItemRelationDo) Clauses(conds ...clause.Expression) *categoryItemRelationDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoryItemRelationDo) Returning(value interface{}, columns ...string) *categoryItemRelationDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoryItemRelationDo) Not(conds ...gen.Condition) *categoryItemRelationDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoryItemRelationDo) Or(conds ...gen.Condition) *categoryItemRelationDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoryItemRelationDo) Select(conds ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoryItemRelationDo) Where(conds ...gen.Condition) *categoryItemRelationDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoryItemRelationDo) Order(conds ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoryItemRelationDo) Distinct(cols ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoryItemRelationDo) Omit(cols ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoryItemRelationDo) Join(table schema.Tabler, on ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoryItemRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoryItemRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoryItemRelationDo) Group(cols ...field.Expr) *categoryItemRelationDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoryItemRelationDo) Having(conds ...gen.Condition) *categoryItemRelationDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoryItemRelationDo) Limit(limit int) *categoryItemRelationDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoryItemRelationDo) Offset(offset int) *categoryItemRelationDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoryItemRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *categoryItemRelationDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoryItemRelationDo) Unscoped() *categoryItemRelationDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoryItemRelationDo) Create(values ...*model.CategoryItemRelation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoryItemRelationDo) CreateInBatches(values []*model.CategoryItemRelation, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoryItemRelationDo) Save(values ...*model.CategoryItemRelation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoryItemRelationDo) First() (*model.CategoryItemRelation, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryItemRelation), nil
	}
}

func (c categoryItemRelationDo) Take() (*model.CategoryItemRelation, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryItemRelation), nil
	}
}

func (c categoryItemRelationDo) Last() (*model.CategoryItemRelation, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryItemRelation), nil
	}
}

func (c categoryItemRelationDo) Find() ([]*model.CategoryItemRelation, error) {
	result, err := c.DO.Find()
	return result.([]*model.CategoryItemRelation), err
}

func (c categoryItemRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CategoryItemRelation, err error) {
	buf := make([]*model.CategoryItemRelation, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoryItemRelationDo) FindInBatches(result *[]*model.CategoryItemRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoryItemRelationDo) Attrs(attrs ...field.AssignExpr) *categoryItemRelationDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoryItemRelationDo) Assign(attrs ...field.AssignExpr) *categoryItemRelationDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoryItemRelationDo) Joins(fields ...field.RelationField) *categoryItemRelationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoryItemRelationDo) Preload(fields ...field.RelationField) *categoryItemRelationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoryItemRelationDo) FirstOrInit() (*model.CategoryItemRelation, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryItemRelation), nil
	}
}

func (c categoryItemRelationDo) FirstOrCreate() (*model.CategoryItemRelation, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryItemRelation), nil
	}
}

func (c categoryItemRelationDo) FindByPage(offset int, limit int) (result []*model.CategoryItemRelation, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c categoryItemRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoryItemRelationDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoryItemRelationDo) Delete(models ...*model.CategoryItemRelation) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoryItemRelationDo) withDO(do gen.Dao) *categoryItemRelationDo {
	c.DO = *do.(*gen.DO)
	return c
}