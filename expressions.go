package gojang

import ()

type expression interface {
	asExpr() string
}

// type sortExpression interface {
//   Asc() string
//   Desc() string
// }

func (f *AutoField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BigAutoField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BigIntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BooleanField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *FloatField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *IntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *SmallIntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *TextField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *ForeignKey) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *OneToOneField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *AutoField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BigAutoField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BigIntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BooleanField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *FloatField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *IntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *SmallIntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *TextField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *ForeignKey) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *OneToOneField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}
