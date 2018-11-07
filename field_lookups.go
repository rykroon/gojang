package gojang

import ()

type lookup struct {
	not        bool
	lhs        field
	lookupName string
	rhs        string
}

//type lookup string

func exactIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "=", rhs: intAsSql(value)}
}

func inIntField(field IntField, values []int) lookup {
	return lookup{lhs: field, lookupName: "IN", rhs: integersAsSql(values)}
}

func gtIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: ">", rhs: intAsSql(value)}
}

func gteIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: ">=", rhs: intAsSql(value)}
}

func ltIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "<", rhs: intAsSql(value)}
}

func lteIntField(field IntField, value int) lookup {
	return lookup{lhs: field, lookupName: "<=", rhs: intAsSql(value)}
}

func rangeIntField(field IntField, from, to int) lookup {
	lookup := lookup{lhs: field, lookupName: "BETWEEN"}
	lookup.rhs = intAsSql(from) + " AND " + intAsSql(to)
	return lookup
}

func fieldIsNull(field field, value bool) lookup {
	lookup := lookup{lhs: field, lookupName: "IS"}

	if value {
		lookup.rhs = "NULL"
	} else {
		lookup.rhs = "NOT NULL"
	}

	return lookup
}
