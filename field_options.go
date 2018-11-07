package gojang

import (
	"reflect"
)

type fieldOptions struct {
	dbColumn   string
	primaryKey bool
	null       bool
	unique     bool
}

func getFieldOptions(structField reflect.StructField) (fieldOptions, error) {
	options, err := lookupFieldOptionTags(structField.Tag)
	if err != nil {
		return options, err
	}

	if options.dbColumn == "" {
		options.dbColumn = snakeCase(structField.Name)
	}

	return options, nil
}

func setFieldOptions(field field, options fieldOptions) {
	field.setColumnName(options.dbColumn)
	field.setPrimaryKeyConstraint(options.primaryKey)
	field.setNullConstraint(options.null)
	field.setUniqueConstraint(options.unique)
}

func lookupFieldOptionTags(tag reflect.StructTag) (fieldOptions, error) {
	options := fieldOptions{}
	var err error

	options.dbColumn = lookupDbColumnTag(tag)

	options.primaryKey, err = lookupPrimaryKeyTag(tag)
	if err != nil {
		return fieldOptions{}, err
	}

	options.null, err = lookupNullTag(tag)
	if err != nil {
		return fieldOptions{}, err
	}

	options.unique, err = lookupUniqueTag(tag)
	if err != nil {
		return fieldOptions{}, err
	}

	return options, nil
}
