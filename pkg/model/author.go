package model

import "github.com/jinzhu/gorm"

// Author is the structure that record the author data
type Author struct {
	gorm.Model
	Name  string
	Works []Book `gorm:"many2many:book_authors"`
}

func (a Author) String() string {
	return a.Name
}
