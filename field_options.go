package gojang

import (
	"errors"
	"reflect"
	"strconv"
)

type fieldOptions struct {
	dbColumn   string
	primaryKey bool
	null       bool
	unique     bool
}

func lookupDbTableTag(structField reflect.StructField) string {
	dbTable, _ := structField.Tag.Lookup("dbTable")
	return dbTable
}

func lookupFieldTags(structField reflect.StructField) (fieldOptions, error) {
	tag := fieldOptions{}
	var err error

	tag.dbColumn = lookupDbColumnTag(structField)
	if tag.dbColumn == "" {
		tag.dbColumn = snakeCase(structField.Name)
	}

	tag.primaryKey, err = lookupPrimaryKeyTag(structField)
	if err != nil {
		return fieldOptions{}, err
	}

	tag.null, err = lookupNullTag(structField)
	if err != nil {
		return fieldOptions{}, err
	}

	tag.unique, err = lookupUniqueTag(structField)
	if err != nil {
		return fieldOptions{}, err
	}

	return tag, nil
}

func lookupDbColumnTag(structField reflect.StructField) string {
	dbColumn, _ := structField.Tag.Lookup("dbColumn")
	return dbColumn
}

func lookupPrimaryKeyTag(structField reflect.StructField) (bool, error) {
	pkeyTag, ok := structField.Tag.Lookup("primaryKey")

	if ok {
		pkey, err := strconv.ParseBool(pkeyTag)

		if err != nil {
			return false, errors.New("Invalid value for primaryKey tag")
		} else {
			return pkey, nil
		}
	}

	return false, nil
}

func lookupNullTag(structField reflect.StructField) (bool, error) {
	nullTag, ok := structField.Tag.Lookup("null")

	if ok {
		null, err := strconv.ParseBool(nullTag)

		if err != nil {
			return false, errors.New("Invalid value for null tag")
		} else {
			return null, nil
		}
	}

	return false, nil
}

func lookupUniqueTag(structField reflect.StructField) (bool, error) {
	uniqueTag, ok := structField.Tag.Lookup("unique")

	if ok {
		unique, err := strconv.ParseBool(uniqueTag)

		if err != nil {
			return false, errors.New("Invalid value for unique tag")
		} else {
			return unique, nil
		}
	}

	return false, nil
}

//maybe
func lookupBoolTag(tagKey string, structField reflect.StructField) (bool, error) {
	tagValue, ok := structField.Tag.Lookup(tagKey)

	if ok {
		tagBool, err := strconv.ParseBool(tagValue)

		if err != nil {
			return false, errors.New("Tag value is not a boolean")
		}

		return tagBool, nil
	}

	return false, nil
}
