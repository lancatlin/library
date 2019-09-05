package model

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// Book is the structure of the catalog of books
// Catalog "has many" items
// 	Many to Many Author
//	Has one Publisher
//	Has one Classification
//	Has many Tags
type Book struct {
	gorm.Model
	Name string
	// Many to many authors
	Authors []Author `gorm:"many2many:book_authors"`
	// belongs to one publisher
	Publisher   Publisher
	PublisherID int
	ISBN        int
	Year        int
	// belongs to one category
	Category   Category
	CategoryID int
	ClassNums  []ClassNum `gorm:"many2many:book_class_nums"`
	// has many items
	Items []Item
	// many to many tags
	Tags        []Tag `gorm:"many2many:book_tags"`
	Cover       string
	Description string
}

// Author return the all the authors and join them into a string
func (b Book) Author() string {
	// 將所有作者以頓號分隔排列
	l := make([]string, len(b.Authors))
	for i, v := range b.Authors {
		l[i] = v.String()
	}
	return strings.Join(l, "、")
}

func (c Category) checkAndGenerateBarcode() string {
	return c.generateBarcode()
}

func (b *Book) Create() {
	if err := db.Create(b).Error; err != nil {
		panic(err)
	}
}
