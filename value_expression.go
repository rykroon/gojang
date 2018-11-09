package gojang

import (
	"fmt"
	"strconv"
	"strings"
)

// type ValueExpression struct {
// 	outputField field
// }
//
// func NewValueFromBool(value bool) *ValueExpression {
// 	valExpr := &ValueExpression{}
// 	field := NewBooleanField()
// 	field.Value = value
// 	valExpr.outputField = field
// 	return valExpr
// }
//
// func NewValueFromFloat(value float64) *ValueExpression {
// 	valExpr := &ValueExpression{}
// 	field := NewFloatField()
// 	field.Value = value
// 	valExpr.outputField = field
// 	return valExpr
// }
//
// func NewValueFromInt(value int) *ValueExpression {
// 	valExpr := &ValueExpression{}
// 	field := NewBigIntegerField()
// 	field.Value = int64(value)
// 	valExpr.outputField = field
// 	return valExpr
// }
//
// func NewValueFromString(value string) *ValueExpression {
// 	valExpr := &ValueExpression{}
// 	field := NewTextField()
// 	field.Value = value
// 	valExpr.outputField = field
// 	return valExpr
// }

func boolAsSql(b bool) string {
	if b {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func intAsSql(i int) string {
	return strconv.Itoa(i)
}

func float64AsSql(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func stringAsSql(s string) string {
	return fmt.Sprintf("'%v'", s)
}

func boolsAsSql(booleans []bool) string {
	var valueList []string
	for _, value := range booleans {
		valueList = append(valueList, boolAsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}

func floatsAsSql(floats []float64) string {
	var valueList []string
	for _, value := range floats {
		valueList = append(valueList, float64AsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}

func integersAsSql(integers []int) string {
	var valueList []string
	for _, value := range integers {
		valueList = append(valueList, intAsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}

func stringsAsSql(stringList []string) string {
	var valueList []string
	for _, value := range stringList {
		valueList = append(valueList, stringAsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}
