package gojang

import ()

type aggregate struct {
	field    field
	function string
	name     string
}

func (a aggregate) toSql() string {
	return a.function + "(" + a.field.toSql() + ") AS " + a.name
}

func (a aggregate) As(name string) aggregate {
	a.name = name
	return a
}

//use Count aggregate in queryset count
func Count(field field, distinct bool) aggregate {
	name := field.getDbColumn() + "__count"
	a := aggregate{field: field, function:"COUNT", name: name}
	return a
}

//maybe create a numericField interface?
// func Avg(numField numericField) aggregate {
// 	field := numField.(field)
// 	name := field.getDbColumn() + "__avg"
// 	a := aggregate{field: field, function: "AVG", name: name}
// }
