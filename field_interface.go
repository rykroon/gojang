package gojang

import ()

type field interface {
	selectExpression

	//getters and setters
	// DbColumn() string
	// setDbColumn(string)
	ColumnName() string
	setColumnName(string)

	Model() *Model
	setModel(*Model)
	HasModel() bool

	Null() bool
	setNull(bool)

	Unique() bool
	setUnique(bool)

	PrimaryKey() bool
	setPrimaryKey(bool)

	HasRelation() bool

	//asAssignment() assignment
	Asc() orderByExpression
	Desc() orderByExpression
	Count(bool) *aggregate

	copyField() field

	validate()
	valueAsSql() string
}

func (f *CharField) copyField() field {
	return f.copy()
}

func (f *DecimalField) copyField() field {
	return f.copy()
}

func (f *DecimalField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *DecimalField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *ForeignKey) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *OneToOneField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}

	if !f.unique {
		panic(NewForceConstraint(f, "unique"))
	}
}

func (f *DecimalField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return f.Value.String()
	}
}
