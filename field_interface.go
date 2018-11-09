package gojang

import ()

type field interface {
	selectExpression

	//getters and setters
	// DbColumn() string
	// setDbColumn(string)
	ColumnName() string
	setColumnName(string)

	Model() *Model
	setModel(*Model)
	HasModel() bool

	Null() bool
	setNull(bool)

	Unique() bool
	setUnique(bool)

	PrimaryKey() bool
	setPrimaryKey(bool)

	HasRelation() bool

	//asAssignment() assignment
	Asc() orderByExpression
	Desc() orderByExpression
	Count(bool) *aggregate

	copyField() field

	validate()
	valueAsSql() string
}
