package gojang

import ()

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

func NewForeignKey(to *Model, onDelete onDelete) *ForeignKey {
	field := &ForeignKey{}
	field.BigIntegerField = NewBigIntegerField()

	field.isRelation = true
	field.manyToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	return field
}
