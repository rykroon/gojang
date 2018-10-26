package gojang

import (
	"strings"
)

type aggregate struct {
	function    string
	expression  expression
	alias       string
	template    string
	outputField field
}

func newAggregate(function string, expr expression, outputField field) aggregate {
	a := aggregate{}
	a.function = strings.ToUpper(function)
	a.expression = expr
	a.outputField = outputField
	a.alias = strings.ToLower(function)
	a.template = "%v(%v) AS \"%v\""
	return a
}

// func newAvg(expr expression) aggregate {
// 	a := newAggregate("avg", expr, NewFloatField())
// 	a.template = "%v(%v)::float AS \"%v\""
// 	return a
// }

// func newCount(expr expression, distinct bool) aggregate {
// 	a := newAggregate("count", expr, NewIntegerField())
//
// 	if distinct {
// 		a.template = "%v(DISTINCT %v) AS %v"
// 	}
//
// 	return a
// }

func newMax(expr expression, outputField field) aggregate {
	return newAggregate("max", expr, outputField)
}

func newMin(expr expression, outputField field) aggregate {
	return newAggregate("min", expr, outputField)
}

// func newSum(expr expression, outputField field) aggregate {
// 	return newAggregate("sum", expr, outputField)
// }

func (a aggregate) As(alias string) aggregate {
	a.alias = alias
	return a
}

// func (f *AutoField) Avg() aggregate {
// 	a := newAvg(f)
// 	a.alias = f.dbColumn + "__avg"
// 	return a
// }
//
// func (f *BigAutoField) Avg() aggregate {
// 	a := newAvg(f)
// 	a.alias = f.dbColumn + "__avg"
// 	return a
// }
//
// func (f *BigIntegerField) Avg() aggregate {
// 	a := newAvg(f)
// 	a.alias = f.dbColumn + "__avg"
// 	return a
// }
//
// func (f *FloatField) Avg() aggregate {
// 	a := newAvg(f)
// 	a.alias = f.dbColumn + "__avg"
// 	return a
// }
//
// func (f *IntegerField) Avg() aggregate {
// 	a := newAvg(f)
// 	a.alias = f.dbColumn + "__avg"
// 	return a
// }
//
// func (f *SmallIntegerField) Avg() aggregate {
// 	a := newAvg(f)
// 	a.alias = f.dbColumn + "__avg"
// 	return a
// }

func (s star) Count() aggregate {
	return aggregate{}
	//return newCount(s, false)
}

//
// func (f *AutoField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *BigAutoField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *BigIntegerField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *BooleanField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *FloatField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *IntegerField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *SmallIntegerField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *TextField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *ForeignKey) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *OneToOneField) Count(distinct bool) aggregate {
// 	a := newCount(f, distinct)
// 	a.alias = f.dbColumn + "__count"
// 	return a
// }
//
// func (f *AutoField) Sum() aggregate {
// 	a := newSum(f, NewBigIntegerField())
// 	a.alias = f.dbColumn + "__sum"
// 	return a
// }
//
// func (f *BigAutoField) Sum() aggregate {
// 	a := newSum(f, NewBigIntegerField())
// 	a.alias = f.dbColumn + "__sum"
// 	return a
// }
//
// func (f *BigIntegerField) Sum() aggregate {
// 	a := newSum(f, NewBigIntegerField())
// 	a.alias = f.dbColumn + "__sum"
// 	return a
// }
//
// func (f *FloatField) Sum() aggregate {
// 	a := newSum(f, NewFloatField())
// 	a.alias = f.dbColumn + "__sum"
// 	//a.template = "%v(%v)::float AS \"%v\""
// 	return a
// }
//
// func (f *IntegerField) Sum() aggregate {
// 	a := newSum(f, NewIntegerField())
// 	a.alias = f.dbColumn + "__sum"
// 	return a
// }
//
// func (f *SmallIntegerField) Sum() aggregate {
// 	a := newSum(f, NewSmallIntegerField())
// 	a.alias = f.dbColumn + "__sum"
// 	return a
// }
