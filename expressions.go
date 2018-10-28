package gojang

import (
	"fmt"
	//"reflect"
	//"unsafe"
)

//Expressions describe a value or a computation that can be used as part of an
//update, create, filter, order by, annotation, or aggregate.
type expression interface {
	asSql() string
}

type selectExpression interface {
	expression
	Scan(interface{}) error
	getValue() interface{}
	getAlias() string
}

// type annotation interface {
// 	expression
// 	As()
// }

type sortExpression struct {
	field field
	desc  bool
}

type star string

func (a aggregate) asSql() string {
	return function(a).asSql()
}

func (f function) asSql() string {
	return fmt.Sprintf(f.template, f.args...)
}

func (l lookup) asSql() string {
	sql := l.lhs.asSql() + " " + l.lookupName + " " + l.rhs

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
	if field.hasModel() {
		return dbq(field.getModel().dbTable) + "." + dbq(field.getDbColumn())
	} else {
		return field.getExpr().asSql() + " AS " + dbq(field.getDbColumn())
	}
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

func (a aggregate) Scan(v interface{}) error {
	return function(a).Scan(v)
}

//
//Select Expression Method Set
//

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

func (a aggregate) getAlias() string {
	return function(a).getAlias()
}

func (f function) getAlias() string {
	return f.outputField.getDbColumn()
}

//Use the field's attribute name if it is part of a model
func getFieldAlias(field field) string {
	if field.hasModel() {
		attrName, ok := field.getModel().colToAttr[field.getDbColumn()]
		if ok {
			return attrName
		}
	}
	return field.getDbColumn()
}

func (f *BigIntegerField) getAlias() string {
	return getFieldAlias(f)
}

func (f *BooleanField) getAlias() string {
	return getFieldAlias(f)
}

func (f *FloatField) getAlias() string {
	return getFieldAlias(f)
}

func (f *IntegerField) getAlias() string {
	return getFieldAlias(f)
}

func (f *SmallIntegerField) getAlias() string {
	return getFieldAlias(f)
}

func (f *TextField) getAlias() string {
	return getFieldAlias(f)
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
