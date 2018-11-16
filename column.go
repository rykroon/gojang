package gojang

import (
	"fmt"
)

type onDelete string

const (
	Cascade    onDelete = "CASCADE"
	Protect    onDelete = "RESTRICT"
	SetNull    onDelete = "SET NULL"
	SetDefault onDelete = "SET DEFAULT"
)

type foreignKeyConstraint struct {
	foreignKey   bool
	onDelete     onDelete
	relatedModel *Model
}

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

type columner interface {
	expression

	ColumnName() string
	setColumnName(string)

	DataType() string
	setDataType(string)

	Null() bool
	setNull(bool)

	Unique() bool
	setUnique(bool)

	PrimaryKey() bool
	setPrimaryKey(bool)
}

type orderByExpression string

//Constructor
func newColumn(dataType string) *Column {
	return &Column{dataType: dataType}
}

func (c *Column) asSql() string {
	return c.QualifiedName()
}

//Getters and Setters

func (c *Column) HasModel() bool {
	return c.model != nil
}

func (c *Column) Model() *Model {
	return c.model
}

func (c *Column) setModel(model *Model) {
	c.model = model
}

func (c *Column) QualifiedName() string {
	if c.HasModel() {
		return fmt.Sprintf("%v.%v", dbq(c.model.dbTable), c.UnqualifiedName())
	}

	return c.UnqualifiedName()
}

func (c *Column) UnqualifiedName() string {
	return dbq(c.columnName)
}

func (c *Column) ColumnName() string {
	return c.columnName
}

func (c *Column) setColumnName(name string) {
	c.columnName = name
}

func (c *Column) DataType() string {
	return c.dataType
}

func (c *Column) setDataType(dataType string) {
	c.dataType = dataType
}

func (c *Column) PrimaryKey() bool {
	return c.primaryKey
}

func (c *Column) setPrimaryKey(primaryKey bool) {
	c.primaryKey = primaryKey
}

func (c *Column) Unique() bool {
	return c.unique
}

func (c *Column) setUnique(unique bool) {
	c.unique = unique
}

func (c *Column) Null() bool {
	return c.null
}

func (c *Column) setNull(null bool) {
	c.null = null
}

func (c *Column) Alias() string {
	return c.alias
}

func (c *Column) As(alias string) {
	c.alias = alias
}

//
// OrderByExpressions
//

func (c *Column) Asc() orderByExpression {
	return orderByExpression(c.asSql() + " ASC")
}

func (c *Column) Desc() orderByExpression {
	return orderByExpression(c.asSql() + " DESC")
}

// func (c *Column) copy() *Column {
// 	copy := newColumn(c.dataType)
// 	copy.model = c.model
// 	copy.columnName = c.columnName
// 	copy.constraints = c.constraints
// 	copy.alias = c.alias
// 	return copy
// }

func (c *Column) HasRelation() bool {
	return c.isRelation
}

func create(f field) string {
	s := dbq(f.ColumnName()) + " " + f.DataType()

	if f.PrimaryKey() {
		s += " PRIMARY KEY"
	} else {

		if f.HasRelation() {
			fkey := f.(relatedField)
			s += " REFERENCES " + dbq(fkey.getRelatedModel().dbTable) + " ON DELETE " + fkey.getOnDelete()
		}

		if f.Null() {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.Unique() {
			s += " UNIQUE"
		}
	}

	return s
}
