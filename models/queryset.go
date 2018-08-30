package models

import ()

type Queryset struct {
  query string

  //select_ string
  //from string
  //where []string
  //groupBy string
  //having string
  //orderBy string
}

func (q Queryset) Filter() Queryset {

}

func (q Queryset) Exclude() Queryset {

}

func (q Queryset) Delete() Queryset {

}
