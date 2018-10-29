package gojang

import (
	"fmt"
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
	Scan(interface{}) error
}

type sortExpression struct {
	desc  bool
	field field
}

type star string

func (a *aggregate) asSql() string {
	f := function(*a)
	return f.asSql()
}

func (a assignment) asSql() string {
	return dbq(a.lhs.getDbColumn()) + " " + a.lookupName + " " + a.rhs
}

func (f *function) asSql() string {
	return fmt.Sprintf(f.template, f.args...)
}

func (l lookup) asSql() string {
	sql := l.lhs.asSql() + " " + l.lookupName + " " + l.rhs
	//sql := fieldAsSql(l.lhs, false) + " " + l.lookupName + " " + l.rhs

	if l.not {
		sql = "NOT(" + sql + ")"
	}
	return sql
}

func (e sortExpression) asSql() string {
	if e.desc {
		return e.field.asSql() + " DESC"
	} else {
		return e.field.asSql() + " ASC"
	}
}

func (s star) asSql() string {
	return "*"
}

func fieldAsSql(field field) string {
	sql := ""

	if field.hasModel() {
		tableName := dbq(field.getModel().dbTable)
		colName := dbq(field.getDbColumn())
		sql = tableName + "." + colName

	} else {
		sql = field.getDbColumn()
	}

	return sql
}

func (f *BigIntegerField) asSql() string {
	return fieldAsSql(f)
}

func (f *BooleanField) asSql() string {
	return fieldAsSql(f)
}

func (f *FloatField) asSql() string {
	return fieldAsSql(f)
}

func (f *IntegerField) asSql() string {
	return fieldAsSql(f)
}

func (f *SmallIntegerField) asSql() string {
	return fieldAsSql(f)
}

func (f *TextField) asSql() string {
	return fieldAsSql(f)
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

//Use the field's attribute name if it is part of a model
// func getFieldAlias(field field) string {
// 	if field.hasModel() {
// 		attrName, ok := field.getModel().colToAttr[field.getDbColumn()]
// 		if ok {
// 			return attrName
// 		}
// 	}
// 	return field.getDbColumn()
// }

func (f *BigIntegerField) Alias() string {
	return f.alias
}

func (f *BooleanField) Alias() string {
	return f.alias
}

func (f *FloatField) Alias() string {
	return f.alias
}

func (f *IntegerField) Alias() string {
	return f.alias
}

func (f *SmallIntegerField) Alias() string {
	return f.alias
}

func (f *TextField) Alias() string {
	return f.alias
}

func (a *aggregate) As(alias string) {
	a.outputField.As(alias)
}

func (f *function) As(alias string) {
	f.outputField.As(alias)
}

func (f *BigIntegerField) As(alias string) {
	f.alias = alias
}

func (f *BooleanField) As(alias string) {
	f.alias = alias
}

func (f *FloatField) As(alias string) {
	f.alias = alias
}

func (f *IntegerField) As(alias string) {
	f.alias = alias
}

func (f *SmallIntegerField) As(alias string) {
	f.alias = alias
}

func (f *TextField) As(alias string) {
	f.alias = alias
}

func (a aggregate) Scan(v interface{}) error {
	return function(a).Scan(v)
}

func (f function) Scan(v interface{}) error {
	return f.outputField.Scan(v)
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
	f.Value, f.valid = value.(float64)
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
//Methods that return Sort Expressions
//

func (f *BigIntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BooleanField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *FloatField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *IntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *SmallIntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *TextField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BigIntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BooleanField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *FloatField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *IntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *SmallIntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *TextField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}
