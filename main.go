package main

import (
	"./models"
	"fmt"
)

func main() {
	// id := models.IntegerField{}.Init("id").PrimaryKey(true)
	// text := models.TextField{}.Init("text")
	// ratio := models.FloatField{}.Init("float")
	// salary := models.DecimalField{}.Init("salary",19,4)
	// initials := models.CharField{}.Init("initials",3)

	// fmt.Println(id.CreateString())
	// fmt.Println(text.CreateString())
	// fmt.Println(ratio.CreateString())
	// fmt.Println(salary.CreateString())
	// fmt.Println(initials.CreateString())

	//fields := make([]models.Field, 7)

	// fields[0] = models.IntegerField{}.Init("population")
	// fields[1] = models.TextField{}.Init("text")
	// fields[2] = models.FloatField{}.Init("float")
	// fields[3] = models.DecimalField{}.Init("salary",19,4).Null(true)
	// fields[4] = models.CharField{}.Init("initials",3).Unique(true)
	// fields[5] = models.BooleanField{}.Init("us citizen")
	// fields[6] = models.AutoField{}.Init("id").PrimaryKey(true)

	// fields = append(fields, models.IntegerField{}.Init("population"))
	// fields[1] = models.TextField{}.Init("text")
	// fields[2] = models.FloatField{}.Init("float")
	// fields[3] = models.DecimalField{}.Init("salary",19,4).Null(true)
	// fields[4] = models.CharField{}.Init("initials",3).Unique(true)
	// fields[5] = models.BooleanField{}.Init("us citizen")
	// fields[6] = models.AutoField{}.Init("id").PrimaryKey(true)

  table := models.Model("myTable")

  //table.Field["population"].AddField(models.IntegerField())

	table.Fields[0] = models.IntegerField("population")
	table.Fields[1] = models.TextField("text")
	table.Fields[2] = models.FloatField("float")
	table.Fields[3] = models.DecimalField("salary", 19, 4)
	table.Fields[4] = models.CharField("initials", 3)
	table.Fields[5] = models.BooleanField("us citizen")
	table.Fields[6] = models.AutoField("id")

	//for _, v := range table.Fields {
//		fmt.Println(v.CreateString())
	//}
  fmt.Println(table.CreateTable())


	//table := models.Model{}.Init("dbtable")

	//table = table.SetFields(fields)

	//fmt.Println(table.CreateTable())

}
