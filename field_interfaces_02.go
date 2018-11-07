package gojang

import ()

type stringField interface {
	Val() string
	set(string)

	Contains(string) lookup
	IContains(string) lookup
	Exact(string) lookup
	IExact(string) lookup
	In(...string) lookup
	Gt(string) lookup
	Gte(string) lookup
	Lt(string) lookup
	Lte(string) lookup
	Range(string, string) lookup
	StartsWith(string) lookup
	IStartsWith(string) lookup
	EndsWith(string) lookup
	IEndsWith(string) lookup

	Length() *IntegerField
	Upper() *TextField
	Lower() *TextField
}

type NumericField interface {
	field
	Avg() *aggregate
	Sum() *aggregate
}

type IntField interface {
	NumericField
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

type PrimaryKeyField interface {
	IntField
	Id() int
	isAutoField() bool
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
