package gojang

import (
	"strings"
)

type function struct {
	name        string
	expr        expression
	args        []interface{}
	template    string
	outputField field
}

type aggregate function

//Create a new function
func newFunc(name string, expr expression, outputField field) *function {
	funct := &function{}
	funct.name = strings.ToUpper(name)
	funct.expr = expr
	funct.args = []interface{}{funct.name, funct.expr.asSql()}
	funct.template = "%v(%v)"

	funct.outputField = outputField
	funct.As(strings.ToLower(name))
	return funct
}

func newAgg(name string, expr expression, outputField field) *aggregate {
	f := newFunc(name, expr, outputField)
	a := aggregate(*f)
	return &a
}

//Creates a new AVG Function
func newAvg(expr expression) *aggregate {
	return newAgg("AVG", expr, NewFloatField())
}

//Creates a new CAST Function
func newCast(expr expression, outputField field) *function {
	cast := newFunc("CAST", expr, outputField)
	cast.args = append(cast.args, cast.outputField.DbType())
	cast.template = "%v(%v AS %v)"
	return cast
}

//Creates a new COUNT Function
func newCount(expr expression, distinct bool) *aggregate {
	count := newAgg("COUNT", expr, NewIntegerField())

	if distinct {
		count.template = "%v(DISTINCT %v)"
	}

	return count
}

//Creates a new SUM Function
func newSum(expr expression, outputField field) *function {
	return newFunc("SUM", expr, outputField)
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

//Creates a new Count function for a field
func countField(field field, distinct bool) *aggregate {
	count := newCount(field, distinct)
	alias := field.DbColumn() + "__count"
	count.As(alias)
	return count
}

// func sumField(field field) function {
//   sum := newSum(field)
//   alias := field.DbColumn() + "__sum"
//   //sum.outputField.setDbColumn(alias)
// 	sum = sum.As(alias)
//   return sum
// }

func (f *AutoField) Avg() *aggregate {
	return avgField(f)
}

func (f *BigAutoField) Avg() *aggregate {
	return avgField(f)
}

func (f *BigIntegerField) Avg() *aggregate {
	return avgField(f)
}

func (f *FloatField) Avg() *aggregate {
	return avgField(f)
}

func (f *IntegerField) Avg() *aggregate {
	return avgField(f)
}

func (f *SmallIntegerField) Avg() *aggregate {
	return avgField(f)
}

func (s star) Count() *aggregate {
	return newCount(s, false)
}

func (f *BigIntegerField) Count(distinct bool) *aggregate {
	return countField(f, distinct)
}

func (f *BooleanField) Count(distinct bool) *aggregate {
	return countField(f, distinct)
}

func (f *FloatField) Count(distinct bool) *aggregate {
	return countField(f, distinct)
}

func (f *IntegerField) Count(distinct bool) *aggregate {
	return countField(f, distinct)
}

func (f *SmallIntegerField) Count(distinct bool) *aggregate {
	return countField(f, distinct)
}

func (f *TextField) Count(distinct bool) *aggregate {
	return countField(f, distinct)
}
