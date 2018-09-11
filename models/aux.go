package models

import (
	"reflect"
	"strconv"
	"database/sql"
)

//auxillary functions used in all programs

func doubleQuotes(s string) string {
	return "\"" + s + "\""
}

func singleQuotes(s string) string {
	return "'" + s + "'"
}


func interfaceSliceToSql(slice []interface{}) string {
	s := "("

	for _, value := range slice {
		s += interfaceToSql(value) + ", "
	}

	s = s[0: len(s) - 2] + ")"
	return s
}

func interfaceToSql(i interface{}) string {
	t := reflect.TypeOf(i).String()
	k := reflect.TypeOf(i).Kind()

	switch k {
	case reflect.Slice:
		switch t {
		case "[]interface {}":
			slice := i.([]interface{})
			return interfaceSliceToSql(slice)

		case "[]string":
			slice := i.([]string)
			return sliceStringToSql(slice)

		case "[]int":
			slice := i.([]int)
			return sliceIntToSql(slice)
		}
	}

	switch t {
	case "sql.NullBool":
		nullValue := i.(sql.NullBool)

		if nullValue.Valid {
			return boolToSql(nullValue.Bool)
		} else {
			return "NULL"
		}

	case "sql.NullFloat64":
		nullValue := i.(sql.NullFloat64)

		if nullValue.Valid {
			return float64ToSql(nullValue.Float64)
		} else {
			return "NULL"
		}

	case "sql.NullInt64":
		nullValue := i.(sql.NullInt64)

		if nullValue.Valid {
			return int64ToSql(nullValue.Int64)
		} else {
			return "NULL"
		}

	case "sql.NullString":
		nullValue := i.(sql.NullString)

		if nullValue.Valid {
			return stringToSql(nullValue.String)
		} else {
			return "NULL"
		}

	case "int":
		return intToSql(i.(int))

	case "int32":
		return int32ToSql(i.(int32))

	case "int64":
		return int64ToSql(i.(int64))

	case "string":
		return stringToSql(i.(string))

	case "float64":
		return float64ToSql(i.(float64))

	case "bool":
		return boolToSql(i.(bool))
	}

	return ""
}


func sliceStringToSql(slice []string) string {
	s := "("

	for _, value := range slice {
		s += stringToSql(value) + ", "
	}

	s = s[0: len(s) - 2] + ")"

	return s
}

func stringToSql(s string) string {
	return singleQuotes(s)
}


func sliceIntToSql(slice []int) string {
	s := "("

	for _, value := range slice {
		s += intToSql(value) + ", "
	}

	s = s[0: len(s) - 2] + ")"

	return s
}

func intToSql(i int) string {
	return strconv.Itoa(i)
}


func int32ToSql(i int32) string {
	return strconv.Itoa(int(i))
}


func int64ToSql(i int64) string {
	return strconv.Itoa(int(i))
}


func float64ToSql(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}


func boolToSql(b bool) string {
	if b {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

//Constants

const Cascade string = "CASCADE"
const Protect string = "RESTRICT"
const SetNull string = "SET NULL"
const SetDefault string = "SET DEFAULT"
