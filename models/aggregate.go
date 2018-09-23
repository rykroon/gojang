package models

import ()

type aggregate struct {
	field    field
	function string
	name     string
}

func (a aggregate) asSql() string {
	//name := doubleQuotes(a.field.dbColumn + "__" + strings.ToLower(a.function))
	name := doubleQuotes(a.name)
	return a.function + "(" + a.field.toSql() + ") as " + name
}

func (f field) Avg() aggregate {
	//Add check to make sure that the field is capable of doing AVG()
	name := f.dbColumn + "__avg"
	return aggregate{field: f, function: "AVG", name: name}
}

func (f field) Count() aggregate {
	name := f.dbColumn + "__count"
	return aggregate{field: f, function: "COUNT", name: name}
}

func (f field) Max() aggregate {
	name := f.dbColumn + "__max"
	return aggregate{field: f, function: "MAX", name: name}
}

func (f field) Min() aggregate {
	name := f.dbColumn + "__min"
	return aggregate{field: f, function: "MIN", name: name}
}

func (f field) Sum() aggregate {
	name := f.dbColumn + "__sum"
	return aggregate{field: f, function: "SUM", name: name}
}
