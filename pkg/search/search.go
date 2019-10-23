package search

import (
	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
)

func searchByColumn(dest interface{}, word, column string) error {
	return db.Where(gorm.ToColumnName(column)+` LIKE ?`, `%`+word+`%`).Find(dest).Error
}

func SearchBooks(keyword string) (books []model.Book) {
	if err := searchByColumn(&books, keyword, "name"); err != nil {
		panic(err)
	}
	return
}

func merge(s1, s2, dest []model.Merger) {
	dest = make([]model.Merger, len(s1), len(s1)+len(s2))
	copy(dest, s1)
	for _, obj1 := range s2 {
		isAdd := true
		for _, obj2 := range dest {
			if obj1.Equal(obj2) {
				isAdd = false
				break
			}
		}
		if isAdd {
			dest = append(dest, obj1)
		}
	}
	return
}

func SearchAccounts(keyword string) []model.Account {
	var byName []model.Account
	if err := searchByColumn(byName, keyword, "name"); err != nil {
		panic(err)
	}
	var byPhone []model.Account
	if err := searchByColumn(byPhone, keyword, "phone"); err != nil {
		panic(err)
	}
	set := NewAccountSet()
	set.Add(byName)
	set.Add(byPhone)
	return set.List()
}
