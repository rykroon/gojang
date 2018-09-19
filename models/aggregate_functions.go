package models

import (
	"strings"
)

type aggregate struct {
	field    Field
	function string
	//name string
}

func (a aggregate) asSql() string {
	name := doubleQuotes(a.field.dbColumn + "__" + strings.ToLower(a.function))
	return a.function + "(" + a.field.toSql() + ") as " + name
}

func (f Field) Avg() aggregate {
	return aggregate{field: f, function: "AVG"}
}

func (f Field) Count() aggregate {
	return aggregate{field: f, function: "COUNT"}
}

func (f Field) Max() aggregate {
	return aggregate{field: f, function: "MAX"}
}

func (f Field) Min() aggregate {
	return aggregate{field: f, function: "MIN"}
}

func (f Field) Sum() aggregate {
	return aggregate{field: f, function: "SUM"}
}
