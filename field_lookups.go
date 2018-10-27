package gojang

import (
//"reflect"
)

type lookup struct {
	not        bool
	lhs        field
	lookupName string
	rhs        string
}

// func (f *AutoField) Exact(value int) lookup {
// 	return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
// }
//
// func (f *BigAutoField) Exact(value int) lookup {
// 	return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
// }

func exactIntField(field intField, value int) lookup {
	return lookup{lhs: field, lookupName: "=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
	//return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
}

func (f *BooleanField) Exact(value bool) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: boolAsSql(value)}
}

func (f *FloatField) Exact(value float64) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: float64AsSql(value)}
}

func (f *IntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
	//return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
}

func (f *SmallIntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
	//return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
}

func (f *TextField) Exact(value string) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: stringAsSql(value)}
}

// func (f *ForeignKey) Exact(value int) lookup {
// 	return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
// }
//
// func (f *OneToOneField) Exact(value int) lookup {
// 	return lookup{lhs: f, lookupName: "=", rhs: intAsSql(value)}
// }

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

// func (f *AutoField) In(values ...int) lookup {
// 	return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
// }
//
// func (f *BigAutoField) In(values ...int) lookup {
// 	return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
// }

func inIntField(field intField, values []int) lookup {
	return lookup{lhs: field, lookupName: "IN", rhs: integersAsSql(values)}
}

func (f *BigIntegerField) In(values ...int) lookup {
	return inIntField(f, values)
	//return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
}

func (f *BooleanField) In(values ...bool) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: boolSliceAsSql(values)}
}

func (f *FloatField) In(values ...float64) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: float64SliceAsSql(values)}
}

func (f *IntegerField) In(values ...int) lookup {
	return inIntField(f, values)
	//return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
}

func (f *SmallIntegerField) In(values ...int) lookup {
	return inIntField(f, values)
	//return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
}

func (f *TextField) In(values ...string) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: stringSliceAsSql(values)}
}

// func (f *ForeignKey) In(values ...int) lookup {
// 	return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
// }
//
// func (f *OneToOneField) In(values ...int) lookup {
// 	return lookup{lhs: f, lookupName: "IN", rhs: integersAsSql(values)}
// }

// func (f *AutoField) Gt(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">", rhs: intAsSql(value)}
// }
//
// func (f *BigAutoField) Gt(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">", rhs: intAsSql(value)}
// }

func gtIntField(field intField, value int) lookup {
	return lookup{lhs: field, lookupName: ">", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
	//return lookup{lhs: f, lookupName: ">", rhs: intAsSql(value)}
}

func (f *BooleanField) Gt(value bool) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: boolAsSql(value)}
}

func (f *FloatField) Gt(value float64) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: float64AsSql(value)}
}

func (f *IntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
	//return lookup{lhs: f, lookupName: ">", rhs: intAsSql(value)}
}

func (f *SmallIntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *TextField) Gt(value string) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: stringAsSql(value)}
}

// func (f *ForeignKey) Gt(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">", rhs: intAsSql(value)}
// }
//
// func (f *OneToOneField) Gt(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">", rhs: intAsSql(value)}
// }

// func (f *AutoField) Gte(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
// }
//
// func (f *BigAutoField) Gte(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
// }

func gteIntField(field intField, value int) lookup {
	return lookup{lhs: field, lookupName: ">=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
	//return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
}

func (f *BooleanField) Gte(value bool) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: boolAsSql(value)}
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

// func (f *ForeignKey) Gte(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
// }
//
// func (f *OneToOneField) Gte(value int) lookup {
// 	return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
// }

// func (f *AutoField) Lt(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
// }
//
// func (f *BigAutoField) Lt(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
// }

func ltIntField(field intField, value int) lookup {
	return lookup{lhs: field, lookupName: "<", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
	//return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
}

func (f *BooleanField) Lt(value bool) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: boolAsSql(value)}
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

// func (f *ForeignKey) Lt(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
// }
//
// func (f *OneToOneField) Lt(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
// }

// func (f *AutoField) Lte(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
// }
//
// func (f *BigAutoField) Lte(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
// }

func lteIntField(field intField, value int) lookup {
	return lookup{lhs: field, lookupName: "<=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
	//return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
}

func (f *BooleanField) Lte(value bool) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: boolAsSql(value)}
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

// func (f *ForeignKey) Lte(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
// }
//
// func (f *OneToOneField) Lte(value int) lookup {
// 	return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
// }

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

// func (f *AutoField) Range(from int, to int) lookup {
// 	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
// 	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
// 	return lookup
// }
//
// func (f *BigAutoField) Range(from int, to int) lookup {
// 	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
// 	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
// 	return lookup
// }

func rangeIntField(field intField, from, to int) lookup {
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

// func (f *ForeignKey) Range(from int, to int) lookup {
// 	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
// 	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
// 	return lookup
// }
//
// func (f *OneToOneField) Range(from int, to int) lookup {
// 	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
// 	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
// 	return lookup
// }

// func (f *AutoField) IsNull(value bool) lookup {
// 	lookup := lookup{lhs: f, lookupName: "IS"}
//
// 	if value {
// 		lookup.rhs = "NULL"
// 	} else {
// 		lookup.rhs = "NOT NULL"
// 	}
//
// 	return lookup
// }
//
// func (f *BigAutoField) IsNull(value bool) lookup {
// 	lookup := lookup{lhs: f, lookupName: "IS"}
//
// 	if value {
// 		lookup.rhs = "NULL"
// 	} else {
// 		lookup.rhs = "NOT NULL"
// 	}
//
// 	return lookup
// }

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

func (f *FloatField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *IntegerField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *TextField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

// func (f *ForeignKey) IsNull(value bool) lookup {
// 	lookup := lookup{lhs: f, lookupName: "IS"}
//
// 	if value {
// 		lookup.rhs = "NULL"
// 	} else {
// 		lookup.rhs = "NOT NULL"
// 	}
//
// 	return lookup
// }
//
// func (f *OneToOneField) IsNull(value bool) lookup {
// 	lookup := lookup{lhs: f, lookupName: "IS"}
//
// 	if value {
// 		lookup.rhs = "NULL"
// 	} else {
// 		lookup.rhs = "NOT NULL"
// 	}
//
// 	return lookup
// }
