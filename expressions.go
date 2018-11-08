package gojang

import (
	"fmt"
	//"reflect"
	//"github.com/shopspring/decimal"
	//"strconv"
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
	DataType() string
	Scan(interface{}) error
}

type orderByExpression string

type star string

func (a *aggregate) asSql() string {
	f := function(*a)
	return f.asSql()
}

func (a assignment) asSql() string {
	return dbq(a.lhs.ColumnName()) + " " + a.lookupName + " " + a.rhs
}

func (f *function) asSql() string {
	return fmt.Sprintf(f.template, f.args...)
}

// func (l lookup) asSql() string {
// 	sql := l.lhs.asSql() + " " + l.lookupName + " " + l.rhs
//
// 	if l.not {
// 		sql = "NOT(" + sql + ")"
// 	}
// 	return sql
// }

func (s star) asSql() string {
	return "*"
}

func (v *ValueExpression) asSql() string {
	return v.outputField.valueAsSql()
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

func (s star) Alias() string {
	return "*"
}

func (v *ValueExpression) Alias() string {
	return v.outputField.Alias()
}

func (a *aggregate) As(alias string) {
	a.outputField.As(alias)
}

func (f *function) As(alias string) {
	f.outputField.As(alias)
}

func (s star) As(string) {
	return
}

func (v *ValueExpression) As(alias string) {
	v.outputField.As(alias)
}

func (a aggregate) DataType() string {
	return function(a).DataType()
}

func (f function) DataType() string {
	return f.outputField.DataType()
}

func (s star) DataType() string {
	return ""
}

func (v *ValueExpression) DataType() string {
	return v.outputField.DataType()
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

func (v *ValueExpression) Scan(value interface{}) error {
	return v.outputField.Scan(value)
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

func (v *ValueExpression) getValue() interface{} {
	return v.outputField.getValue()
}

//
//Methods that return Order By Expressions
//
