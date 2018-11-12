package gojang

import (
	"strings"
)

type function struct {
	name        string
	args        []interface{}
	template    string
	outputField field
}

func newFunction(name string, expr selectExpression, outputField field) *function {
	funct := &function{}
	funct.name = strings.ToUpper(name)
	funct.template = "%v(%v)"
	funct.args = []interface{}{funct.name, expr.asSql()}
	funct.outputField = outputField
	return funct
}

func (f function) Alias() string {
	return f.outputField.Alias()
}

func (f *function) As(alias string) {
	f.outputField.As(alias)
}

func (f function) DataType() string {
	return f.outputField.DataType()
}

func (f function) getValue() interface{} {
	return f.outputField.getValue()
}

func (f function) Scan(v interface{}) error {
	return f.outputField.Scan(v)
}

func (f *function) toField() field {
	f.outputField.setColumnName(f.asSql())
	return f.outputField
}

//Comparison and conversion functions¶

func Cast(expr selectExpression, outputField field) *function {
	cast := newFunction("CAST", expr, outputField)
	cast.template = "%v(%v AS %v)"
	cast.args = append(cast.args, outputField.DataType())
	return cast
}

//Coalesce

//Date Functions

//Extract
//Now

//Text Functions

func Length(expr selectExpression) *function {
	return newFunction("LENGTH", expr, NewIntegerField())
}

func Upper(expr selectExpression) *function {
	return newFunction("UPPER", expr, NewTextField())
}

func Lower(expr selectExpression) *function {
	return newFunction("LOWER", expr, NewTextField())
}
