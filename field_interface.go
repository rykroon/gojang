package gojang

import ()

type field interface {
	selectExpression

	//getters and setters
	DbColumn() string
	setDbColumn(string)

	Model() *Model
	setModel(*Model)
	HasModel() bool

	HasNullConstraint() bool
	setNullConstraint(bool)

	HasUniqueConstraint() bool
	setUniqueConstraint(bool)

	HasPrimaryKeyConstraint() bool
	setPrimaryKeyConstraint(bool)

	HasRelation() bool

	asAssignment() assignment
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

func (f *FloatField) copyField() field {
	return f.copy()
}

func (f *IntegerField) copyField() field {
	return f.copy()
}

func (f *SmallIntegerField) copyField() field {
	return f.copy()
}

func (f *TextField) copyField() field {
	return f.copy()
}







func (f *DecimalField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *FloatField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *IntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *SmallIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *TextField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *AutoField) validate() {
	if !f.primaryKey {
		panic(NewForceConstraint(f, "primary key"))
	}

	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *BigAutoField) validate() {
	if !f.primaryKey {
		panic(NewForceConstraint(f, "primary key"))
	}

	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}



func (f *DecimalField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *FloatField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *IntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *SmallIntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *TextField) validate() {
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

func (f *FloatField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return float64AsSql(f.Value)
	}
}

func (f *IntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Value))
	}
}

func (f *SmallIntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Value))
	}
}

func (f *TextField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return stringAsSql(f.Value)
	}
}
