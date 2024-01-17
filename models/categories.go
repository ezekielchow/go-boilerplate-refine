package models

const CategoriesTable = "categories"

var CategoriesSearchable = []string{"name"}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
