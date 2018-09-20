package models

import (
	"reflect"
)

type lookup struct {
	lhs        string
	lookupName string
	rhs        string
}

func (f Field) Exact(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f Field) IExact(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "ILIKE"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f Field) Contains(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = "%" + value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f Field) IContains(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "ILIKE"}
	value = "%" + value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f Field) In(values ...interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "IN"}
	lookup.rhs = interfaceToSql(values)
	return lookup
}

func (f Field) Gt(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: ">"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f Field) Gte(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: ">="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f Field) Lt(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "<"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f Field) Lte(value interface{}) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "<="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func (f Field) StartsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f Field) IStartsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f Field) EndsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = "%" + value
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f Field) IEndsWith(value string) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "LIKE"}
	value = "%" + value
	lookup.rhs = stringToSql(value)
	return lookup
}

func (f Field) Range(from interface{}, to interface{}) lookup {
	if reflect.TypeOf(from) != reflect.TypeOf(to) {
		panic("Values have mismatching types")
	}

	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "BETWEEN"}
	lookup.rhs = interfaceToSql(from) + " AND " + interfaceToSql(to)
	return lookup
}

func (f Field) IsNull(value bool) lookup {
	fieldName := f.toSql()
	lookup := lookup{lhs: fieldName, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}
