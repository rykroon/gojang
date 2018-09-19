package models

import ()

type Model struct {
	dbTable string
	Objects Manager
	fields  map[string]Field

	//Meta
	//uniqueTogether []string
}

//Create a new Model
func NewModel(dbTable string) Model {
	m := Model{dbTable: dbTable}
	m.Objects.model = &m
	m.fields = make(map[string]Field)
	return m
}

//Add a Field to the Model
func (m *Model) AddField(fieldName string, field Field) {
	field.model = m

	if field.primaryKey {
		if m.hasPrimaryKey() {
			panic("Model already has a primary key")
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
func (m Model) Field(fieldName string) (Field, bool) {
	field, ok := m.fields[fieldName]
	return field, ok
}

//Create a new Instance of this Model
func (m Model) NewInstance() Instance {
	i := Instance{}
	i.model = &m
	return i
}

func (m Model) getPrimaryKey() Field {
	for _, field := range m.fields {
		if field.primaryKey {
			return field
		}
	}
	return Field{}
}

//Checks if the Model as a Primary Key Field
func (m Model) hasPrimaryKey() bool {
	for _, field := range m.fields {
		if field.primaryKey {
			return true
		}
	}
	return false
}

//returns a list of the Model's fields
func (m Model) fieldList() []Field {
	list := []Field{}

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
	if !m.hasPrimaryKey() {
		pkey := AutoField().PrimaryKey(true)
		pkey.autoCreated = true
		m.AddField("id", pkey)
	}

	//sql := m.CreateTable()

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
