package model

import (
	"github.com/jinzhu/gorm"
	// load driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// SetDB return the db
func SetDB(theDB *gorm.DB) {
	theDB.AutoMigrate(&Book{}, &Item{}, &User{}, &Record{}, &Category{}, &Publisher{}, &Author{}, &Tag{})
	db = theDB
}
