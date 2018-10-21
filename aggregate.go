package gojang

import ()

type aggregate struct {
	expression expression
	function   string
	alias      string
}

func (a aggregate) As(alias string) aggregate {
	a.alias = alias
	return a
}

func (f *BigIntegerField) Avg() aggregate {
	alias := f.dbColumn + "__avg"
	return aggregate{expression: f, function: "AVG", alias: alias}
}

func (f *FloatField) Avg() aggregate {
	alias := f.dbColumn + "__avg"
	return aggregate{expression: f, function: "AVG", alias: alias}
}

func (f *IntegerField) Avg() aggregate {
	alias := f.dbColumn + "__avg"
	return aggregate{expression: f, function: "AVG", alias: alias}
}

func (f *SmallIntegerField) Avg() aggregate {
	alias := f.dbColumn + "__avg"
	return aggregate{expression: f, function: "AVG", alias: alias}
}

func (f *BigIntegerField) Sum() aggregate {
	alias := f.dbColumn + "__sum"
	return aggregate{expression: f, function: "SUM", alias: alias}
}

func (f *FloatField) Sum() aggregate {
	alias := f.dbColumn + "__sum"
	return aggregate{expression: f, function: "SUM", alias: alias}
}

func (f *IntegerField) Sum() aggregate {
	alias := f.dbColumn + "__sum"
	return aggregate{expression: f, function: "SUM", alias: alias}
}

func (f *SmallIntegerField) Sum() aggregate {
	alias := f.dbColumn + "__sum"
	return aggregate{expression: f, function: "SUM", alias: alias}
}

func Avg(expr expression) aggregate {
	numField, isNumField := expr.(numericField)
	if isNumField {
		return numField.Avg()
	}

	a := aggregate{expression: expr, function: "AVG", alias: "avg"}

	return a
}

func Count(expr expression) aggregate {
	a := aggregate{expression: expr, function: "COUNT", alias: "count"}

	field, isAField := expr.(field)
	if isAField {
		a.alias = field.getDbColumn() + "__count"
	}

	return a
}

func Min(expr expression) aggregate {
	a := aggregate{expression: expr, function: "MIN", alias: "min"}

	field, isAField := expr.(field)
	if isAField {
		a.alias = field.getDbColumn() + "__min"
	}

	return a
}

func Max(expr expression) aggregate {
	a := aggregate{expression: expr, function: "MAX", alias: "max"}

	field, isAField := expr.(field)
	if isAField {
		a.alias = field.getDbColumn() + "__max"
	}

	return a
}

func Sum(expr expression) aggregate {
	numField, isNumField := expr.(numericField)
	if isNumField {
		return numField.Sum()
	}

	a := aggregate{expression: expr, function: "SUM", alias: "sum"}

	return a
}
