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

func (f *DecimalField) Avg() *aggregate {
	return Avg(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *IntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *SmallIntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (s star) Count() *aggregate {
	count := Count(s, false)
	count.As("count")
	return count
}

func (f *DecimalField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *IntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *SmallIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *TextField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *DecimalField) Max() *aggregate {
	return Max(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *IntegerField) Max() *aggregate {
	return Max(f, NewIntegerField())
}

func (f *SmallIntegerField) Max() *aggregate {
	return Max(f, NewSmallIntegerField())
}

func (f *TextField) Max() *aggregate {
	return Max(f, NewTextField())
}

func (f *DecimalField) Min() *aggregate {
	return Min(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *IntegerField) Min() *aggregate {
	return Min(f, NewIntegerField())
}

func (f *SmallIntegerField) Min() *aggregate {
	return Min(f, NewSmallIntegerField())
}

func (f *TextField) Min() *aggregate {
	return Min(f, NewTextField())
}

func (f *DecimalField) Sum() *aggregate {
	return Sum(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *IntegerField) Sum() *aggregate {
	return Sum(f, NewIntegerField())
}

func (f *SmallIntegerField) Sum() *aggregate {
	return Sum(f, NewSmallIntegerField())
}
