package gojang

import (
	"fmt"
)

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

type Column struct {
	model    *Model
	dbColumn string
	dbType   string
	alias    string

	constraints
	isRelation bool //foreignKey
}

//A field type for each data type
type BigIntegerField struct {
	*Column

	Valid bool
	Value int64
}

type BooleanField struct {
	*Column

	Valid bool
	Value bool
}

type FloatField struct {
	*Column

	Valid bool
	Value float64
}

type IntegerField struct {
	*Column

	Valid bool
	Value int32
}

type SmallIntegerField struct {
	*Column

	Valid bool
	Value int16
}

type TextField struct {
	*Column

	Valid bool
	Value string
}

type AutoField struct {
	*IntegerField
}

type BigAutoField struct {
	*BigIntegerField
}

type CharField struct {
	*TextField
	maxLength int
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

func newColumn(dbType string) *Column {
	return &Column{dbType: dbType}
}

func NewBigIntegerField() *BigIntegerField {
	field := &BigIntegerField{}
	field.Column = newColumn("INT8")
	field.Valid = true
	return field
}

func NewBooleanField() *BooleanField {
	field := &BooleanField{}
	field.Column = newColumn("BOOL")
	field.Valid = true
	return field
}

func NewCharField(maxLength int) *CharField {
	field := &CharField{maxLength: maxLength}
	field.TextField = NewTextField()
	dbType := fmt.Sprintf("VARCHAR(%v)", maxLength)
	field.dbType = dbType
	return field
}

func NewFloatField() *FloatField {
	field := &FloatField{}
	field.Column = newColumn("FLOAT8")
	field.Valid = true
	return field
}

func NewIntegerField() *IntegerField {
	field := &IntegerField{}
	field.Column = newColumn("INT4")
	field.Valid = true
	return field
}

func NewSmallIntegerField() *SmallIntegerField {
	field := &SmallIntegerField{}
	field.Column = newColumn("INT2")
	field.Valid = true
	return field
}

func NewTextField() *TextField {
	field := &TextField{}
	field.Column = newColumn("TEXT")
	field.Valid = true
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

func (c *Column) copy() *Column {
	copy := newColumn(c.dbType)
	copy.model = c.model
	copy.dbColumn = c.dbColumn
	copy.alias = c.alias
	copy.constraints = c.constraints
	return copy
}

func (f *BigIntegerField) copy() *BigIntegerField {
	copy := NewBigIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *BooleanField) copy() *BooleanField {
	copy := NewBooleanField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *FloatField) copy() *FloatField {
	copy := NewFloatField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *IntegerField) copy() *IntegerField {
	copy := NewIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *SmallIntegerField) copy() *SmallIntegerField {
	copy := NewSmallIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *TextField) copy() *TextField {
	copy := NewTextField()
	copy.Column = f.Column.copy()
	return copy
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
