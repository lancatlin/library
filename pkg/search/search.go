package search

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
)

type Searcher interface {
	SearchBooks(keyword string) []model.Book
	SearchAccounts(keyword string) []model.Account
}

func New(database *gorm.DB) Searcher {
	return &searchImpl{
		db:   database,
		lock: &sync.Mutex{},
	}
}

type searchImpl struct {
	db   *gorm.DB
	lock sync.Locker
}

func (s *searchImpl) SearchBooks(keyword string) (books []model.Book) {
	if err := s.searchByColumn(keyword, books, "name"); err != nil {
		panic(err)
	}
	return
}

func (s *searchImpl) SearchAccounts(keyword string) (accounts []model.Account) {
	if err := s.searchByColumn(keyword, &accounts, "name", "phone"); err != nil {
		panic(err)
	}
	return
}

func (s *searchImpl) searchByColumn(keyword string, dest interface{}, columns ...string) error {
	var queries []string = make([]string, len(columns))
	for i, c := range columns {
		queries[i] = fmt.Sprintf(`%s LIKE '%%%s%%'`, gorm.ToColumnName(c), keyword)
	}
	query := strings.Join(queries, " OR ")
	log.Println(query)
	return s.db.Where(query).Order("id asc").Find(dest).Error
}

/*
Not using anymore
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
*/
