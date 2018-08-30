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

	fields := make([]models.Field, 7)

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

	fields[0] = models.IntegerField("population")
	fields[1] = models.TextField("text")
	fields[2] = models.FloatField("float")
	fields[3] = models.DecimalField("salary", 19, 4)
	fields[4] = models.CharField("initials", 3)
	fields[5] = models.BooleanField("us citizen")
	fields[6] = models.AutoField("id")

	for _, v := range fields {
		fmt.Println(v.CreateString())
	}

	//table := models.Model{}.Init("dbtable")

	//table = table.SetFields(fields)

	//fmt.Println(table.CreateTable())

}
