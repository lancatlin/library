/*
Package model defines the structure.
model need to set the db variable *gorm.DB
*/
package model

import (
	"io"
	"os"

	"github.com/jinzhu/gorm"
	// load driver
	"encoding/json"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	filename := "test.sqlite"
	if err := os.Remove(filename); err != nil {
		panic(err)
	}
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	SetDB(db)
	InitCategoriesFromConfigs()
}

// SetDB return the db
func SetDB(theDB *gorm.DB) {
	if err := theDB.AutoMigrate(&Book{}, &Item{}, &Account{}, &Record{}, &Category{}, &Publisher{}, &Author{}, &Tag{}).Error; err != nil {
		panic(err)
	}
	db = theDB
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
	return
}

func initCategories(categories []Category) {
	for _, category := range categories {
		if err := db.FirstOrCreate(&category).Error; err != nil {
			panic(err)
		}
	}
}
