package gojang

import (
	"errors"
)

type ModelInstance interface {
	Save() error
	Delete() error
	RefreshFromDb() error
}

func (m *Model) ToObj() object {
	obj := newObj()

	for _, field := range m.fields {
		attrName := m.colToAttr[field.ColumnName()]
		obj.SetAttr(attrName, field.getValue())
	}

	return obj
}

func (m *Model) FromObj(obj object) error {
	for _, field := range m.fields {
		if obj.HasAttr(field.Alias()) {
			//some sort of switch statement
		}
	}

	return nil
}

//If instance does not have a primary key then it will insert into the database
//Otherwise it updates the record
func (m *Model) Save() error {
	if m.Pk.Id() == 0 {
		id, err := m.insert()
		if err != nil {
			return err
		}

		m.Pk.setInt(id)

	} else {
		err := m.update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) insert() (int, error) {
	var createList []assignment

	for _, field := range m.fields {
		if !field.PrimaryKey() {
			createList = append(createList, field.asAssignment())
		}
	}

	obj, err := m.Objects.Create(createList...)
	if err != nil {
		return 0, err
	}

	pkeyAttr := m.colToAttr[m.Pk.ColumnName()]
	if obj.HasAttr(pkeyAttr) {
		return obj.GetAttr(pkeyAttr).(int), nil
	}

	return 0, errors.New("gojang: idk, we lost the key, sorry")
}

//
func (m *Model) update() error {
	var updateList []assignment

	for _, field := range m.fields {
		if !field.PrimaryKey() {
			updateList = append(updateList, field.asAssignment())
		}
	}

	qs := m.Objects.Filter(m.Pk.Exact(m.Pk.Id()))
	_, err := qs.Update(updateList...)
	return err
}

func (m *Model) Delete() error {
	qs := m.Objects.Filter(m.Pk.Exact(m.Pk.Id()))
	_, err := qs.Delete()
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) RefreshFromDb() error {
	obj, err := m.Objects.Get(m.Pk.Exact(m.Pk.Id()))
	if err != nil {
		return err
	}

	m.FromObj(obj)
	return nil
}
