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
func (m *Model) AddField(fieldName string, f Field) {
	if f.foreignKey && (f.dbColumn == "") {
		f.dbColumn = doubleQuotes(fieldName + "_id")
	}

	if f.dbColumn == "" {
		f.dbColumn = doubleQuotes(fieldName)
	}

	m.fields[fieldName] = f
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

//Checks if the Model as a Primary Key Field
func (m Model) hasPrimaryKey() bool {
	for _, field := range m.fields {
		if field.primaryKey {
			return true
		}
	}

	return false
}

//returns a list of the Model's field names
func (m Model) fieldList() []string {
	list := []string{}

	for key, _ := range m.fields {
		list = append(list, key)
	}

	return list
}


func (m *Model) Migrate() {
	if !m.hasPrimaryKey() {
		m.AddField("id", AutoField().PrimaryKey(true))
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
