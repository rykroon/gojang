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

	//setOptions(fieldOptions)
	hasRelation() bool

	IsNil() bool
	SetNil() error
	UnSetNil()

	asAssignment() assignment
	Asc() sortExpression
	Desc() sortExpression
	copy() field
	Count(bool) aggregate

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

// func (f *AutoField) copy() field {
// 	copy := NewAutoField()
// 	copy.model = f.model
// 	copy.dbColumn = f.dbColumn
// 	return copy
// }
//
// func (f *BigAutoField) copy() field {
// 	copy := NewBigAutoField()
// 	copy.model = f.model
// 	copy.dbColumn = f.dbColumn
// 	return copy
// }

func (f *BigIntegerField) copy() field {
	copy := NewBigIntegerField()
	copy.model = f.model //maybe change to be a copy of the model
	copy.dbColumn = f.dbColumn
	copy.constraints = f.constraints
	return copy
}

func (f *BooleanField) copy() field {
	copy := NewBooleanField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.constraints = f.constraints
	return copy
}

func (f *FloatField) copy() field {
	copy := NewFloatField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.constraints = f.constraints
	return copy
}

func (f *IntegerField) copy() field {
	copy := NewIntegerField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.constraints = f.constraints
	return copy
}

func (f *SmallIntegerField) copy() field {
	copy := NewSmallIntegerField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.constraints = f.constraints
	return copy
}

func (f *TextField) copy() field {
	copy := NewTextField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.constraints = f.constraints
	return copy
}

// func (f *ForeignKey) copy() field {
// 	copy := NewForeignKey(f.relatedModel, f.onDelete)
// 	copy.model = f.model
// 	copy.dbColumn = f.dbColumn
//	copy.constraints = f.constraints
// 	return copy
// }
//
// func (f *OneToOneField) copy() field {
// 	copy := NewOneToOneField(f.relatedModel, f.onDelete)
// 	copy.model = f.model
// 	copy.dbColumn = f.dbColumn
//	copy.constraints = f.constraints
// 	return copy
// }

// func (f *AutoField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *BigAutoField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }

// func (f *BigIntegerField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *BooleanField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *FloatField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *IntegerField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *SmallIntegerField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *TextField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }

// func (f *ForeignKey) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }
//
// func (f *OneToOneField) setOptions(options fieldOptions) {
// 	f.dbColumn = options.dbColumn
// 	f.primaryKey = options.primaryKey
// 	f.null = options.null
// 	f.unique = options.unique
// }

// func (f *AutoField) hasNullConstraint() bool {
// 	return f.null
// }
//
// func (f *BigAutoField) hasNullConstraint() bool {
// 	return f.null
// }

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

// func (f *ForeignKey) setNullConstraint(null bool) {
// 	f.null = null
//
// 	if f.null {
// 		f.valid = false
// 	}
// }
//
// func (f *OneToOneField) setNullConstraint(null bool) {
// 	f.null = null
//
// 	if f.null {
// 		f.valid = false
// 	}
// }

// func (f *AutoField) hasUniqueConstraint() bool {
// 	return f.unique
// }
//
// func (f *BigAutoField) hasUniqueConstraint() bool {
// 	return f.unique
// }

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

// func (f *ForeignKey) hasUniqueConstraint() bool {
// 	return f.unique
// }
//
// func (f *OneToOneField) hasUniqueConstraint() bool {
// 	return f.unique
// }

// func (f *AutoField) setUniqueConstraint(unique bool) {
// 	f.unique = unique
// }
//
// func (f *BigAutoField) setUniqueConstraint(unique bool) {
// 	f.unique = unique
// }

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

// func (f *ForeignKey) setUniqueConstraint(unique bool) {
// 	f.unique = unique
// }
//
// //The 'Unique' field option is valid on all field types except ManyToManyField and OneToOneField.
// func (f *OneToOneField) setUniqueConstraint(unique bool) {
// 	f.unique = unique
// }
//
// func (f *AutoField) hasPrimaryKeyConstraint() bool {
// 	return f.primaryKey
// }
//
// func (f *BigAutoField) hasPrimaryKeyConstraint() bool {
// 	return f.primaryKey
// }

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

