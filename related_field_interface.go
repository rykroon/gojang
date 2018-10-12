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

func (f *ForeignKeyField) isManyToMany() bool {
	return f.manyToMany
}

func (f *OneToOneField) isManyToMany() bool {
	return f.manyToMany
}

func (f *ForeignKeyField) isManyToOne() bool {
	return f.manyToOne
}

func (f *OneToOneField) isManyToOne() bool {
	return f.manyToOne
}

func (f *ForeignKeyField) isOneToMany() bool {
	return f.oneToMany
}

func (f *OneToOneField) isOneToMany() bool {
	return f.oneToMany
}

func (f *ForeignKeyField) isOneToOne() bool {
	return f.oneToOne
}

func (f *OneToOneField) isOneToOne() bool {
	return f.oneToOne
}

func (f *ForeignKeyField) getRelatedModel() *Model {
	return f.relatedModel
}

func (f *OneToOneField) getRelatedModel() *Model {
	return f.relatedModel
}

func (f *ForeignKeyField) getOnDelete() string {
	return string(f.onDelete)
}

func (f *OneToOneField) getOnDelete() string {
	return string(f.onDelete)
}

func (f *ForeignKeyField) Fetch() error {
	model := f.relatedModel

	table := dbq(model.dbTable)
	pkCol := model.Pk.(field).sqlField()

	//when queryset logic is finished use that to build the string
	sql := "SELECT * FROM " + table + " WHERE " + pkCol + " = " + f.sqlValue() + ";"
	rows, err := model.db.Query(sql)

	if err != nil {
		return err
	}

	columns, err := rows.Columns()

	if err != nil {
		return err
	}

	pointers := model.getPointers(columns)

	for rows.Next() {
		err := rows.Scan(pointers...)

		if err != nil {
			return err
		}

		break
	}

	return nil
}
