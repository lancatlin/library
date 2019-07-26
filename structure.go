package main

import (
	"strings"
	"time"

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

var (
	// BorrowingPeriods 為圖書館的借閱期限，預設為 30 天
	BorrowingPeriods time.Duration = time.Hour * 24 * 30
)

type Role int

const (
	RoleAdmin Role = iota
	RoleMember
	RoleUser
)

// Book is the structure of the catalog of books
// Catalog "has many" items
// 	Many to Many Author
//	Has one Publisher
//	Has one Classification
//	Has many Tags
type Book struct {
	gorm.Model
	BookName string
	// Many to many authors
	Authors []Author `gorm:"many2many:book_authors"`
	// belongs to one publisher
	Publisher   Publisher
	PublisherID int
	Year        int
	// belongs to one category
	Category             Category
	CategoryID           int
	ClassificationNumber string
	// has many items
	Items []Item
	// many to many tags
	Tags        []Tag `gorm:"many2many:book_tags"`
	Cover       string
	Description string
}

func (b Book) Author() string {
	// 將所有作者以頓號分隔排列
	l := make([]string, len(b.Authors))
	for i, v := range b.Authors {
		l[i] = v.String()
	}
	return strings.Join(l, "、")
}

// Author is the structure that record the author data
type Author struct {
	gorm.Model
	Name  string
	Works []Book `gorm:"many2many:book_authors"`
}

func (a Author) String() string {
	return a.Name
}

// Publisher is just like Author, a structure of publisher
type Publisher struct {
	gorm.Model
	Name        string
	Publication []Book
}

func (p Publisher) String() string {
	return p.Name
}

// Classification is define by the library
type Category struct {
	gorm.Model
	Name  string
	Books []Book
}

func (c Category) String() string {
	return c.Name
}

// Item is the instance of a book in the library
type Item struct {
	Barcode      string `gorm:"primary_key"`
	Book         Book
	BookID       int
	NewBookLabel string
	Supporter    User `gorm:"foreignkey:SupporterID"`
	SupporterID  int
	Records      []Record
}

func (i *Item) Status() Status {
	record := i.Record()
	if record == nil {
		return StatusInside
	}
	if time.Now().Sub(record.LendingTime) > BorrowingPeriods {
		return StatusMissing
	}
	return StatusLending
}

func (i *Item) Record() *Record {
	// 返回還沒有完成的 Record，如果無返回 nil
	return nil
}

// User the structure of users
type User struct {
	gorm.Model
	UserName  string
	Email     string
	Phone     string
	Role      Role
	Login     bool
	Records   []Record `gorm:"foreignkey:BorrowerID"`
	Donations []Item   `gorm:"foreignkey:SupporterID"`
	Password  []byte
}

// Tag records 關鍵字
// Many2Many with catalogs
type Tag struct {
	gorm.Model
	Name  string `gorm:"many2many:catalog_tags"`
	Books []Book
}

func (t Tag) String() string {
	return t.Name
}

type Record struct {
	gorm.Model
	Borrower    User `gorm:"foreignkey:BorrowerID"`
	BorrowerID  int
	Item        Item
	ItemID      int
	LendingTime time.Time
	ReturnTime  time.Time
}
