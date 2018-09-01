package main

import (
	"./models"
	"fmt"
)

type Book struct {
	models.Model

	Isbn   models.Field
	Title  models.Field
	Author models.Field
	Price  models.Field
  Stuff int
}

func main() {

	Book := Book{}
  Book.Init("book")
  qs := Book.Objects.All()
  fmt.Println(qs.Query)


	//table.Field["population"].AddField(models.IntegerField())

	Book.Isbn = models.IntegerField()
	Book.Title = models.TextField()
	Book.Author = models.CharField(30)
	Book.Price = models.DecimalField(9, 2)
	//table.Fields[1] = models.TextField("text")
	//table.Fields[2] = models.FloatField("float")
	//table.Fields[3] = models.DecimalField("salary", 19, 4)
	//table.Fields[4] = models.CharField("initials", 3)
	//table.Fields[5] = models.BooleanField("us citizen")
	//table.Fields[6] = models.AutoField("id")

	models.Migrate(Book)
  fmt.Println("")
	//for _, v := range table.Fields {
	//		fmt.Println(v.CreateString())
	//}
	//fmt.Println(table.CreateTable())

	//table := models.Model{}.Init("dbtable")

	//table = table.SetFields(fields)

	//fmt.Println(table.CreateTable())

}
