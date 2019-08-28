/*
Package utils can do the method to library
utils need to set the db variable *gorm.DB
*/
package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
)

var db *gorm.DB

// SetDB return the db
func SetDB(theDB *gorm.DB) {
	if err := theDB.AutoMigrate(&model.Book{}, &model.Item{}, &model.User{}, &model.Record{}, &model.Category{}, &model.Publisher{}, &model.Author{}, &model.Tag{}).Error; err != nil {
		panic(err)
	}
	db = theDB
}
