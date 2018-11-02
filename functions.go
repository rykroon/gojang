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
	funct.args = []interface{}{funct.name, expr.asSql()}
	funct.outputField = outputField
	funct.outputField.setDbColumn(funct.asSql())
	funct.As(strings.ToLower(name))
	return funct
}

func newAggregate(name string, expr selectExpression, outputField field) *aggregate {
	f := newFunction(name, expr, outputField)
	a := &aggregate{}
	a = *f
	return a
}

func Avg(expr selectExpression) *aggregate {
	agg := NewAggregate("AVG", expr, NewFloatField())
	agg.As(expr.Alias() + "__avg")
	return agg
}

func Count(expr selectExpression, distinct bool) *aggregate {
	agg := NewAggregate("COUNT", expr, NewBigIntegerField())

	if distinct {
		agg.template = "%v(DISTINCT %v)"
		agg.outputField.setDbColumn(agg.asSql())
	}

	agg.As(expr.Alias() + "__count")
	return agg
}


//Creates a new CAST Function
func newCast(expr expression, outputField field) *function {
	cast := newFunc("CAST", expr, outputField)
	cast.args = append(cast.args, cast.outputField.DbType())
	cast.template = "%v(%v AS %v)"
	return cast
}

//Creates a new AVG function for a field
func avgField(field field) *aggregate {
	avg := newAvg(field)
	alias := field.DbColumn() + "__avg"
	avg.As(alias)

	if field.DbType() != avg.outputField.DbType() {
		cast := newCast(avg, avg.outputField)
		cast.As(alias)
		aggCast := aggregate(*cast)
		return &aggCast
	}

	return avg
}

// //Creates a new Count function for a field
// func countField(field field, distinct bool) *aggregate {
// 	count := newCount(field, distinct)
// 	alias := field.DbColumn() + "__count"
// 	count.As(alias)
// 	return count
// }

// func sumField(field field) function {
//   sum := newSum(field)
//   alias := field.DbColumn() + "__sum"
//   //sum.outputField.setDbColumn(alias)
// 	sum = sum.As(alias)
//   return sum
// }

// func (f *AutoField) Avg() *aggregate {
// 	return avgField(f)
// }
//
// func (f *BigAutoField) Avg() *aggregate {
// 	return avgField(f)
// }

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
