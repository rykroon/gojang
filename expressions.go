package gojang

import ()

//Expressions describe a value or a computation that can be used as part of an
//update, create, filter, order by, annotation, or aggregate.
type expression interface {
	asExpr() string
}

type sortExpression struct {
	field expression
	desc  bool
}

type star string

func (s star) asExpr() string {
	return "*"
}

func (l lookup) asExpr() string {
	sql := l.lhs.asExpr() + " " + l.lookupName + " " + l.rhs

	if l.not {
		sql = "NOT(" + sql + ")"
	}
	return sql
}

func (e sortExpression) asExpr() string {
	if e.desc {
		return e.field.asExpr() + " DESC"
	} else {
		return e.field.asExpr() + " ASC"
	}
}

func (a aggregate) asExpr() string {
	return a.function + "(" + a.expression.asExpr() + ") AS " + a.alias
}

func (f *AutoField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigAutoField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigIntegerField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BooleanField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *FloatField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *IntegerField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *SmallIntegerField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *TextField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *ForeignKey) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *OneToOneField) asExpr() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

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
