package gojang

import (
//"fmt"
)

type relatedField interface {
	isManyToMany() bool
	isManyToOne() bool
	isOneToMany() bool
	isOneToOne() bool
	getRelatedModel() *Model
	getOnDelete() string
}

func (f *ForeignKey) isManyToMany() bool {
	return f.manyToMany
}

func (f *OneToOneField) isManyToMany() bool {
	return f.manyToMany
}

func (f *ForeignKey) isManyToOne() bool {
	return f.manyToOne
}

func (f *OneToOneField) isManyToOne() bool {
	return f.manyToOne
}

func (f *ForeignKey) isOneToMany() bool {
	return f.oneToMany
}

func (f *OneToOneField) isOneToMany() bool {
	return f.oneToMany
}

func (f *ForeignKey) isOneToOne() bool {
	return f.oneToOne
}

func (f *OneToOneField) isOneToOne() bool {
	return f.oneToOne
}

func (f *ForeignKey) getRelatedModel() *Model {
	return f.relatedModel
}

func (f *OneToOneField) getRelatedModel() *Model {
	return f.relatedModel
}

func (f *ForeignKey) getOnDelete() string {
	return string(f.onDelete)
}

func (f *OneToOneField) getOnDelete() string {
	return string(f.onDelete)
}

// func (f *ForeignKey) Fetch() error {
// 	model := f.relatedModel
// 	return model.Objects.Get(model.Pk.Exact(f.Val()))
// }
//
// func (f *OneToOneField) Fetch() error {
// 	model := f.relatedModel
// 	return model.Objects.Get(model.Pk.Exact(f.Val()))
// }
