package models

import (
	"strconv"
)

type lookup struct {
	lhs        string
	lookupName string
	rhs        string
}

func (l lookup) toSql() string {
	return l.lhs + " " + l.lookupName + " " + l.rhs
}

func ExactNum(field string, value float64) lookup {
	lookup := lookup{lhs: field, lookupName: "="}
	lookup.rhs = strconv.FormatFloat(value, 'f', -1, 64)
	return lookup
}

func ExactChar(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "="}
	lookup.rhs = "'" + value + "'"
	return lookup
}

func Contains(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	lookup.rhs = "'%" + value + "'%"
	return lookup
}

func IContains(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "ILIKE"}
	lookup.rhs = "'%" + value + "'%"
	return lookup
}

func InNum(field string, values ...int) lookup {
	return lookup{}
}

func InChar(field string, values ...string) lookup {
	return lookup{}
}

func InQs(field string, qs QuerySet) lookup {
	return lookup{}
}

func Gt(field string, value int) lookup {
	lookup := lookup{lhs: field, lookupName: ">"}
	lookup.rhs = strconv.Itoa(value)
	return lookup
}

func Gte(field string, value int) lookup {
	lookup := lookup{lhs: field, lookupName: ">="}
	lookup.rhs = strconv.Itoa(value)
	return lookup
}

func Lt(field string, value int) lookup {
	lookup := lookup{lhs: field, lookupName: "<"}
	lookup.rhs = strconv.Itoa(value)
	return lookup
}

func Lte(field string, value int) lookup {
	lookup := lookup{lhs: field, lookupName: "<="}
	lookup.rhs = strconv.Itoa(value)
	return lookup
}

func StartsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	lookup.rhs = "'" + value + "'%"
	return lookup
}

func IStartsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	lookup.rhs = "'" + value + "'%"
	return lookup
}

func EndsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	lookup.rhs = "'%" + value + "'"
	return lookup
}

func IEndsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	lookup.rhs = "'%" + value + "'"
	return lookup
}

func IsNull(field string, value bool) lookup {
	lookup := lookup{lhs: field, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}
