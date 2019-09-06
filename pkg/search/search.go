package search

import (
	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
)

func searchByColumn(word, column string) (books []model.Book) {
	err := db.Where(gorm.ToColumnName(column)+` LIKE ?`, `%`+word+`%`).Find(&books).Error
	if err != nil {
		panic(err)
	}
	return
}
