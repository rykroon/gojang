package main

import (
	"./models"
	"fmt"
)

type MyModel struct {
	models.Model

	Bool    models.Field
	Char    models.Field
	Decimal models.Field
	Float   models.Field
	Number  models.Field

	Text models.Field
}

func main() {

	MyModel := MyModel{}
	MyModel.Init("mymodel")

	//table.Field["population"].AddField(models.IntegerField())

	MyModel.Bool = models.BooleanField().IsNull()
	MyModel.Char = models.CharField(30).Unique()
	MyModel.Decimal = models.DecimalField(9, 2)
	MyModel.Float = models.FloatField()
	MyModel.Number = models.IntegerField()
	MyModel.Text = models.TextField()

	models.Migrate(MyModel)

  qs := MyModel.Objects.All()
	fmt.Println(qs.Query)

  qs = qs.Filter("char__exact","Hello").Filter("float__gte",5000).Exclude("text__startswith","Meow")
  fmt.Println(qs.Query)

	fmt.Println("")

}
