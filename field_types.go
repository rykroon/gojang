package gojang

import ()

//A field type for each data type

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
