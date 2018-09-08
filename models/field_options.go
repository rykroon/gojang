package models

import ()

type fieldOption func(*Field)

func PrimaryKey(pkey bool) fieldOption {
	return func(f *Field) {
		f.primaryKey = pkey
	}
}

func Unique(unique bool) fieldOption {
	return func(f *Field) {
		f.unique = unique
	}
}

func Null(null bool) fieldOption {
	return func(f *Field) {
		f.null = null
	}
}





//Add Primary Key Option
func (f AutoField) PrimaryKey(value bool) AutoField {
	f.primaryKey = value
	return f
}

func (f BooleanField) PrimaryKey(value bool) BooleanField {
	f.primaryKey = value
	return f
}

func (f CharField) PrimaryKey(value bool) CharField {
	f.primaryKey = value
	return f
}

func (f DecimalField) PrimaryKey(value bool) DecimalField {
	f.primaryKey = value
	return f
}

func (f FloatField) PrimaryKey(value bool) FloatField {
	f.primaryKey = value
	return f
}

func (f IntegerField) PrimaryKey(value bool) IntegerField {
	f.primaryKey = value
	return f
}

func (f TextField) PrimaryKey(value bool) TextField {
	f.primaryKey = value
	return f
}



//Add Unique Field Option
func (f AutoField) Unique(value bool) AutoField {
	f.unique = value
	return f
}

func (f BooleanField) Unique(value bool) BooleanField {
	f.unique = value
	return f
}

func (f CharField) Unique(value bool) CharField {
	f.unique = value
	return f
}

func (f DecimalField) Unique(value bool) DecimalField {
	f.unique = value
	return f
}

func (f FloatField) Unique(value bool) FloatField {
	f.unique = value
	return f
}

func (f IntegerField) Unique(value bool) IntegerField {
	f.unique = value
	return f
}

func (f TextField) Unique(value bool) TextField {
	f.unique = value
	return f
}



//Add Null Field Option
func (f AutoField) Null(value bool) AutoField {
	f.null = value
	return f
}

func (f BooleanField) Null(value bool) BooleanField {
	f.null = value
	return f
}

func (f CharField) Null(value bool) CharField {
	f.null = value
	return f
}

func (f DecimalField) Null(value bool) DecimalField {
	f.null = value
	return f
}

func (f FloatField) Null(value bool) FloatField {
	f.null = value
	return f
}

func (f IntegerField) Null(value bool) IntegerField {
	f.null = value
	return f
}

func (f TextField) Null(value bool) TextField {
	f.null = value
	return f
}


//Default Values
func (f BooleanField) Default(value bool) BooleanField {
	f.defaultVal = value
	return f
}

func (f CharField) Default(value string) CharField {
	f.defaultVal = value
	return f
}

func (f DecimalField) Default(value float64) DecimalField {
	f.defaultVal = value
	return f
}

func (f FloatField) Default(value float64) FloatField {
	f.defaultVal = value
	return f
}

func (f IntegerField) Default(value int) IntegerField {
	f.defaultVal = value
	return f
}

func (f TextField) Default(value string) TextField {
	f.defaultVal = value
	return f
}
