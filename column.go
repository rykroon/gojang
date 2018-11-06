package gojang

import ()

type constraints struct {
	null       bool
	unique     bool
	primaryKey bool
}

type Column struct {
	model    *Model
	dbColumn string
	dbType   string
	alias    string

	constraints
	isRelation bool //foreignKey
}

func newColumn(dbType string) *Column {
	return &Column{dbType: dbType}
}

func (c *Column) Alias() string {
	return c.alias
}

func (c *Column) As(alias string) {
	c.alias = alias
}

func (c *Column) Asc() orderByExpression {
	return orderByExpression(c.asSql() + " ASC")
}

func (c *Column) asSql() string {
	sql := ""

	if c.HasModel() {
		tableName := dbq(c.model.dbTable)
		colName := dbq(c.dbColumn)
		sql = tableName + "." + colName

	} else {
		sql = c.dbColumn
	}

	return sql
}

func (c *Column) copy() *Column {
	copy := newColumn(c.dbType)
	copy.model = c.model
	copy.dbColumn = c.dbColumn
	copy.alias = c.alias
	copy.constraints = c.constraints
	return copy
}

func (c *Column) DbColumn() string {
	return c.dbColumn
}

func (c *Column) Desc() orderByExpression {
	return orderByExpression(c.asSql() + " DESC")
}

func (c *Column) HasModel() bool {
	return c.model != nil
}

func (f *Column) HasNullConstraint() bool {
	return f.null
}

func (c *Column) HasPrimaryKeyConstraint() bool {
	return c.primaryKey
}

func (c *Column) HasRelation() bool {
	return c.isRelation
}

func (f *Column) HasUniqueConstraint() bool {
	return f.unique
}

func (c *Column) Model() *Model {
	return c.model
}

func (c *Column) setDbColumn(col string) {
	c.dbColumn = col
}

func (c *Column) setModel(model *Model) {
	c.model = model
}

func (c *Column) setPrimaryKeyConstraint(primaryKey bool) {
	c.primaryKey = primaryKey
}

func (c *Column) setUniqueConstraint(unique bool) {
	c.unique = unique
}
