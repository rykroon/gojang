package gojang

import (
//"reflect"
)

type lookup struct {
	lhs        string
	lookupName string
	rhs        string
}

func (f AutoField) Exact(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: intToSql(value)}
}

func (f BigAutoField) Exact(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: intToSql(value)}
}

func (f BigIntegerField) Exact(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: intToSql(value)}
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

func (f ForeignKeyField) Exact(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: intToSql(value)}
}

func (f OneToOneField) Exact(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "=", rhs: intToSql(value)}
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

func (f AutoField) In(values ...int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: intSliceToSql(values)}
}

func (f BigAutoField) In(values ...int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: intSliceToSql(values)}
}

func (f BigIntegerField) In(values ...int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: intSliceToSql(values)}
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

func (f ForeignKeyField) In(values ...int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: intSliceToSql(values)}
}

func (f OneToOneField) In(values ...int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "IN", rhs: intSliceToSql(values)}
}

func (f AutoField) Gt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f BigAutoField) Gt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f BigIntegerField) Gt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f BooleanField) Gt(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: boolToSql(value)}
}

func (f FloatField) Gt(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: float64ToSql(value)}
}

func (f IntegerField) Gt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f TextField) Gt(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: stringToSql(value)}
}

func (f ForeignKeyField) Gt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f OneToOneField) Gt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">", rhs: intToSql(value)}
}

func (f AutoField) Gte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f BigAutoField) Gte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f BigIntegerField) Gte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f BooleanField) Gte(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: boolToSql(value)}
}

func (f FloatField) Gte(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: float64ToSql(value)}
}

func (f IntegerField) Gte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f TextField) Gte(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: stringToSql(value)}
}

func (f ForeignKeyField) Gte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f OneToOneField) Gte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: ">=", rhs: intToSql(value)}
}

func (f AutoField) Lt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f BigAutoField) Lt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f BigIntegerField) Lt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f BooleanField) Lt(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: boolToSql(value)}
}

func (f FloatField) Lt(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: float64ToSql(value)}
}

func (f IntegerField) Lt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f TextField) Lt(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: stringToSql(value)}
}

func (f ForeignKeyField) Lt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f OneToOneField) Lt(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<", rhs: intToSql(value)}
}

func (f AutoField) Lte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
}

func (f BigAutoField) Lte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
}

func (f BigIntegerField) Lte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
}

func (f BooleanField) Lte(value bool) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: boolToSql(value)}
}

func (f FloatField) Lte(value float64) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: float64ToSql(value)}
}

func (f IntegerField) Lte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
}

func (f TextField) Lte(value string) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: stringToSql(value)}
}

func (f ForeignKeyField) Lte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
}

func (f OneToOneField) Lte(value int) lookup {
	return lookup{lhs: f.dbColumn, lookupName: "<=", rhs: intToSql(value)}
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

func (f AutoField) Range(from int, to int) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = intToSql(from) + " AND " + intToSql(to)
	return lookup
}

func (f BigAutoField) Range(from int, to int) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = intToSql(from) + " AND " + intToSql(to)
	return lookup
}

func (f BigIntegerField) Range(from int, to int) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = intToSql(from) + " AND " + intToSql(to)
	return lookup
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

func (f ForeignKeyField) Range(from int, to int) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = intToSql(from) + " AND " + intToSql(to)
	return lookup
}

func (f OneToOneField) Range(from int, to int) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "BETWEEN"}
	lookup.rhs = intToSql(from) + " AND " + intToSql(to)
	return lookup
}

func (f AutoField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f BigAutoField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f BigIntegerField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

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

func (f ForeignKeyField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}

func (f OneToOneField) IsNull(value bool) lookup {
	lookup := lookup{lhs: f.dbColumn, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}
