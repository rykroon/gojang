package gojang

import (
	"errors"
	"reflect"
	"strconv"
)

type fieldTags struct {
	dbColumn   string
	primaryKey bool
	null       bool
	unique     bool
}

func processTags(structField reflect.StructField) (fieldTags, error) {
	tag := fieldTags{}
	var err error

	tag.dbColumn = processDbColumnTag(structField)

	tag.primaryKey, err = processPrimaryKeyTag(structField)
	if err != nil {
		return fieldTags{}, err
	}

	tag.null, err = processNullTag(structField)
	if err != nil {
		return fieldTags{}, err
	}

	tag.unique, err = processUniqueTag(structField)
	if err != nil {
		return fieldTags{}, err
	}

	return tag, nil
}

func processDbColumnTag(structField reflect.StructField) string {
	dbColumn, ok := structField.Tag.Lookup("dbColumn")

	if ok {
		return dbColumn
	}

	return snakeCase(structField.Name)
}

func processPrimaryKeyTag(structField reflect.StructField) (bool, error) {
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

func processNullTag(structField reflect.StructField) (bool, error) {
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

func processUniqueTag(structField reflect.StructField) (bool, error) {
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
