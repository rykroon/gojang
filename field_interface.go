package gojang

import ()

type field interface {
	selectExpression

	//getters and setters
	getDbColumn() string
	setDbColumn(string)

	getDbType() string

	getExpr() expression
	setExpr(expression)

	getModel() *Model
	setModel(*Model)
	hasModel() bool

	hasNullConstraint() bool
	setNullConstraint(bool)

	hasUniqueConstraint() bool
	setUniqueConstraint(bool)

	hasPrimaryKeyConstraint() bool
	setPrimaryKeyConstraint(bool)

	hasRelation() bool

	IsNil() bool
	SetNil() error
	UnSetNil()

	asAssignment() assignment
	Asc() sortExpression
	Desc() sortExpression
	Count(bool) *aggregate

	copyField() field
	//new() field //haven't found a use for this yet

	validate()
	valueAsSql() string
}

func (f *BigIntegerField) getExpr() expression {
	return f.expr
}

func (f *BooleanField) getExpr() expression {
	return f.expr
}

func (f *FloatField) getExpr() expression {
	return f.expr
}

func (f *IntegerField) getExpr() expression {
	return f.expr
}

func (f *SmallIntegerField) getExpr() expression {
	return f.expr
}

func (f *TextField) getExpr() expression {
	return f.expr
}

func (f *BigIntegerField) setExpr(expr expression) {
	f.expr = expr
}

func (f *BooleanField) setExpr(expr expression) {
	f.expr = expr
}

func (f *FloatField) setExpr(expr expression) {
	f.expr = expr
}

func (f *IntegerField) setExpr(expr expression) {
	f.expr = expr
}

func (f *SmallIntegerField) setExpr(expr expression) {
	f.expr = expr
}

func (f *TextField) setExpr(expr expression) {
	f.expr = expr
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

func (f *BigIntegerField) new() field {
	return NewBigIntegerField()
}

func (f *BooleanField) new() field {
	return NewBooleanField()
}

func (f *FloatField) new() field {
	return NewFloatField()
}

func (f *IntegerField) new() field {
	return NewIntegerField()
}

func (f *SmallIntegerField) new() field {
	return NewSmallIntegerField()
}

func (f *TextField) new() field {
	return NewTextField()
}

func (f *BigIntegerField) hasNullConstraint() bool {
	return f.null
}

func (f *BooleanField) hasNullConstraint() bool {
	return f.null
}

func (f *FloatField) hasNullConstraint() bool {
	return f.null
}

func (f *IntegerField) hasNullConstraint() bool {
	return f.null
}

func (f *SmallIntegerField) hasNullConstraint() bool {
	return f.null
}

func (f *TextField) hasNullConstraint() bool {
	return f.null
}

func (f *BigIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.valid = false
	}
}

func (f *BooleanField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.valid = false
	}
}

func (f *FloatField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.valid = false
	}
}

func (f *IntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.valid = false
	}
}

func (f *SmallIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.valid = false
	}
}

func (f *TextField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.valid = false
	}
}

func (f *BigIntegerField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *BooleanField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *FloatField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *IntegerField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *SmallIntegerField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *TextField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *BigIntegerField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *BooleanField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *FloatField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *IntegerField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *SmallIntegerField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *TextField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *BigIntegerField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *BooleanField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *FloatField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *IntegerField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *SmallIntegerField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *TextField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *BigIntegerField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *BooleanField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *FloatField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *IntegerField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *SmallIntegerField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *TextField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *BigIntegerField) hasRelation() bool {
	return f.isRelation
}

func (f *BooleanField) hasRelation() bool {
	return f.isRelation
}

func (f *FloatField) hasRelation() bool {
	return f.isRelation
}

func (f *IntegerField) hasRelation() bool {
	return f.isRelation
}

func (f *SmallIntegerField) hasRelation() bool {
	return f.isRelation
}

func (f *TextField) hasRelation() bool {
	return f.isRelation
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

func (f *BigIntegerField) setModel(model *Model) {
	f.model = model
}

func (f *BooleanField) setModel(model *Model) {
	f.model = model
}

func (f *FloatField) setModel(model *Model) {
	f.model = model
}

func (f *IntegerField) setModel(model *Model) {
	f.model = model
}

func (f *SmallIntegerField) setModel(model *Model) {
	f.model = model
}

func (f *TextField) setModel(model *Model) {
	f.model = model
}

func (f *BigIntegerField) getModel() *Model {
	return f.model
}

func (f *BooleanField) getModel() *Model {
	return f.model
}

func (f *FloatField) getModel() *Model {
	return f.model
}

func (f *IntegerField) getModel() *Model {
	return f.model
}

func (f *SmallIntegerField) getModel() *Model {
	return f.model
}

func (f *TextField) getModel() *Model {
	return f.model
}

func (f *BigIntegerField) hasModel() bool {
	return f.model != nil
}

func (f *BooleanField) hasModel() bool {
	return f.model != nil
}

func (f *FloatField) hasModel() bool {
	return f.model != nil
}

func (f *IntegerField) hasModel() bool {
	return f.model != nil
}

func (f *SmallIntegerField) hasModel() bool {
	return f.model != nil
}

func (f *TextField) hasModel() bool {
	return f.model != nil
}

func (f *BigIntegerField) getDbColumn() string {
	return f.dbColumn
}

func (f *BooleanField) getDbColumn() string {
	return f.dbColumn
}

func (f *FloatField) getDbColumn() string {
	return f.dbColumn
}

func (f *IntegerField) getDbColumn() string {
	return f.dbColumn
}

func (f *SmallIntegerField) getDbColumn() string {
	return f.dbColumn
}

func (f *TextField) getDbColumn() string {
	return f.dbColumn
}

func (f *BigIntegerField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BooleanField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *FloatField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *IntegerField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *SmallIntegerField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *TextField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BigIntegerField) getDbType() string {
	return f.dbType
}

func (f *BooleanField) getDbType() string {
	return f.dbType
}

func (f *FloatField) getDbType() string {
	return f.dbType
}

func (f *IntegerField) getDbType() string {
	return f.dbType
}

func (f *SmallIntegerField) getDbType() string {
	return f.dbType
}

func (f *TextField) getDbType() string {
	return f.dbType
}

func (f *BigIntegerField) IsNil() bool {
	return !f.valid
}

func (f *BooleanField) IsNil() bool {
	return !f.valid
}

func (f *FloatField) IsNil() bool {
	return !f.valid
}

func (f *IntegerField) IsNil() bool {
	return !f.valid
}

func (f *SmallIntegerField) IsNil() bool {
	return !f.valid
}

func (f *TextField) IsNil() bool {
	return !f.valid
}

func (f *BigIntegerField) valueAsSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64AsSql(f.Value)
	}
}

func (f *BooleanField) valueAsSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return boolAsSql(f.Value)
	}
}

func (f *FloatField) valueAsSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return float64AsSql(f.Value)
	}
}

func (f *IntegerField) valueAsSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int32AsSql(f.Value)
	}
}

func (f *SmallIntegerField) valueAsSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int16AsSql(f.Value)
	}
}

func (f *TextField) valueAsSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return stringAsSql(f.Value)
	}
}
