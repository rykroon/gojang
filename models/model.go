package models

import ()

type Model struct {
	dbTable string
	Objects Manager
	fields  map[string]Field

	//Meta
	//uniqueTogether []string
}

func NewModel(dbTable string) Model {
	m := Model{dbTable: dbTable}
	m.Objects.model = &m
	m.fields = make(map[string]Field)
	return m
}

func (m *Model) AddField(fieldName string, f Field) {
	if f.foreignKey && (f.dbColumn == "") {
		f.dbColumn = fieldName + "_id"
	}

	if f.dbColumn == "" {
		f.dbColumn = fieldName
	}

	m.fields[fieldName] = f
}

func (m Model) Field(fieldName string) (Field, bool) {
	field, exists := m.fields[fieldName]
	return field, exists
}

func (m Model) NewInstance() Instance {
	i := Instance{}
	i.model = &m
	return i
}

func (m Model) hasPrimaryKey() bool {
	for _, field := range m.fields {
		if field.primaryKey {
			return true
		}
	}

	return false
}

func (m *Model) Migrate() {
	if !m.hasPrimaryKey() {
		m.AddField("id", AutoField().PrimaryKey(true))
	}

	//sql := m.CreateTable()

}

func (m Model) CreateTable() string {
	s := "CREATE TABLE " + m.dbTable + "("

	for _, field := range m.fields {
		s += field.create() + ", "
	}

	s = s[0:len(s)-2] + ");"

	return s
}
