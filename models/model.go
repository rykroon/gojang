package models

import (
	"database/sql"
)

type Model struct {
	dbTable string
	Objects Manager
	fields  map[string]field

	db sql.DB

	//Meta
	//uniqueTogether []string
}

//Create a new Model
func NewModel(dbTable string) Model {
	m := Model{dbTable: dbTable}
	m.Objects.model = &m
	m.fields = make(map[string]field)

	pkey := AutoField().PrimaryKey(true)
	pkey.autoCreated = true
	m.AddField("id", pkey)
	return m
}

//Add a Field to the Model
func (m *Model) AddField(fieldName string, field field) {
	field.model = m

	_, found := m.fields[fieldName]

	if found { //check for duplicate field
		panic("Field name already exists")
	}

	if field.primaryKey {
		if m.hasPrimaryKey() {
			pkey := m.getPrimaryKey()

			if pkey.autoCreated {
				delete(m.fields, "id")
			} else {
				panic("Model already has a primary key")
			}
		}
	}

	if field.isRelation {
		if field.dbColumn == "" {
			field.dbColumn = fieldName + "_id"
		}

		if field.manyToOne { //Foreign Key
			//reverseFieldName :+ fieldName + "_set"
			//reverseField :=
			//reverseField.concrete = false //should proably set this to true for other situations
			//field.relatedModel.AddField(reverseFieldName)
		}
	}

	if field.dbColumn == "" {
		field.dbColumn = fieldName
	}

	m.fields[fieldName] = field
}

//Get a Field from the Model
func (m Model) Field(fieldName string) (field, bool) {
	field, ok := m.fields[fieldName]
	return field, ok
}

//Create a new Instance of this Model
func (m Model) NewInstance() Instance {
	i := Instance{}
	i.model = &m
	return i
}

func (m Model) getPrimaryKey() field {
	for _, field := range m.fields {
		if field.primaryKey {
			return field
		}
	}

	panic("Primary Key was not found")
}

//alias for getPrimaryKey()
func (m Model) PK() field {
	return m.getPrimaryKey()
}

//Checks if the Model has a Primary Key field
// ** Technically there should never be a reason why a model has no Primary Key
func (m Model) hasPrimaryKey() bool {
	for _, field := range m.fields {
		if field.primaryKey {
			return true
		}
	}
	return false
}

//returns a list of the Model's fields
func (m Model) fieldList() []field {
	list := []field{}

	for _, field := range m.fields {
		list = append(list, field)
	}

	return list
}

//returns a list of the model's fields in SQL format
func (m Model) sqlFieldList() []string {
	list := []string{}

	for _, field := range m.fields {
		list = append(list, field.toSql())
	}

	return list
}

func (m *Model) Migrate() {

}

//Returns an SQL Query that will create a new table
func (m Model) CreateTable() string {
	s := "CREATE TABLE " + m.dbTable + "("

	for _, field := range m.fields {
		s += field.create() + ", "
	}

	s = s[0:len(s)-2] + ");"

	return s
}
