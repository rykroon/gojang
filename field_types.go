package gojang

import ()

type onDelete string

const Cascade onDelete = "CASCADE"
const Protect onDelete = "RESTRICT"
const SetNull onDelete = "SET NULL"
const SetDefault onDelete = "SET DEFAULT"

type constraints struct {
	null       bool
	unique     bool
	primaryKey bool
}

type column struct {
	model    *Model
	dbColumn string
	dbType   string
	alias    string

	constraints
	isRelation bool //foreignKey
}

//A field type for each data type
type BigIntegerField struct {
	*column

	valid bool
	Value int64
}

type BooleanField struct {
	*column

	valid bool
	Value bool
}

type FloatField struct {
	*column

	valid bool
	Value float64
}

type IntegerField struct {
	*column

	valid bool
	Value int32
}

type SmallIntegerField struct {
	*column

	valid bool
	Value int16
}

type TextField struct {
	*column

	valid bool
	Value string
}

type AutoField struct {
	*IntegerField
}

type BigAutoField struct {
	*BigIntegerField
}

type ForeignKey struct {
	*BigIntegerField

	//specific for related fields
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model
	onDelete     onDelete
}

type OneToOneField struct {
	*ForeignKey
}

//Constructors

func NewColumn(dbType string) *column {
	return &column{dbType: dbType}
}

func NewBigIntegerField() *BigIntegerField {
	field := &BigIntegerField{}
	field.column = NewColumn("INT8")
	field.valid = true
	return field
}

func NewBooleanField() *BooleanField {
	field := &BooleanField{}
	field.column = NewColumn("BOOL")
	field.valid = true
	return field
}

func NewFloatField() *FloatField {
	field := &FloatField{}
	field.column = NewColumn("FLOAT8")
	field.valid = true
	return field
}

func NewIntegerField() *IntegerField {
	field := &IntegerField{}
	field.column = NewColumn("INT4")
	field.valid = true
	return field
}

func NewSmallIntegerField() *SmallIntegerField {
	field := &SmallIntegerField{}
	field.column = NewColumn("INT2")
	field.valid = true
	return field
}

func NewTextField() *TextField {
	field := &TextField{}
	field.column = NewColumn("TEXT")
	field.valid = true
	return field
}

func NewAutoField() *AutoField {
	field := &AutoField{}
	field.IntegerField = NewIntegerField()
	field.dbType = "SERIAL4"
	return field
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{}
	field.BigIntegerField = NewBigIntegerField()
	field.dbType = "SERIAL8"
	return field
}

func NewForeignKey(to *Model, onDelete onDelete) *ForeignKey {
	field := &ForeignKey{}
	field.BigIntegerField = NewBigIntegerField()

	field.isRelation = true
	field.manyToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	return field
}

func NewOneToOneField(to *Model, onDelete onDelete) *OneToOneField {
	field := &OneToOneField{}
	field.ForeignKey = NewForeignKey(to, onDelete)

	field.manyToOne = false
	field.oneToOne = true

	//unique constraint must be true for OneToOne Field
	field.unique = true

	return field
}

func (f *BigIntegerField) copy() *BigIntegerField {
	copy := NewBigIntegerField()
	copy.model = f.model //maybe change to be a copy of the model
	copy.dbColumn = f.dbColumn
	copy.alias = f.alias
	copy.constraints = f.constraints
	return copy
}

func (f *BooleanField) copy() *BooleanField {
	copy := NewBooleanField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.alias = f.alias
	copy.constraints = f.constraints
	return copy
}

func (f *FloatField) copy() *FloatField {
	copy := NewFloatField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.alias = f.alias
	copy.constraints = f.constraints
	return copy
}

func (f *IntegerField) copy() *IntegerField {
	copy := NewIntegerField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.alias = f.alias
	copy.constraints = f.constraints
	return copy
}

func (f *SmallIntegerField) copy() *SmallIntegerField {
	copy := NewSmallIntegerField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.alias = f.alias
	copy.constraints = f.constraints
	return copy
}

func (f *TextField) copy() *TextField {
	copy := NewTextField()
	copy.model = f.model
	copy.dbColumn = f.dbColumn
	copy.alias = f.alias
	copy.constraints = f.constraints
	return copy
}

func (f *BigIntegerField) SetNil() error {
	if f.HasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *BooleanField) SetNil() error {
	if f.HasNullConstraint() {
		f.valid = false
		f.Value = false
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *FloatField) SetNil() error {
	if f.HasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *IntegerField) SetNil() error {
	if f.HasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *SmallIntegerField) SetNil() error {
	if f.HasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *TextField) SetNil() error {
	if f.HasNullConstraint() {
		f.valid = false
		f.Value = ""
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f BigIntegerField) UnSetNil() {
	f.valid = true
}

func (f BooleanField) UnSetNil() {
	f.valid = true
}

func (f FloatField) UnSetNil() {
	f.valid = true
}

func (f IntegerField) UnSetNil() {
	f.valid = true
}

func (f SmallIntegerField) UnSetNil() {
	f.valid = true
}

func (f TextField) UnSetNil() {
	f.valid = true
}

func create(f field) string {
	s := dbq(f.DbColumn()) + " " + f.DbType()

	if f.HasPrimaryKeyConstraint() {
		s += " PRIMARY KEY"
	} else {

		if f.HasRelation() {
			fkey := f.(relatedField)
			s += " REFERENCES " + dbq(fkey.getRelatedModel().dbTable) + " ON DELETE " + fkey.getOnDelete()
		}

		if f.HasNullConstraint() {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.HasUniqueConstraint() {
			s += " UNIQUE"
		}
	}

	return s
}
