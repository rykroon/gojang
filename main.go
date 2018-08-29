package main

import (
  "fmt"
  "./models"
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

  var fields [7]models.Field

  fields[0] = models.IntegerField{}.Init("population")
  fields[1] = models.TextField{}.Init("text")
  fields[2] = models.FloatField{}.Init("float")
  fields[3] = models.DecimalField{}.Init("salary",19,4).Null(true)
  fields[4] = models.CharField{}.Init("initials",3).Unique(true)
  fields[5] = models.BooleanField{}.Init("us citizen")
  fields[6] = models.AutoField{}.Init("id").PrimaryKey(true)


  for _,v := range fields {
    fmt.Println(v.CreateString())
  }

  table := models.Model{}

}
