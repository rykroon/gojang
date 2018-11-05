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

func (f *BigIntegerField) copyField() field {
	return f.copy()
}

func (f *BooleanField) copyField() field {
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

func (f *Column) HasNullConstraint() bool {
	return f.null
}

func (f *BigIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *BooleanField) setNullConstraint(null bool) {
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

func (f *Column) HasUniqueConstraint() bool {
	return f.unique
}

func (c *Column) setUniqueConstraint(unique bool) {
	c.unique = unique
}

func (c *Column) HasPrimaryKeyConstraint() bool {
	return c.primaryKey
}

func (c *Column) setPrimaryKeyConstraint(primaryKey bool) {
	c.primaryKey = primaryKey
}

func (c *Column) HasRelation() bool {
	return c.isRelation
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

func (f *BigIntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *BooleanField) validate() {
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

func (c *Column) setModel(model *Model) {
	c.model = model
}

func (c *Column) Model() *Model {
	return c.model
}

func (f *Column) HasModel() bool {
	return f.model != nil
}

func (f *Column) DbColumn() string {
	return f.dbColumn
}

func (f *Column) setDbColumn(col string) {
	f.dbColumn = col
}

func (f *BigIntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return int64AsSql(f.Value)
	}
}

func (f *BooleanField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return boolAsSql(f.Value)
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
		return int32AsSql(f.Value)
	}
}

func (f *SmallIntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return int16AsSql(f.Value)
	}
}

func (f *TextField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return stringAsSql(f.Value)
	}
}
