package models

import (
	"strconv"
)

type dbDataType struct {
	dataType  string
	n         int //arbitray n
	precision int //precision
	scale     int //scale
}

func (d *dbDataType) setPrecision(p int) {
	d.precision = p
}

func (d *dbDataType) setScale(s int) {
	d.scale = s
}

func (d *dbDataType) setN(n int) {
	d.n = n
}

func (d dbDataType) String() string {

	switch d.dataType {
	case "VARCHAR":
		n := "(" + strconv.Itoa(d.n) + ")"
		return d.dataType + n

	case "NUMERIC":
		if d.precision == 0 && d.scale == 0 {
			return d.dataType

		} else if d.precision != 0 && d.scale == 0 {
			p := "(" + strconv.Itoa(d.precision) + ")"
			return d.dataType + p

		} else if d.precision != 0 && d.scale != 0 {
			p := strconv.Itoa(d.precision)
			s := strconv.Itoa(d.scale)
			return d.dataType + "(" + p + ", " + s + ")"
		}
	}

	return d.dataType
}

// bigInt := dbDataType{dataType:"BIGINT"}
// bigSerial := dbDataType{dataType:"BIGSERIAL"}
// boolean := dbDataType{dataType:"BOOLEAN"}
// varChar := dbDataType{dataType:"VARCHAR"}
// doublePrecision := dbDataType{dataType:"DOUBLE PRECISION"}
// integer := dbDataType{dataType:"INTEGER"}
// numeric := dbDataType{dataType:"NUMERIC"}
// serial := dbDataType{dataType:"SERIAL"}
// text := dbDataType{dataType:"TEXT"}
