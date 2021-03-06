package gojang

import ()

// type stringField interface {
// 	Val() string
// 	set(string)
// }

type numericField interface {
	field
	Avg() *aggregate
	Sum() *aggregate
}

type intField interface {
	numericField
	Val() int
	setInt(int)

	//lookups
	Exact(int) lookup
	In(...int) lookup
	Gt(int) lookup
	Gte(int) lookup
	Lt(int) lookup
	Lte(int) lookup
	Range(int, int) lookup
}

type primaryKeyField interface {
	intField
	Id() int
	isAutoField() bool
}

func (f *AutoField) Id() int {
	return f.Val()
}

func (f *BigAutoField) Id() int {
	return f.Val()
}

func (f *AutoField) isAutoField() bool {
	return true
}

func (f *BigAutoField) isAutoField() bool {
	return true
}

func (f *BigIntegerField) setInt(num int) {
	f.Value = int64(num)
}

func (f *IntegerField) setInt(num int) {
	f.Value = int32(num)
}

func (f *SmallIntegerField) setInt(num int) {
	f.Value = int16(num)
}

func (f BigIntegerField) Val() int {
	return int(f.Value)
}

func (f IntegerField) Val() int {
	return int(f.Value)
}

func (f SmallIntegerField) Val() int {
	return int(f.Value)
}
