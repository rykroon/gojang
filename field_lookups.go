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

func (f *BigIntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
}

func (f *BooleanField) Exact(value bool) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: boolAsSql(value)}
}

func (f *DecimalField) Exact(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: f.Value.String()}
}

func (f *FloatField) Exact(value float64) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: float64AsSql(value)}
}

func (f *IntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
}

func (f *SmallIntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
}

func (f *TextField) Exact(value string) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: stringAsSql(value)}
}

func (f *TextField) IExact(value string) lookup {
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func (f *TextField) Contains(value string) lookup {
	value = "%" + value + "%"
	return lookup{lhs: f, lookupName: "LIKE", rhs: stringAsSql(value)}
}

func (f *TextField) IContains(value string) lookup {
	value = "%" + value + "%"
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func inIntField(field IntField, values []int) lookup {
	return lookup{lhs: field, lookupName: "IN", rhs: integersAsSql(values)}
}

func (f *BigIntegerField) In(values ...int) lookup {
	return inIntField(f, values)
}

func (f *BooleanField) In(values ...bool) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: boolSliceAsSql(values)}
}

func (f *FloatField) In(values ...float64) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: float64SliceAsSql(values)}
}

func (f *IntegerField) In(values ...int) lookup {
	return inIntField(f, values)
}

func (f *SmallIntegerField) In(values ...int) lookup {
	return inIntField(f, values)
}

func (f *TextField) In(values ...string) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: stringSliceAsSql(values)}
}

func gtIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: ">", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *BooleanField) Gt(value bool) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: boolAsSql(value)}
}

func (f *DecimalField) Gt(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: f.Value.String()}
}

func (f *FloatField) Gt(value float64) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: float64AsSql(value)}
}

func (f *IntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *SmallIntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *TextField) Gt(value string) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: stringAsSql(value)}
}

func gteIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: ">=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
	//return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
}

func (f *BooleanField) Gte(value bool) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: boolAsSql(value)}
}

func (f *DecimalField) Gte(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: f.Value.String()}
}

func (f *FloatField) Gte(value float64) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: float64AsSql(value)}
}

func (f *IntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
}

func (f *SmallIntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
}

func (f *TextField) Gte(value string) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: stringAsSql(value)}
}

func ltIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "<", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
	//return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
}

func (f *BooleanField) Lt(value bool) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: boolAsSql(value)}
}

func (f *DecimalField) Lt(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: f.Value.String()}
}

func (f *FloatField) Lt(value float64) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: float64AsSql(value)}
}

func (f *IntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
}

func (f *SmallIntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
}

func (f *TextField) Lt(value string) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: stringAsSql(value)}
}

func lteIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "<=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
	//return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
}

func (f *BooleanField) Lte(value bool) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: boolAsSql(value)}
}

func (f *DecimalField) Lte(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: f.Value.String()}
}

func (f *FloatField) Lte(value float64) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: float64AsSql(value)}
}

func (f *IntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
}

func (f *SmallIntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
}

func (f *TextField) Lte(value string) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: stringAsSql(value)}
}

func (f *TextField) StartsWith(value string) lookup {
	value = value + "%"
	return lookup{lhs: f, lookupName: "LIKE", rhs: stringAsSql(value)}
}

func (f *TextField) IStartsWith(value string) lookup {
	value = value + "%"
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func (f *TextField) EndsWith(value string) lookup {
	value = "%" + value
	return lookup{lhs: f, lookupName: "LIKE", rhs: stringAsSql(value)}
}

func (f *TextField) IEndsWith(value string) lookup {
	value = "%" + value
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func rangeIntField(field IntField, from, to int) lookup {
	lookup := lookup{lhs: field, lookupName: "BETWEEN"}
	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
	return lookup
}

func (f *BigIntegerField) Range(from, to int) lookup {
	return rangeIntField(f, from, to)
}

func (f *BooleanField) Range(from, to bool) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = boolAsSql(from) + " AND " + boolAsSql(to)
	return lookup
}

func (f *DecimalField) Range(from, to decimal.Decimal) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = from.String() + " AND " + to.String()
	return lookup
}

func (f *FloatField) Range(from, to float64) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = float64AsSql(from) + " AND " + float64AsSql(to)
	return lookup
}

func (f *IntegerField) Range(from, to int) lookup {
	return rangeIntField(f, from, to)
}

func (f *SmallIntegerField) Range(from, to int) lookup {
	return rangeIntField(f, from, to)
}

func (f *TextField) Range(from, to string) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = stringAsSql(from) + " AND " + stringAsSql(to)
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

func (f *BigIntegerField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *BooleanField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *DecimalField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *FloatField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *IntegerField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *TextField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}
