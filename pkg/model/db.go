/*
Package model defines the structure.
model need to set the db variable *gorm.DB
*/
package model

import (
	"github.com/jinzhu/gorm"
	// load driver

	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	log.SetFlags(log.Lshortfile)
	filename := "test.sqlite"
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	SetDB(db)
}

// SetDB return the db
func SetDB(theDB *gorm.DB) {
	if err := theDB.AutoMigrate(&Book{}, &Item{}, &Account{}, &Record{}, &Category{}, &Publisher{}, &Author{}, &Tag{}, &ClassNum{}).Error; err != nil {
		panic(err)
	}
	db = theDB
	db.DB().SetMaxOpenConns(1)
	InitCategoriesFromConfigs()
}
