package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
)

var (
	ErrCategoryNotFound     = errors.New("This category prefix is not defined")
	ErrInvalidBarcodeFormat = errors.New("utils: Invalid barcode format")
)

// Category is define by the library.
type Category struct {
	gorm.Model
	Name   string `json:"name"`
	Books  []Book
	Prefix string `json:"prefix"`
	Amount int
}

func (c Category) String() string {
	return c.Name
}

func (c Category) addAmountAndSave() {
	c.Amount++
	if err := db.Save(&c).Error; err != nil {
		panic(err)
	}
}

func isValidBarcodeFormat(barcode string) bool {
	return regexp.MustCompile(`^[A-Z][a-z]*\d+$`).MatchString(barcode)
}

func (c Category) generateBarcode() (barcode string) {
	barcode = fmt.Sprintf("%s%d", c.Prefix, c.Amount+1)
	return
}

func (c Category) isRightNumber(barcode string) bool {
	number, err := strconv.Atoi(regexp.MustCompile(`\d+$`).FindString(barcode))
	if err != nil {
		panic(err)
	}
	return number == c.Amount+1
}

func findCategoryFromBarcode(barcode string) (category Category, err error) {
	prefix := regexp.MustCompile(`^[A-Z][a-z]*`).FindString(barcode)
	if err = db.Where(&Category{Prefix: prefix}).First(&category).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = fmt.Errorf("model error: in %s: category prefix %s was not found", barcode, prefix)
		}
		return
	}
	return
}

func GetCategory(barcodes []string) (category Category, err error) {
	isSet := false
	for _, barcode := range barcodes {
		if !isValidBarcodeFormat(barcode) {
			err = ErrInvalidBarcodeFormat
			return
		}
		var c Category
		c, err = findCategoryFromBarcode(barcode)
		if err != nil {
			return
		}
		if isSet && category.ID != c.ID {
			err = fmt.Errorf("%v has different categories prefix", barcodes)
			return
		}
		category = c
		isSet = true
	}
	return
}

func InitCategoriesFromConfigs() {
	file, err := os.Open("../../configs/categories.json")
	if err != nil {
		panic(err)
	}
	categories := loadCategoriesFromJSON(file)
	initCategories(categories)
}

func loadCategoriesFromJSON(file io.Reader) (categories []Category) {
	dec := json.NewDecoder(file)
	if err := dec.Decode(&categories); err != nil {
		panic(err)
	}
	log.Println(categories)
	return
}

func initCategories(categories []Category) {
	for _, category := range categories {
		if err := db.Where(category).FirstOrCreate(&category).Error; err != nil {
			panic(err)
		}
	}
}
