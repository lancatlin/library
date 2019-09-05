package utils

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/lancatlin/library/pkg/model"
)

func TestImport(t *testing.T) {
	file, err := os.Open("./testData/success_example.csv")
	if err != nil {
		t.Error(err)
	}
	var categories []model.Category
	db.Find(&categories)
	log.Println(categories)
	errList := ImportBooks(file)
	for _, v := range errList {
		t.Error(v)
	}
}

func TestReadRaws(t *testing.T) {
	file := strings.NewReader(
		`"Barcodes","BookName","Authors","Supporter","Publisher","Year","ISBN","ClassNum","Tags"
"barcodes","bookname","authors","supporter","publisher","year","isbn","classnum","tags"`)
	data := readRaws(file)
	if len(data) != 1 {
		t.Errorf("wrong data length: %d \n %v", len(data), data)
	}
	for _, raw := range data {
		for key, value := range raw {
			if value != strings.ToLower(key) {
				t.Errorf("Answer not equal: key %s value %s \n %v", key, value, raw)
			}
		}
	}
}

func TestParse(t *testing.T) {
	data := map[string]string{
		"Barcodes":  "A001;A002",
		"BookName":  "TestBook",
		"Authors":   "author1;author2",
		"Supporter": "Donater",
		"Publisher": "TestPublish",
		"Year":      "2018",
		"ISBN":      "9198756321",
		"ClassNum":  "147.987",
		"Tags":      "book;test",
	}
	err := parse(data)
	if err != nil {
		t.Error(err)
	}
}
