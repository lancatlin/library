package model

import "github.com/jinzhu/gorm"

// Category is define by the library.
type Category struct {
	gorm.Model
	Name   string
	Books  []Book
	Prefix string
}

func (c Category) String() string {
	return c.Name
}
