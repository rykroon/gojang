package gojang

import ()

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

func (f *TextField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *TextField) Assign(value string) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *TextField) Contains(value string) lookup {
	value = "%" + value + "%"
	return lookup{lhs: f, lookupName: "LIKE", rhs: stringAsSql(value)}
}

func (f *TextField) copy() *TextField {
	copy := NewTextField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *TextField) copyField() field {
	return f.copy()
}

func (f *TextField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *TextField) EndsWith(value string) lookup {
	value = "%" + value
	return lookup{lhs: f, lookupName: "LIKE", rhs: stringAsSql(value)}
}

func (f *TextField) Exact(value string) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: stringAsSql(value)}
}

func (f *TextField) getValue() interface{} {
	return f.Value
}

func (f *TextField) Gt(value string) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: stringAsSql(value)}
}

func (f *TextField) Gte(value string) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: stringAsSql(value)}
}

func (f *TextField) IContains(value string) lookup {
	value = "%" + value + "%"
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func (f *TextField) IEndsWith(value string) lookup {
	value = "%" + value
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func (f *TextField) IExact(value string) lookup {
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func (f *TextField) In(values ...string) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: stringsAsSql(values)}
}

func (f *TextField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *TextField) IStartsWith(value string) lookup {
	value = value + "%"
	return lookup{lhs: f, lookupName: "ILIKE", rhs: stringAsSql(value)}
}

func (f *TextField) Length() *IntegerField {
	length := Length(f)
	return length.toField().(*IntegerField)
}

func (f *TextField) Lower() *TextField {
	lower := Lower(f)
	return lower.toField().(*TextField)
}

func (f *TextField) Lt(value string) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: stringAsSql(value)}
}

func (f *TextField) Lte(value string) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: stringAsSql(value)}
}

func (f *TextField) Max() *aggregate {
	return Max(f, NewTextField())
}

func (f *TextField) Min() *aggregate {
	return Min(f, NewTextField())
}

func (f *TextField) StartsWith(value string) lookup {
	value = value + "%"
	return lookup{lhs: f, lookupName: "LIKE", rhs: stringAsSql(value)}
}

func (f *TextField) Range(from, to string) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = stringAsSql(from) + " AND " + stringAsSql(to)
	return lookup
}

func (f *TextField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(string)
	return nil
}

func (f *TextField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *TextField) Upper() *TextField {
	upper := Upper(f)
	return upper.toField().(*TextField)
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
