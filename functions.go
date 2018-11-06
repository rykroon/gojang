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

//Comparison and conversion functions¶

func Cast(expr selectExpression, outputField field) *function {
	cast := newFunction("CAST", expr, outputField)
	cast.template = "%v(%v AS %v)"
	cast.args = append(cast.args, outputField.DbType())
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

func (f *TextField) Length() *IntegerField {
	length := Length(f)
	return length.outputField.(*IntegerField)
}

func (f *TextField) Upper() *TextField {
	upper := Upper(f)
	return upper.outputField.(*TextField)
}

func (f *TextField) Lower() *TextField {
	lower := Lower(f)
	return lower.outputField.(*TextField)
}
