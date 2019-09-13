package utils

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/lancatlin/library/pkg/model"
)

func TestImportFile(t *testing.T) {
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

func TestImportOne(t *testing.T) {
	data := map[string]string{
		"Barcodes":  "A101;A102",
		"BookName":  "TestBook",
		"Authors":   "author1;author2",
		"Supporter": "Donater",
		"Publisher": "TestPublish",
		"Year":      "2018",
		"ISBN":      "9198756321",
		"ClassNum":  "147.987",
		"Tags":      "book;test",
	}
	errChan := make(chan string)
	go importBook(data, errChan)
	select {
	case msg := <-errChan:
		if msg != "" {
			t.Error(msg)
		}
	}
}

func TestImportRepeated(t *testing.T) {
	data := map[string]string{
		"Barcodes":  "A103;A104",
		"BookName":  "TestRepeat",
		"Authors":   "author1;author2",
		"Supporter": "Donater",
		"Publisher": "TestPublish",
		"Year":      "2018",
		"ISBN":      "1234567890",
		"ClassNum":  "147.987",
		"Tags":      "book;test",
	}
	errChan := make(chan string)
	go importBook(data, errChan)
	select {
	case msg := <-errChan:
		if msg != "" {
			t.Errorf("result not nil: %s", msg)
		}
	}
	go importBook(data, errChan)
	select {
	case msg := <-errChan:
		if msg == "" {
			t.Errorf("result is nil: %s", msg)
		}
		t.Log(msg)
	}
}
