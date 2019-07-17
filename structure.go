package main

import (
	"github.com/jinzhu/gorm"
)

// Status to record the status of a item
type Status int

func (s Status) String() string {
	switch s {
	case 0:
		return "館內"
	case 1:
		return "借出中"
	case 2:
		return "遺失"
	default:
		return ""
	}
}

const (
	// StatusInside means the book is in the library
	StatusInside Status = iota
	// StatusLending means the book is lending by someone
	StatusLending
	// StatusMissing means the book is missing
	StatusMissing
)

// Catalog is the structure of the catalog of books
// Catalog "has many" items
// 	Many to Many Author
//	Has one Publisher
//	Has one Classification
//	Has many Tags
type Catalog struct {
	gorm.Model
	Name                 string
	Authors              []Author
	Publisher            Publisher
	Year                 int
	Classification       Classification
	ClassificationNumber string
	Items                []Item
	Tags                 []Tag
}

// Author is the structure that record the author data
type Author struct {
	gorm.Model
	Name  string
	Works []Catalog
}

// Publisher is just like Author, a structure of publisher
type Publisher struct {
	gorm.Model
	Name        string
	Publication []Catalog
}

// Classification is define by the library
type Classification struct {
	gorm.Model
	Name  string
	Books []Catalog
}

// Item is the instance of a book in the library
type Item struct {
	gorm.Model
	Barcode      string `gorm:"primary_key"`
	Catalog      Catalog
	Status       Status
	NewBookLabel string
	Borrower     User
	SupportBy    User
}

// User the structure of users
type User struct {
	gorm.Model
	Name      string
	Email     string
	Phone     string
	Lendings  []Item
	Donations []Item
	Password  []byte
}

// Tag records 關鍵字
// Many2Many with catalogs
type Tag struct {
	gorm.Model
	Name  string `gorm:"many2many:catalog_tags"`
	Books []Catalog
}
