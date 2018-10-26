package gojang

import (
	"fmt"
	"strings"
)

type function struct {
	name        string
	expr        expression
	args        []interface{}
	template    string
	outputField field
}

//type aggregate function

func (f function) asSql() string {
	result := fmt.Sprintf(f.template, f.args...)
	fmt.Println(result)
	return result
}

//Create a new function
func newFunc(name string, expr expression, outputField field) function {
	funct := function{}
	funct.name = strings.ToUpper(name)
	funct.expr = expr
	funct.args = []interface{}{funct.name, funct.expr.asSql()}
	funct.template = "%v(%v)"
	funct.outputField = outputField
	funct.outputField.setDbColumn(strings.ToLower(name))
	funct.outputField.setExpr(funct)
	return funct
}

//Creates a new AVG Function
func newAvg(expr expression) function {
	return newFunc("AVG", expr, NewFloatField())
}

//Creates a new CAST Function
func newCast(expr expression, outputField field) function {
	cast := newFunc("CAST", expr, outputField)
	cast.args = append(cast.args, cast.outputField.getDbType())
	cast.template = "%v(%v AS %v)"
	cast.outputField.setExpr(cast)
	return cast
}

//Creates a new COUNT Function
func newCount(expr expression, distinct bool) function {
	count := newFunc("COUNT", expr, NewIntegerField())

	if distinct {
		count.template = "%v(DISTINCT %v)"
	}

	count.outputField.setExpr(count)

	return count
}

//Creates a new SUM Function
func newSum(expr expression, outputField field) function {
	return newFunc("SUM", expr, outputField)
}

//Creates a new AVG function for a field
func avgField(field field) function {
	avg := newAvg(field)
	alias := field.getDbColumn() + "__avg"
	avg.outputField.setDbColumn(alias)

	if field.getDbType() != avg.outputField.getDbType() {
		cast := newCast(avg, avg.outputField)
		cast.outputField.setDbColumn(alias)
		return cast
	}

	return avg
}

func (f *AutoField) Avg() function {
	return avgField(f)
}

func (f *BigAutoField) Avg() function {
	return avgField(f)
}

func (f *BigIntegerField) Avg() function {
	return avgField(f)
}

func (f *FloatField) Avg() function {
	return avgField(f)
}

func (f *IntegerField) Avg() function {
	return avgField(f)
}

func (f *SmallIntegerField) Avg() function {
	return avgField(f)
}
