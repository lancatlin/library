package model

import "github.com/jinzhu/gorm"

// Publisher is just like Author, a structure of publisher
type Publisher struct {
	gorm.Model
	Name        string
	Publication []Book
}

func (p Publisher) String() string {
	return p.Name
}
