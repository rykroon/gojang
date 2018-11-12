package gojang

import (
	"strings"
)

type aggregate function

func newAggregate(name string, expr selectExpression, outputField field) *aggregate {
	f := newFunction(name, expr, outputField)
	f.As(expr.Alias() + "__" + strings.ToLower(name))
	a := aggregate(*f)
	return &a
}

func (a *aggregate) asSql() string {
	f := function(*a)
	return f.asSql()
}

func (a aggregate) Alias() string {
	return function(a).Alias()
}

func (a *aggregate) As(alias string) {
	a.outputField.As(alias)
}

func (a aggregate) DataType() string {
	return function(a).DataType()
}

func (a aggregate) getValue() interface{} {
	return function(a).getValue()
}

func (a aggregate) Scan(v interface{}) error {
	return function(a).Scan(v)
}

func Avg(expr selectExpression, outputField field) *aggregate {
	return newAggregate("AVG", expr, outputField)
}

func Count(expr selectExpression, distinct bool) *aggregate {
	count := newAggregate("COUNT", expr, NewBigIntegerField())

	if distinct {
		count.template = "%v(DISTINCT %v)"
	}

	return count
}

func Max(expr selectExpression, outputField field) *aggregate {
	return newAggregate("MAX", expr, outputField)
}

func Min(expr selectExpression, outputField field) *aggregate {
	return newAggregate("MIN", expr, outputField)
}

func Sum(expr selectExpression, outputField field) *aggregate {
	return newAggregate("SUM", expr, outputField)
}

func (s star) Count() *aggregate {
	count := Count(s, false)
	count.As("count")
	return count
}
