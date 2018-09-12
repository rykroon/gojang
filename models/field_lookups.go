package models

import (
//"strconv"
	//"strings"
)

type otherLookup struct {
	
}

type lookup struct {
	lhs        string
	lookupName string
	rhs        string
}

func (l lookup) toSql() string {
	return l.lhs + " " + l.lookupName + " " + l.rhs
}

func lookupToOperator(lookup string) string {
	switch lookup {
	case "exact":
		return "="

	case "iexact":
		return "LIKE"

	case "contains":
		return "LIKE"

	case "icontains":
		return "ILIKE"

	case "in":
		return "IN"

	case "gt":
		return ">"

	case "gte":
		return "<="

	case "lt":
		return "<"

	case "lte":
		return "<="

	case "startswith":
		return "LIKE"

	case "istartswith":
		return "ILIKE"

	case "endswith":
		return "LIKE"

	case "iendswith":
		return "ILIKE"

	case "isnull":
		return "IS"
	}

	return ""

}




func Exact(field string, value interface{}) lookup {
	lookup := lookup{lhs: field, lookupName: "="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func Contains(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	value = "%" + value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func IContains(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "ILIKE"}
	value = "%" + value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func In(field string, values ...interface{}) lookup {
	lookup := lookup{lhs: field, lookupName: "IN"}
	lookup.rhs = interfaceToSql(values)
	return lookup
}

// func InInt(field string, values ...int) lookup {
// 	return lookup{}
// }
//
// func InString(field string, values ...string) lookup {
// 	return lookup{}
// }
//
// func InQs(field string, qs QuerySet) lookup {
// 	return lookup{}
// }

func Gt(field string, value interface{}) lookup {
	lookup := lookup{lhs: field, lookupName: ">"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func Gte(field string, value interface{}) lookup {
	lookup := lookup{lhs: field, lookupName: ">="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func Lt(field string, value interface{}) lookup {
	lookup := lookup{lhs: field, lookupName: "<"}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func Lte(field string, value interface{}) lookup {
	lookup := lookup{lhs: field, lookupName: "<="}
	lookup.rhs = interfaceToSql(value)
	return lookup
}

func StartsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	value = value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func IStartsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	value = value + "%"
	lookup.rhs = stringToSql(value)
	return lookup
}

func EndsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	value = "%" + value
	lookup.rhs = stringToSql(value)
	return lookup
}

func IEndsWith(field string, value string) lookup {
	lookup := lookup{lhs: field, lookupName: "LIKE"}
	value = "%" + value
	lookup.rhs = stringToSql(value)
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
