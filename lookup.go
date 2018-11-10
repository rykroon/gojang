package gojang

import (
	"database/sql/driver"
	"fmt"
)

type lookup string

type filterer interface {
	driver.Valuer
	QualifiedName() string
}

func newLookup(lookupName string, lhs filterer, rhs string) lookup {
	return lookup(fmt.Sprintf("%v %v %v", lhs.QualifiedName(), lookupName, rhs))
}

func exact(lhs filterer, rhs string) lookup {
	return newLookup("=", lhs, rhs)
}

func iexact(lhs filterer, rhs string) lookup {
	return newLookup("ILIKE", lhs, rhs)
}

func contains(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v%%", rhs)
	return newLookup("LIKE", lhs, rhs)
}

func icontains(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v%%", rhs)
	return newLookup("ILIKE", lhs, rhs)
}

func in(lhs filterer, rhs string) lookup {
	return newLookup("IN", lhs, rhs)
}

func gt(lhs filterer, rhs string) lookup {
	return newLookup(">", lhs, rhs)
}

func gte(lhs filterer, rhs string) lookup {
	return newLookup(">=", lhs, rhs)
}

func lt(lhs filterer, rhs string) lookup {
	return newLookup("<", lhs, rhs)
}

func lte(lhs filterer, rhs string) lookup {
	return newLookup("<=", lhs, rhs)
}

func startsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%v%%", rhs)
	return newLookup("LIKE", lhs, rhs)
}

func iStartsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%v%%", rhs)
	return newLookup("ILIKE", lhs, rhs)
}

func endsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v", rhs)
	return newLookup("LIKE", lhs, rhs)
}

func iEndsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v", rhs)
	return newLookup("ILIKE", lhs, rhs)
}

func between(lhs filterer, rhs1, rhs2 string) lookup {
	rhs := fmt.Sprintf("%v AND %v", rhs1, rhs2)
	return newLookup("BETWEEN", lhs, rhs)
}

func isNull(lhs filterer, null bool) lookup {
	rhs := "NULL"

	if !null {
		rhs = "NOT NULL"
	}

	return newLookup("IS", lhs, rhs)
}
