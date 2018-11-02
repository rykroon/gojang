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

//Creates a new CAST Function
func Cast(expr selectExpression, outputField field) *function {
	cast := newFunction("CAST", expr, outputField)
	cast.args = append(cast.args, cast.outputField.DbType())
	cast.template = "%v(%v AS %v)"
	return cast
}

func newAggregate(name string, expr selectExpression, outputField field) *aggregate {
	f := newFunction(name, expr, outputField)
	f.As(expr.Alias() + "__" + strings.ToLower(name))
	a := aggregate(*f)
	return &a
}

func Avg(expr selectExpression) *aggregate {
	avg := newAggregate("AVG", expr, NewFloatField())

	//Once I create Decimal Field use that for certain cases
	if expr.DbType() != avg.outputField.DbType() {
		cast := Cast(avg, avg.outputField)
		aggCast := aggregate(*cast)
		return &aggCast
	}

	return avg
}

func Count(expr selectExpression, distinct bool) *aggregate {
	agg := newAggregate("COUNT", expr, NewBigIntegerField())

	if distinct {
		agg.template = "%v(DISTINCT %v)"
	}

	if expr.Alias() != "*" {
		agg.As(expr.Alias() + "__count")
	}

	return agg
}

// func Min(expr selectExpression) *aggregate {
// 	//min := newAggregate("MIN", expr, )
// }

func Sum(expr selectExpression) *aggregate {
	sum := newAggregate("SUM", expr, NewBigIntegerField())
	sum.As(expr.Alias() + "__sum")
	return sum
}

func (f *BigIntegerField) Avg() *aggregate {
	return Avg(f)
}

func (f *FloatField) Avg() *aggregate {
	return Avg(f)
}

func (f *IntegerField) Avg() *aggregate {
	return Avg(f)
}

func (f *SmallIntegerField) Avg() *aggregate {
	return Avg(f)
}

func (s star) Count() *aggregate {
	return Count(s, false)
}

func (f *BigIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *BooleanField) Count(distinct bool) *aggregate {
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

func (f *BigIntegerField) Sum() *aggregate {
	return Sum(f)
}

func (f *FloatField) Sum() *aggregate {
	return Sum(f)
}

func (f *IntegerField) Sum() *aggregate {
	return Sum(f)
}

func (f *SmallIntegerField) Sum() *aggregate {
	return Sum(f)
}
