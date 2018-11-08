package gojang

import (
	"database/sql/driver"
)

type TextField struct {
	*Column

	Valid bool
	Value string
}

func NewTextField() *TextField {
	field := &TextField{}
	field.Column = newColumn("TEXT")
	field.Valid = true
	return field
}

//
// func (f *TextField) asAssignment() assignment {
// 	return assignment(f.Exact(f.Value))
// }

func (f *TextField) Assign(value string) assignment {
	return newAssignment(f, stringAsSql(value))
}

func (f *TextField) copy() *TextField {
	copy := NewTextField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *TextField) copyField() field {
	return f.copy()
}

//Aggregates

func (f *TextField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *TextField) Max() *aggregate {
	return Max(f, NewTextField())
}

func (f *TextField) Min() *aggregate {
	return Min(f, NewTextField())
}

//
// Lookups
//

func (f *TextField) Exact(value string) lookup {
	rhs := stringAsSql(value)
	return exact(f, rhs)
}

func (f *TextField) IExact(value string) lookup {
	rhs := stringAsSql(value)
	return iexact(f, rhs)
}

func (f *TextField) Contains(value string) lookup {
	rhs := stringAsSql(value)
	return contains(f, rhs)
}

func (f *TextField) IContains(value string) lookup {
	rhs := stringAsSql(value)
	return icontains(f, rhs)
}

func (f *TextField) In(values ...string) lookup {
	rhs := stringsAsSql(values)
	return in(f, rhs)
}

func (f *TextField) Gt(value string) lookup {
	rhs := stringAsSql(value)
	return gt(f, rhs)
}

func (f *TextField) Gte(value string) lookup {
	rhs := stringAsSql(value)
	return gte(f, rhs)
}

func (f *TextField) Lt(value string) lookup {
	rhs := stringAsSql(value)
	return lt(f, rhs)
}

func (f *TextField) Lte(value string) lookup {
	rhs := stringAsSql(value)
	return lte(f, rhs)
}

func (f *TextField) StartsWith(value string) lookup {
	rhs := stringAsSql(value)
	return startsWith(f, rhs)
}

func (f *TextField) IStartsWith(value string) lookup {
	rhs := stringAsSql(value)
	return iStartsWith(f, rhs)
}

func (f *TextField) EndsWith(value string) lookup {
	rhs := stringAsSql(value)
	return endsWith(f, rhs)
}

func (f *TextField) IEndsWith(value string) lookup {
	rhs := stringAsSql(value)
	return iEndsWith(f, rhs)
}

func (f *TextField) Range(from, to string) lookup {
	rhs1 := stringAsSql(from)
	rhs2 := stringAsSql(to)
	return between(f, rhs1, rhs2)
}

func (f *TextField) IsNull(value bool) lookup {
	return isNull(f, value)
}

//
// Transforms
//

func (f *TextField) Length() *IntegerField {
	length := Length(f)
	return length.toField().(*IntegerField)
}

func (f *TextField) Lower() *TextField {
	lower := Lower(f)
	return lower.toField().(*TextField)
}

func (f *TextField) Upper() *TextField {
	upper := Upper(f)
	return upper.toField().(*TextField)
}

//
// Scanner, Valuer
//

func (f *TextField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(string)
	return nil
}

func (f *TextField) xValue() (driver.Value, error) {
	return f.Value, nil
}

func (f *TextField) getValue() interface{} {
	return f.Value
}

func (f *TextField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *TextField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *TextField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return stringAsSql(f.Value)
	}
}
