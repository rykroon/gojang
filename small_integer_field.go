package gojang

import (
	"database/sql/driver"
)

type SmallIntegerField struct {
	*Column

	Valid bool
	Value int16
}

func NewSmallIntegerField() *SmallIntegerField {
	field := &SmallIntegerField{}
	field.Column = newColumn("INT2")
	field.Valid = true
	return field
}

// func (f *SmallIntegerField) asAssignment() assignment {
// 	return assignment(f.Exact(int(f.Value)))
// }

func (f *SmallIntegerField) Assign(value int16) assignment {
	return newAssignment(f, intAsSql(int(value)))
}

func (f *SmallIntegerField) copy() *SmallIntegerField {
	copy := NewSmallIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *SmallIntegerField) copyField() field {
	return f.copy()
}

//
// Aggreagates
//

func (f *SmallIntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *SmallIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *SmallIntegerField) Max() *aggregate {
	return Max(f, NewSmallIntegerField())
}

func (f *SmallIntegerField) Min() *aggregate {
	return Min(f, NewSmallIntegerField())
}

func (f *SmallIntegerField) Sum() *aggregate {
	return Sum(f, NewSmallIntegerField())
}

//
// Lookups
//

func (f *SmallIntegerField) Exact(value int) lookup {
	rhs := intAsSql(value)
	return exact(f, rhs)
}

func (f *SmallIntegerField) In(values ...int) lookup {
	rhs := integersAsSql(values)
	return in(f, rhs)
}

func (f *SmallIntegerField) Gt(value int) lookup {
	rhs := intAsSql(value)
	return gt(f, rhs)
}

func (f *SmallIntegerField) Gte(value int) lookup {
	rhs := intAsSql(value)
	return gte(f, rhs)
}

func (f *SmallIntegerField) Lt(value int) lookup {
	rhs := intAsSql(value)
	return lt(f, rhs)
}

func (f *SmallIntegerField) Lte(value int) lookup {
	rhs := intAsSql(value)
	return lte(f, rhs)
}

func (f *SmallIntegerField) Range(from, to int) lookup {
	rhs1 := intAsSql(from)
	rhs2 := intAsSql(to)
	return between(f, rhs1, rhs2)
}

func (f *SmallIntegerField) IsNull(value bool) lookup {
	return isNull(f, value)
}

//
// Scanner, Valuer
//

func (f *SmallIntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.Valid = int16(result), ok
	return nil
}

func (f *SmallIntegerField) xValue() (driver.Value, error) {
	return int64(f.Value), nil
}

func (f *SmallIntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *SmallIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *SmallIntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *SmallIntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Value))
	}
}
