package gojang

import ()

type onDelete string

const (
	Cascade    onDelete = "CASCADE"
	Protect    onDelete = "RESTRICT"
	SetNull    onDelete = "SET NULL"
	SetDefault onDelete = "SET DEFAULT"
)

type constraints struct {
	null       bool
	unique     bool
	primaryKey bool
	foreignKey bool
	onDelete   onDelete
}

type Column struct {
	model      *Model
	columnName string
	dataType   string
	constraints

	alias      string
	isRelation bool //foreignKey
}

func newColumn(dataType string) *Column {
	return &Column{dataType: dataType}
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
		colName := dbq(c.columnName)
		sql = tableName + "." + colName

	} else {
		sql = c.columnName
	}

	return sql
}

func (c *Column) copy() *Column {
	copy := newColumn(c.dataType)
	copy.model = c.model
	copy.columnName = c.columnName
	copy.alias = c.alias
	copy.constraints = c.constraints
	return copy
}

func (c *Column) ColumnName() string {
	return c.columnName
}

func (c *Column) setColumnName(name string) {
	c.columnName = name
}

func (f *Column) DataType() string {
	return f.dataType
}

func (c *Column) setDataType(dataType string) {
	c.dataType = dataType
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

func (c *Column) setModel(model *Model) {
	c.model = model
}

func (c *Column) setPrimaryKeyConstraint(primaryKey bool) {
	c.primaryKey = primaryKey
}

func (c *Column) setUniqueConstraint(unique bool) {
	c.unique = unique
}