// func (f *ForeignKey) hasPrimaryKeyConstraint() bool {
// 	return f.primaryKey
// }
//
// func (f *OneToOneField) hasPrimaryKeyConstraint() bool {
// 	return f.primaryKey
// }
//
// func (f *AutoField) setPrimaryKeyConstraint(primaryKey bool) {
// 	f.primaryKey = primaryKey
// }
//
// func (f *BigAutoField) setPrimaryKeyConstraint(primaryKey bool) {
// 	f.primaryKey = primaryKey
// }

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

// func (f *ForeignKey) setPrimaryKeyConstraint(primaryKey bool) {
// 	f.primaryKey = primaryKey
// }
//
// func (f *OneToOneField) setPrimaryKeyConstraint(primaryKey bool) {
// 	f.primaryKey = primaryKey
// }
//
// func (f *AutoField) hasRelation() bool {
// 	return f.isRelation
// }
//
// func (f *BigAutoField) hasRelation() bool {
// 	return f.isRelation
// }

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

// func (f *ForeignKey) hasRelation() bool {
// 	return f.isRelation
// }
//
// func (f *OneToOneField) hasRelation() bool {
// 	return f.isRelation
// }

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

// func (f *AutoField) setModel(model *Model) {
// 	f.model = model
// }
//
// func (f *BigAutoField) setModel(model *Model) {
// 	f.model = model
// }

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

// func (f *ForeignKey) setModel(model *Model) {
// 	f.model = model
// }
//
// func (f *OneToOneField) setModel(model *Model) {
// 	f.model = model
// }

// func (f *AutoField) getModel() *Model {
// 	return f.model
// }
//
// func (f *BigAutoField) getModel() *Model {
// 	return f.model
// }

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

// func (f *ForeignKey) getModel() *Model {
// 	return f.model
// }
//
// func (f *OneToOneField) getModel() *Model {
// 	return f.model
// }

// func (f *AutoField) hasModel() bool {
// 	return f.model != nil
// }
//
// func (f *BigAutoField) hasModel() bool {
// 	return f.model != nil
// }

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

// func (f *ForeignKey) hasModel() bool {
// 	return f.model != nil
// }
//
// func (f *OneToOneField) hasModel() bool {
// 	return f.model != nil
// }

// func (f *AutoField) getDbColumn() string {
// 	return f.dbColumn
// }
//
// func (f *BigAutoField) getDbColumn() string {
// 	return f.dbColumn
// }

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

// func (f *ForeignKey) getDbColumn() string {
// 	return f.dbColumn
// }
//
// func (f *OneToOneField) getDbColumn() string {
// 	return f.dbColumn
// }

// func (f *AutoField) setDbColumn(columnName string) {
// 	f.dbColumn = columnName
// }
//
// func (f *BigAutoField) setDbColumn(columnName string) {
// 	f.dbColumn = columnName
// }

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

// func (f *ForeignKey) setDbColumn(columnName string) {
// 	f.dbColumn = columnName
// }
//
// func (f OneToOneField) setDbColumn(columnName string) {
// 	f.dbColumn = columnName
// }
//
// func (f *AutoField) getDbType() string {
// 	return f.dbType
// }
//
// func (f *BigAutoField) getDbType() string {
// 	return f.dbType
// }

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

// func (f *ForeignKey) getDbType() string {
// 	return f.dbType
// }
//
// func (f *OneToOneField) getDbType() string {
// 	return f.dbType
// }

// func (f *AutoField) IsNil() bool {
// 	return !f.valid
// }
//
// func (f *BigAutoField) IsNil() bool {
// 	return f.valid
// }

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

// func (f *ForeignKey) IsNil() bool {
// 	return !f.valid
// }
//
// func (f *OneToOneField) IsNil() bool {
// 	return !f.valid
// }

// func (f *AutoField) valueAsSql() string {
// 	if f.IsNil() {
// 		return "NULL"
// 	} else {
// 		return int32AsSql(f.Value)
// 	}
// }
//
// func (f *BigAutoField) valueAsSql() string {
// 	if f.IsNil() {
// 		return "NULL"
// 	} else {
// 		return int64AsSql(f.Value)
// 	}
// }

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

// func (f *ForeignKey) valueAsSql() string {
// 	if f.IsNil() {
// 		return "NULL"
// 	} else {
// 		return int64AsSql(f.Value)
// 	}
// }
//
// func (f *OneToOneField) valueAsSql() string {
// 	if f.IsNil() {
// 		return "NULL"
// 	} else {
// 		return int64AsSql(f.Value)
// 	}
// }
