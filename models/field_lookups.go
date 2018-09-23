package models

import (
	"reflect"
)

type lookup struct {
	lhs        string
	lookupName string
	rhs        string
}

func (f field) Exact(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f field) IExact(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "ILIKE"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f field) Contains(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = "%" + value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f field) IContains(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "ILIKE"}
	value = "%" + value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f field) In(values ...interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "IN"}
	lookup.rhs = interfaceToSql(values)
	return lookup
}

func (f field) Gt(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: ">"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f field) Gte(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: ">="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f field) Lt(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "<"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f field) Lte(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "<="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f field) StartsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f field) IStartsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f field) EndsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = "%" + value
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f field) IEndsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = "%" + value
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f field) Range(from interface{}, to interface{}) lookup {
	if reflect.TypeOf(from) != reflect.TypeOf(to) {
		panic("Values have mismatching types")
	}

	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "BETWEEN"}
	lookup.rhs = interfaceToSql(from) + " AND " + interfaceToSql(to)
	return lookup
}

func (f field) IsNull(value bool) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}
