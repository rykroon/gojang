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
