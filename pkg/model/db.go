/*
Package model defines the structure.
model need to set the db variable *gorm.DB
*/
package model

import (
	"os"

	"github.com/jinzhu/gorm"
	// load driver

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
	if err := theDB.AutoMigrate(&Book{}, &Item{}, &Account{}, &Record{}, &Category{}, &Publisher{}, &Author{}, &Tag{}, &ClassNum{}).Error; err != nil {
		panic(err)
	}
	db = theDB
}
