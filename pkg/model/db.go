/*
Package model defines the structure.
model need to set the db variable *gorm.DB
*/
package model

import (
	"github.com/jinzhu/gorm"
	// load driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// SetDB return the db
func SetDB(theDB *gorm.DB) {
	if err := theDB.AutoMigrate(&Book{}, &Item{}, &User{}, &Record{}, &Category{}, &Publisher{}, &Author{}, &Tag{}).Error; err != nil {
		panic(err)
	}
	db = theDB
}
