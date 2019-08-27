package model

import (
	"log"
	"os"
	"testing"
)

func TestImport(t *testing.T) {
	file, err := os.Open("./static/import_example.csv")
	if err != nil {
		t.Error(err)
	}
	db.Create(&Category{Name: "自然文學", Prefix: "A"})
	var categories []Category
	db.Find(&categories)
	log.Println(categories)
	err = booksImport(file)
	if err != nil {
		t.Error(err)
	}
}
