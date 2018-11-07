package gojang

import ()

func NewBooleanField() *BooleanField {
	field := &BooleanField{}
	field.Column = newColumn("BOOL")
	field.Valid = true
	return field
}

func (f *BooleanField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}


func (f *BooleanField) copy() *BooleanField {
	copy := NewBooleanField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *BooleanField) copyField() field {
	return f.copy()
}

func (f *BooleanField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *BooleanField) Exact(value bool) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: boolAsSql(value)}
}

func (f *BooleanField) getValue() interface{} {
	return f.Value
}

func (f *BooleanField) Gt(value bool) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: boolAsSql(value)}
}

func (f *BooleanField) Gte(value bool) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: boolAsSql(value)}
}

func (f *BooleanField) In(values ...bool) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: boolsAsSql(values)}
}

func (f *BooleanField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *BooleanField) Lt(value bool) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: boolAsSql(value)}
}

func (f *BooleanField) Lte(value bool) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: boolAsSql(value)}
}

func (f *BooleanField) Max() *aggregate {
	return Max(f, NewBooleanField())
}

func (f *BooleanField) Min() *aggregate {
	return Min(f, NewBooleanField())
}

func (f *BooleanField) Range(from, to bool) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = boolAsSql(from) + " AND " + boolAsSql(to)
	return lookup
}

func (f *BooleanField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(bool)
	return nil
}

func (f *BooleanField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *BooleanField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *BooleanField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return boolAsSql(f.Value)
	}
}
