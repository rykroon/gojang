package gojang

import ()

type aggregate struct {
	//field    field
	expression string
	function   string
	alias      string
}

func (a aggregate) As(alias string) aggregate {
	a.alias = alias
	return a
}

//use Count aggregate in queryset count
// func Count() aggregate {
//
// }

//maybe create a numericField interface?
// func Avg(numField numericField) aggregate {
// 	field := numField.(field)
// 	name := field.getDbColumn() + "__avg"
// 	a := aggregate{field: field, function: "AVG", name: name}
// }
