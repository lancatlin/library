package utils

import (
	"log"
	"os"
	"testing"

	"github.com/lancatlin/library/pkg/model"
)

func TestImport(t *testing.T) {
	file, err := os.Open("./static/import_example.csv")
	if err != nil {
		t.Error(err)
	}
	db.Create(&model.Category{Name: "自然文學", Prefix: "A"})
	var categories []model.Category
	db.Find(&categories)
	log.Println(categories)
	err = ImportBooks(file)
	if err != nil {
		t.Error(err)
	}
}
