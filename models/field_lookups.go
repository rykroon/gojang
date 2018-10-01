package models

import (
	//"reflect"
)

type lookup struct {
	lhs        string
	lookupName string
	rhs        string
}

func (f BooleanField) Exact(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: boolToSql(value)}
}

func (f FloatField) Exact(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: float64ToSql(value)}
}

func (f IntegerField) Exact(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: intToSql(value)}
}

func (f TextField) Exact(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: stringToSql(value)}
}

func (f TextField) IExact(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "ILIKE", rhs: stringToSql(value)}
}

func (f TextField) Contains(value string) lookup {
	value = "%" + value + "%"
	return lookup{lhs: f.dbColumn, lookupName: "LIKE", rhs: stringToSql(value)}
}

func (f TextField) IContains(value string) lookup {
	value = "%" + value + "%"
	return lookup{lhs: f.dbColumn, lookupName: "ILIKE", rhs: stringToSql(value)}
}

func (f BooleanField) In(values ...bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: boolSliceToSql(values)}
}

func (f FloatField) In(values ...float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: float64SliceToSql(values)}
}

func (f IntegerField) In(values ...int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: intSliceToSql(values)}
}

func (f TextField) In(values ...string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: stringSliceToSql(values)}
}

func (f BooleanField) GT(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: boolToSql(value)}
}

func (f FloatField) GT(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: float64ToSql(value)}
}

func (f IntegerField) GT(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f TextField) GT(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: stringToSql(value)}
}

func (f BooleanField) GTE(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: boolToSql(value)}
}

func (f FloatField) GTE(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: float64ToSql(value)}
}

func (f IntegerField) GTE(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f TextField) GTE(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: stringToSql(value)}
}

func (f BooleanField) LT(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: boolToSql(value)}
}

func (f FloatField) LT(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: float64ToSql(value)}
}

func (f IntegerField) LT(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f TextField) LT(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: stringToSql(value)}
}

func (f BooleanField) LTE(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: boolToSql(value)}
}

func (f FloatField) LTE(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: float64ToSql(value)}
}

func (f IntegerField) LTE(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
}

func (f TextField) LTE(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: stringToSql(value)}
}

func (f TextField) StartsWith(value string) lookup {
	value = value + "%"
	return lookup{lhs: f.dbColumn, lookupName: "LIKE", rhs: stringToSql(value)}
}

func (f TextField) IStartsWith(value string) lookup {
	value = value + "%"
	return lookup{lhs: f.dbColumn, lookupName: "ILIKE", rhs: stringToSql(value)}
}

func (f TextField) EndsWith(value string) lookup {
	value = "%" + value
	return lookup{lhs: f.dbColumn, lookupName: "LIKE", rhs: stringToSql(value)}
}

func (f TextField) IEndsWith(value string) lookup {
	value = "%" + value
	return lookup{lhs: f.dbColumn, lookupName: "ILIKE", rhs: stringToSql(value)}
}

func (f BooleanField) Range(from bool, to bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = boolToSql(from) + " AND " + boolToSql(to)
	return lookup
}

func (f FloatField) Range(from float64, to float64) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = float64ToSql(from) + " AND " + float64ToSql(to)
	return lookup
}

func (f IntegerField) Range(from int, to int) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = intToSql(from) + " AND " + intToSql(to)
	return lookup
}

func (f TextField) Range(from string, to string) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = stringToSql(from) + " AND " + stringToSql(to)
	return lookup
}

func (f BooleanField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f FloatField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f IntegerField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f TextField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}
