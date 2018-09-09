package models

import (
	"reflect"
	//"strconv"
	"strings"
)

type Lookup struct {
	lhs        string
	lookupName string
	rhs        interface{}
}

func lookupToSql(fieldLookup string, i interface{}) string {
	idx := strings.Index(fieldLookup, "__")
	field := ""
	lookup := ""

	//If no Lookup then default to 'exact'
	if idx < 0 {
		field = fieldLookup
		lookup = "exact"
	} else {
		field = fieldLookup[0:idx]
		lookup = fieldLookup[idx+2:]
	}

	value := reflect.ValueOf(i)

	expression := field

	switch lookup {
	case "exact":
		expression += exact(value)

	case "iexact":
		expression += iexact(value)

	case "contains":
		expression += contains(value)

	case "icontains":
		expression += icontains(value)

	case "in":
		expression += in(value)

	case "gt":
		expression += gt(value)

	case "gte":
		expression += gte(value)

	case "lt":
		expression += lt(value)

	case "lte":
		expression += lte(value)

	case "startswith":
		expression += startsWith(value)

	case "istartswith":
		expression += iStartsWith(value)

	case "endswith":
		expression += endsWith(value)

	case "iendswith":
		expression += iEndsWith(value)

	case "isnull":
		expression += isNull(value)

	}

	return expression
}

//Process Right Hand Side
// func valueToSql(v reflect.Value) string {
// 	t := v.Type().String()
// 	k := v.Kind().String()
//
// 	switch k {
// 	case "slice":
// 		str := "("
// 		for idx := 0; idx < v.Len(); idx++ {
// 			element := v.Index(idx)
//
// 			if t == "[]string" {
// 				str += singleQuotes(valueToSql(element)) + ", "
// 			} else {
// 				str += valueToSql(element) + ", "
// 			}
// 		}
//
// 		return str[0:len(str)-2] + ")"
//
// 	case "struct":
// 		if t == "models.QuerySet" {
// 			subQuery := v.Interface().(QuerySet)
// 			return "(" + subQuery.Query[0:len(subQuery.Query)-1] + ")"
// 		}
//
// 	default: //Primitive Types
//
// 		switch t {
// 		case "int":
// 			return strconv.Itoa(v.Interface().(int))
//
// 		case "string":
// 			return v.Interface().(string)
//
// 		case "bool":
// 			if v.Interface().(bool) {
// 				return "TRUE"
// 			} else {
// 				return "FALSE"
// 			}
// 		}
// 	}
//
// 	return ""
// }

//Helper function for logical operators
func logicalOperator(v reflect.Value, logOp string) string {
	value := valueToSql(v)
	t := v.Type().String()

	if t == "string" {
		value = singleQuotes(value)
	}

	return " " + logOp + " " + value
}

//Helper Function for pattern matching
func patternMatch(v reflect.Value, caseInsensitive bool, startPattern string, endPattern string) string {
	value := valueToSql(v)

	if caseInsensitive {
		return " ILIKE " + singleQuotes(startPattern+value+endPattern)
	} else {
		return " LIKE " + singleQuotes(startPattern+value+endPattern)
	}
}

//Lookup functions
func exact(v reflect.Value) string {
	//if value is nil then call isnull()
	return logicalOperator(v, "=")
}

func iexact(v reflect.Value) string {
	return patternMatch(v, true, "", "")
}

func contains(v reflect.Value) string {
	return patternMatch(v, false, "%", "%")
}

func icontains(v reflect.Value) string {
	return patternMatch(v, true, "%", "%")
}

//check if interface{} is slice or queryset
func in(v reflect.Value) string {
	return " IN " + valueToSql(v)
}

func gt(v reflect.Value) string {
	return logicalOperator(v, ">")
}

func gte(v reflect.Value) string {
	return logicalOperator(v, ">=")
}

func lt(v reflect.Value) string {
	return logicalOperator(v, "<")
}

func lte(v reflect.Value) string {
	return logicalOperator(v, "<=")
}

func startsWith(v reflect.Value) string {
	return patternMatch(v, false, "", "%")
}

func iStartsWith(v reflect.Value) string {
	return patternMatch(v, true, "", "%")
}

func endsWith(v reflect.Value) string {
	return patternMatch(v, false, "%", "")
}

func iEndsWith(v reflect.Value) string {
	return patternMatch(v, true, "%", "")
}

func isNull(v reflect.Value) string {
	t := v.Type().String()

	if t == "bool" {
		if v.Interface().(bool) {
			return " IS NULL"
		} else {
			return " IS NOT NULL"
		}
	}

	return ""
}
