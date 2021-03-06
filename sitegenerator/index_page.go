package sitegenerator

import (
	"fmt"
	"io"
	"log"
)

type IndexPage struct {
	PageName       string
	ResourceModels []ResourceModel
	Path           []string
	CategoryList   []Category
}

type Category struct {
	CategoryName  string
	CategoryCount int
}

var IndexPagePath = []string{"All Resources"}

func NewIndexPage(resourceModels []ResourceModel) IndexPage {
	return IndexPage{
		"index-page",
		resourceModels,
		IndexPagePath,
		createCategoryList(resourceModels),
	}
}

func (i *IndexPage) Generate(w io.Writer) error {
	err := load("index.html").ExecuteTemplate(w, "index.html", i)
	if err != nil {
		return fmt.Errorf("cannot write index.html: %s", err)
	}
	log.Print("Index page generated")
	return nil
}

func createCategoryList(resources []ResourceModel) []Category {
	var (
		categoryList        []string
		uniqueCategoryList  []Category
		uniqueCategoryNames []string
	)

	for _, resource := range resources {
		categoryList = append(categoryList, resource.Categories...)
	}

	foundCategories := make(map[string]int)

	for _, category := range categoryList {
		_, found := foundCategories[category]

		if found {
			foundCategories[category] = foundCategories[category] + 1
		} else {
			foundCategories[category] = 1
			uniqueCategoryNames = append(uniqueCategoryNames, category)
		}
	}

	for _, categoryName := range uniqueCategoryNames {
		uniqueCategoryList = append(uniqueCategoryList, Category{categoryName, foundCategories[categoryName]})
	}

	return uniqueCategoryList
}
