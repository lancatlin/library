package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Category is define by the library.
type Category struct {
	gorm.Model
	Name   string
	Books  []Book
	Prefix string
	Amount int
}

func (c Category) String() string {
	return c.Name
}

func (c Category) append(item *Item) {
	c.Amount++
	item.Barcode = fmt.Sprintf("%s%d", c.Prefix, c.Amount)
	if err := db.Save(&c).Error; err != nil {
		panic(err)
	}
}
