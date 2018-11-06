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

type aggregate function

func newFunction(name string, expr selectExpression, outputField field) *function {
	funct := &function{}
	funct.name = strings.ToUpper(name)
	funct.template = "%v(%v)"
	funct.args = []interface{}{funct.name, expr.asSql()}
	funct.outputField = outputField
	return funct
}

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

func (f *BigIntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *DecimalField) Avg() *aggregate {
	return Avg(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *FloatField) Avg() *aggregate {
	return Avg(f, NewFloatField())
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

func (f *BigIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *BooleanField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *DecimalField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *FloatField) Count(distinct bool) *aggregate {
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

func (f *BigIntegerField) Max() *aggregate {
	return Max(f, NewBigIntegerField())
}

func (f *BooleanField) Max() *aggregate {
	return Max(f, NewBooleanField())
}

func (f *DecimalField) Max() *aggregate {
	return Max(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *FloatField) Max() *aggregate {
	return Max(f, NewFloatField())
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

func (f *BigIntegerField) Min() *aggregate {
	return Min(f, NewBigIntegerField())
}

func (f *BooleanField) Min() *aggregate {
	return Min(f, NewBooleanField())
}

func (f *DecimalField) Min() *aggregate {
	return Min(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *FloatField) Min() *aggregate {
	return Min(f, NewFloatField())
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

func (f *BigIntegerField) Sum() *aggregate {
	return Sum(f, NewBigIntegerField())
}

func (f *DecimalField) Sum() *aggregate {
	return Sum(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *FloatField) Sum() *aggregate {
	return Sum(f, NewFloatField())
}

func (f *IntegerField) Sum() *aggregate {
	return Sum(f, NewIntegerField())
}

func (f *SmallIntegerField) Sum() *aggregate {
	return Sum(f, NewSmallIntegerField())
}
