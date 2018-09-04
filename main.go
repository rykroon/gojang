package main

import (
	"./models"
	"fmt"
)

type Schema struct {
	Table1 models.Model
	Table2 models.Model
	Table3 models.Model
}

type OtherModel struct {
	models.Model
	Id models.Field
}

type MyModel struct {
	models.Model

	Bool    models.Field
	Char    models.Field
	Decimal models.Field
	Float   models.Field
	Number  models.Field
	Fkey	  models.Field


	Text models.Field
}

func main() {

	MyModel := MyModel{}
	MyModel.Init("mymodel")

	OtherModel := OtherModel{}
	OtherModel.Init("other_model")


	MyModel.Bool = models.BooleanField().IsNull()
	MyModel.Char = models.CharField(30).Unique()
	MyModel.Decimal = models.DecimalField(9, 2)
	MyModel.Float = models.FloatField()
	MyModel.Number = models.IntegerField()
	MyModel.Text = models.TextField()
	MyModel.Fkey = models.ForeignKey(&OtherModel.Model, models.Cascade)



	//Other Model

	OtherModel.Id = models.AutoField().PrimaryKey()


	models.Migrate(MyModel)
	models.Migrate(OtherModel)

	qs := MyModel.Objects.All()
	fmt.Println(qs.Query)

	qs = qs.Filter("char__exact", "Hello").Filter("float__gte", 5000).Exclude("text__startswith", "Meow")
	qs = qs.Filter("bool", true).Filter("bool__isnull", true).Filter("number__in", []int{1, 2, 3})

	qs2 := MyModel.Objects.All()
	qs = qs.Filter("id__in", qs2)
	qs = qs.OrderBy("id", "-text")
	fmt.Println(qs.Query)

	fmt.Println("")

}
