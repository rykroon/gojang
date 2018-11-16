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

type sqlValuer interface {
	asSqlValue() string
}

func newLookup(operator string, lhs filterer, rhs string) lookup {
	return lookup(fmt.Sprintf("%v %v %v", lhs.QualifiedName(), operator, rhs))
}

func newPatternLookup(lhs filterer, rhs sqlValuer, pattern string, caseSensitive bool) lookup {
	var value string
	strVal, ok := rhs.(stringValue)

	if ok {
		value = string(strVal)
	} else {
		value = rhs.asSqlValue()
	}

	newRhs := stringValue(fmt.Sprintf(pattern, value))

	var operator string

	if caseSensitive {
		operator = "LIKE"
	} else {
		operator = "ILIKE"
	}

	return newLookup(operator, lhs, newRhs.asSqlValue())
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

	//return newPatternLookup(lhs, rhs, "%%%v%%", true)
}

func icontains(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v%%", rhs)
	return newLookup("ILIKE", lhs, rhs)

	//return newPatternLookup(lhs, rhs, "%%%v%%", false)
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

	//return newPatternLookup(lhs, rhs, "%v%%", true)
}

func iStartsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%v%%", rhs)
	return newLookup("ILIKE", lhs, rhs)

	//return newPatternLookup(lhs, rhs, "%v%%", false)
}

func endsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v", rhs)
	return newLookup("LIKE", lhs, rhs)

	//return newPatternLookup(lhs, rhs, "%%%v", true)
}

func iEndsWith(lhs filterer, rhs string) lookup {
	rhs = fmt.Sprintf("%%%v", rhs)
	return newLookup("ILIKE", lhs, rhs)

	//return newPatternLookup(lhs, rhs, "%%%v", false)
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
