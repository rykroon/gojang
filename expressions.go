package gojang

import (
	"fmt"
	"reflect"
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
	getGoType() string
}

type sortExpression struct {
	field field
	desc  bool
}

type star string

func (s star) asSql() string {
	return "*"
}

func (a aggregate) asSql() string {
	return fmt.Sprintf(a.template, a.function, a.expression.asSql(), a.alias)
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

func fieldAsSql(field field) string {
	if field.hasModel() {
		return dbq(field.getModel().dbTable) + "." + dbq(field.getDbColumn())
	} else {
		return field.getExpr().asSql() + " AS " + dbq(field.getDbColumn())
	}
}

func (f *AutoField) asSql() string {
	return fieldAsSql(f)
}

func (f *BigAutoField) asSql() string {
	return fieldAsSql(f)
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

func (f *ForeignKey) asSql() string {
	return fieldAsSql(f)
}

func (f *OneToOneField) asSql() string {
	return fieldAsSql(f)
}

func (a aggregate) Scan(value interface{}) error {
	return a.outputField.Scan(value)
}

func (f *AutoField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.valid = int32(result), ok
	return nil
}

func (f *BigAutoField) Scan(value interface{}) error {
	f.Value, f.valid = value.(int64)
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

func (f *ForeignKey) Scan(value interface{}) error {
	f.Value, f.valid = value.(int64)
	return nil
}

func (f *OneToOneField) Scan(value interface{}) error {
	f.Value, f.valid = value.(int64)
	return nil
}

func (a aggregate) getValue() interface{} {
	return a.outputField.getValue()
}

func (f *AutoField) getValue() interface{} {
	return f.Value
}

func (f *BigAutoField) getValue() interface{} {
	return f.Value
}

func (f *BigIntegerField) getValue() interface{} {
	return f.Value
}

func (f *BooleanField) getValue() interface{} {
	return f.Value
}

func (f *FloatField) getValue() interface{} {
	return f.Value
}

func (f *IntegerField) getValue() interface{} {
	return f.Value
}

func (f *SmallIntegerField) getValue() interface{} {
	return f.Value
}

func (f *TextField) getValue() interface{} {
	return f.Value
}

func (f *ForeignKey) getValue() interface{} {
	return f.Value
}

func (f *OneToOneField) getValue() interface{} {
	return f.Value
}

func (a aggregate) getGoType() string {
	return a.outputField.getGoType()
}

func (f *AutoField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *BigAutoField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *BigIntegerField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *BooleanField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *FloatField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *IntegerField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *SmallIntegerField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *TextField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *ForeignKey) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *OneToOneField) getGoType() string {
	return reflect.TypeOf(f.Value).String()
}

func (f *AutoField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BigAutoField) Asc() sortExpression {
	return sortExpression{field: f}
}

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

func (f *ForeignKey) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *OneToOneField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *AutoField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BigAutoField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
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

func (f *ForeignKey) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *OneToOneField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}
