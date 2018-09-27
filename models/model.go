package models

import (
	"database/sql"
)

type Model struct {
	dbTable string
	Objects Manager
	fields  map[string]field

	db *sql.DB

	//Meta
	//uniqueTogether []string
}

//Create a new Model
func NewModel(dbTable string, db Database) Model {
	sqlDB, err := db.toDB()

	if err != nil{
		panic(err)
	}

	m := Model{dbTable: dbTable, db: sqlDB}
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

//Returns a map with attribute names as the keys and database columns as the values
func (m Model) attrToDBColumnMap() map[string]string {
	result := make(map[string]string)

	for key, field := range m.fields {
		result[key] = field.dbColumn
	}

	return result
}

//Returns a map of database columns as the keys and attribute names as the values
func (m Model) dbColumnToAttrMap() map[string]string {
	result := make(map[string]string)

	for key, field := range m.fields {
		result[field.dbColumn] = key
	}

	return result
}



//Creates a table
func (m Model) CreateTable() {
	sql := "CREATE TABLE IF NOT EXISTS " + m.dbTable + "("

	for _, field := range m.fields {
		sql += field.create() + ", "
	}

	sql = sql[0:len(sql)-2] + ");"

	//return s
	//fmt.Println(sql)
	_, err := m.db.Exec(sql)

	if err != nil {
		panic(err)
	} else {
		//fmt.Println(result.LastInsertId())
		//fmt.Println(result.RowsAffected())
	}


}
