package gojang

import (
	"fmt"
)

// type lookup struct {
// 	not        bool
// 	lhs        field
// 	lookupName string
// 	rhs        string
// }

type lookup string

func newLookup(lookupName string, lhs expression, rhs string) lookup {
	return lookup(fmt.Sprintf("%v %v %v", lhs.asSql(), lookupName, rhs))
}

func exact(lhs expression, rhs string) lookup {
	return newLookup("=", lhs, rhs)
}

func iexact(lhs expression, rhs string) lookup {
	return newLookup("ILIKE", lhs, rhs)
}

func contains(lhs expression, rhs string) lookup {
	rhs = fmt.Sprintf("%%v%", rhs)
	return newLookup("LIKE", lhs, rhs)
}

func icontains(lhs expression, rhs string) lookup {
	rhs = fmt.Sprintf("%%v%", rhs)
	return newLookup("ILIKE", lhs, rhs)
}

func in(lhs expression, rhs string) lookup {
	return newLookup("IN", lhs, rhs)
}

func gt(lhs expression, rhs string) lookup {
	return newLookup(">", lhs, rhs)
}

func gte(lhs expression, rhs string) lookup {
	return newLookup(">=", lhs, rhs)
}

func lt(lhs expression, rhs string) lookup {
	return newLookup("<", lhs, rhs)
}

func lte(lhs expression, rhs string) lookup {
	return newLookup("<=", lhs, rhs)
}

func startsWith(lhs expression, rhs string) lookup {
	rhs = fmt.Sprintf("%v%", rhs)
	return newLookup("LIKE", lhs, rhs)
}

func istartsWith(lhs expression, rhs string) lookup {
	rhs = fmt.Sprintf("%v%", rhs)
	return newLookup("ILIKE", lhs, rhs)
}

func endsWith(lhs expression, rhs string) lookup {
	rhs = fmt.Sprintf("%%v", rhs)
	return newLookup("LIKE", lhs, rhs)
}

func iendsWith(lhs expression, rhs string) lookup {
	rhs = fmt.Sprintf("%%v", rhs)
	return newLookup("ILIKE", lhs, rhs)
}

func between(lhs expression, rhs1, rhs2 string) lookup {
	rhs := fmt.Sprintf("%v AND %v", rhs1, rhs2)
	return newLookup("BETWEEN", lhs, rhs)
}

func isNull(lhs expression, null bool) lookup {
	rhs := "NULL"

	if !null {
		rhs = "NOT NULL"
	}

	return newLookup("IS", lhs, rhs)
}

// func exactIntField(field IntField, value int) lookup {
// 	return lookup{lhs: field, lookupName: "=", rhs: intAsSql(value)}
// }
//
// func inIntField(field IntField, values []int) lookup {
// 	return lookup{lhs: field, lookupName: "IN", rhs: integersAsSql(values)}
// }
//
// func gtIntField(field IntField, value int) lookup {
// 	return lookup{lhs: field, lookupName: ">", rhs: intAsSql(value)}
// }
//
// func gteIntField(field IntField, value int) lookup {
// 	return lookup{lhs: field, lookupName: ">=", rhs: intAsSql(value)}
// }
//
// func ltIntField(field IntField, value int) lookup {
// 	return lookup{lhs: field, lookupName: "<", rhs: intAsSql(value)}
// }
//
// func lteIntField(field IntField, value int) lookup {
// 	return lookup{lhs: field, lookupName: "<=", rhs: intAsSql(value)}
// }
//
// func rangeIntField(field IntField, from, to int) lookup {
// 	lookup := lookup{lhs: field, lookupName: "BETWEEN"}
// 	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
// 	return lookup
// }
//
// func fieldIsNull(field field, value bool) lookup {
// 	lookup := lookup{lhs: field, lookupName: "IS"}
//
// 	if value {
// 		lookup.rhs = "NULL"
// 	} else {
// 		lookup.rhs = "NOT NULL"
// 	}
//
// 	return lookup
// }
