package utils

import (
	"log"
	"os"
	"testing"

	"github.com/lancatlin/library/pkg/model"
)

func TestImport(t *testing.T) {
	file, err := os.Open("./testData/success_example.csv")
	if err != nil {
		t.Error(err)
	}
	model.SetDB(db)
	db.Create(&model.Category{Name: "自然文學", Prefix: "A"})
	var categories []model.Category
	db.Find(&categories)
	log.Println(categories)
	errList := ImportBooks(file)
	for _, v := range errList {
		t.Error(v)
	}
}
