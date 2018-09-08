package models

import ()

//Primary Key Option
func (f Field) PrimaryKey(value bool) Field {
	f.primaryKey = value
	return f
}

//Unique Field Option
func (f Field) Unique(value bool) Field {
	f.unique = value
	return f
}

//Null Field Option
func (f Field) Null(value bool) Field {
	f.null = value
	return f
}

//Default Values
// func (f Field) Default() Field {
// 	f.defaultVal = value
// 	return f
// }
