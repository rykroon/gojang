package gojang

import (
	"fmt"
	//"reflect"
	"strconv"
)

//Expressions describe a value or a computation that can be used as part of an
//update, create, filter, order by, annotation, or aggregate.
type expression interface {
	asSql() string
}

type selectExpression interface {
	expression
	getValue() interface{}
	As(string) //is essentially the 'setter' method for alias
	Alias() string
	DbType() string
	Scan(interface{}) error
}

type orderByExpression string

type star string

func (a *aggregate) asSql() string {
	f := function(*a)
	return f.asSql()
}

func (a assignment) asSql() string {
	return dbq(a.lhs.DbColumn()) + " " + a.lookupName + " " + a.rhs
}

func (f *function) asSql() string {
	return fmt.Sprintf(f.template, f.args...)
}

func (l lookup) asSql() string {
	sql := l.lhs.asSql() + " " + l.lookupName + " " + l.rhs

	if l.not {
		sql = "NOT(" + sql + ")"
	}
	return sql
}

func (s star) asSql() string {
	return "*"
}

func (c *column) asSql() string {
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

//
//Select Expression Method Set
//

func (a aggregate) Alias() string {
	return function(a).Alias()
}

func (f function) Alias() string {
	return f.outputField.Alias()
}

func (c *column) Alias() string {
	return c.alias
}

func (s star) Alias() string {
	return "*"
}

func (a *aggregate) As(alias string) {
	a.outputField.As(alias)
}

func (f *function) As(alias string) {
	f.outputField.As(alias)
}

func (c *column) As(alias string) {
	c.alias = alias
}

func (s star) As(string) {
	return
}

func (a aggregate) DbType() string {
	return function(a).DbType()
}

func (f function) DbType() string {
	return f.outputField.DbType()
}

func (f *column) DbType() string {
	return f.dbType
}

func (s star) DbType() string {
	return ""
}

func (a aggregate) Scan(v interface{}) error {
	return function(a).Scan(v)
}

func (f function) Scan(v interface{}) error {
	return f.outputField.Scan(v)
}

func (s star) Scan(interface{}) error {
	return nil
}

func (f *BigIntegerField) Scan(value interface{}) error {
	f.Value, f.valid = value.(int64)
	return nil
}

func (f *BooleanField) Scan(value interface{}) error {
	f.Value, f.valid = value.(bool)
	return nil
}

func (f *FloatField) Scan(value interface{}) error {
	switch v := value.(type) {
	case []uint8:
		float, err := strconv.ParseFloat(string(v), 64)
		f.Value = float
		f.valid = err == nil

	case float64:
		f.Value, f.valid = v, true

	default:
		f.Value, f.valid = 0, false
	}

	return nil
}

func (f *IntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.valid = int32(result), ok
	return nil
}

func (f *SmallIntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.valid = int16(result), ok
	return nil
}

func (f *TextField) Scan(value interface{}) error {
	f.Value, f.valid = value.(string)
	return nil
}

func (a aggregate) getValue() interface{} {
	return function(a).getValue()
}

func (f function) getValue() interface{} {
	return f.outputField.getValue()
}

func (s star) getValue() interface{} {
	return nil
}

func (f *BigIntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *BooleanField) getValue() interface{} {
	return f.Value
}

func (f *FloatField) getValue() interface{} {
	return f.Value
}

func (f *IntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *SmallIntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *TextField) getValue() interface{} {
	return f.Value
}

//
//Methods that return Order By Expressions
//

func (c *column) Asc() orderByExpression {
	return orderByExpression(c.asSql() + " ASC")
}

func (c *column) Desc() orderByExpression {
	return orderByExpression(c.asSql() + " DESC")
}
