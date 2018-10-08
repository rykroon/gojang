package gojang

import ()

type relatedField interface {
	isManyToMany() bool
	isManyToOne() bool
	isOneToMany() bool
	isOneToOne() bool
  getRelatedModel() Model
  getOnDelete() string
}

func (f *ForeignKeyField) isManyToMany() bool {
	return f.manyToMany
}

func (f *ForeignKeyField) isManyToOne() bool {
	return f.manyToOne
}

func (f *ForeignKeyField) isOneToMany() bool {
	return f.oneToMany
}

func (f *ForeignKeyField) isOneToOne() bool {
	return f.oneToOne
}

func (f *ForeignKeyField) getRelatedModel() Model {
	return f.relatedModel
}

func (f *ForeignKeyField) getOnDelete() string {
	return string(f.onDelete)
}

func (f *ForeignKeyField) Fetch() string {
  model := f.relatedModel
  pkField := dbq(model.getPKField().(field).DBColumn())
  sql := "SELECT * FROM " + dbq(model.DBTable) + " WHERE " + pkField + " = " + f.sqlValue()

  return sql
}
