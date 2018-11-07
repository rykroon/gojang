package gojang

import (
	"github.com/shopspring/decimal"
)

type lookup struct {
	not        bool
	lhs        field
	lookupName string
	rhs        string
}

func exactIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "=", rhs: intAsSql(value)}
}

func (f *DecimalField) Exact(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: f.Value.String()}
}

func inIntField(field IntField, values []int) lookup {
	return lookup{lhs: field, lookupName: "IN", rhs: integersAsSql(values)}
}

func gtIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: ">", rhs: intAsSql(value)}
}

func (f *DecimalField) Gt(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: f.Value.String()}
}

func gteIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: ">=", rhs: intAsSql(value)}
}

func (f *DecimalField) Gte(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: f.Value.String()}
}

func ltIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "<", rhs: intAsSql(value)}
}

func (f *DecimalField) Lt(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: f.Value.String()}
}

func lteIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "<=", rhs: intAsSql(value)}
}

func (f *DecimalField) Lte(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: f.Value.String()}
}

func rangeIntField(field IntField, from, to int) lookup {
	lookup := lookup{lhs: field, lookupName: "BETWEEN"}
	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
	return lookup
}

func (f *DecimalField) Range(from, to decimal.Decimal) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = from.String() + " AND " + to.String()
	return lookup
}

func fieldIsNull(field field, value bool) lookup {
	lookup := lookup{lhs: field, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f *DecimalField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}
