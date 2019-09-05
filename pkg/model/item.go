package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Status to record the status of a item
type Status int

func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "未設定"
	case StatusInside:
		return "館內"
	case StatusLending:
		return "借出中"
	case StatusMissing:
		return "遺失"
	}
	return ""
}

const (
	// StatusUnknown is the default status
	StatusUnknown Status = iota
	// StatusInside means the book is in the library
	StatusInside
	// StatusLending means the book is lending by someone
	StatusLending
	// StatusMissing means the book is missing
	StatusMissing
)

// Item is the instance of a book in the library
type Item struct {
	gorm.Model
	Barcode   string `gorm:"unique"`
	Book      Book
	BookID    int
	Supporter string
	Records   []Record
}

// Status return the status of a book
func (i *Item) Status() Status {
	record := i.ProcessingRecord()
	if record == nil {
		return StatusInside
	}
	if time.Now().Sub(record.LendingTime) > BorrowingPeriods {
		return StatusMissing
	}
	return StatusLending
}

func (i *Item) ProcessingRecord() *Record {
	return nil
}

func (book *Book) NewItem(supporter string) (item Item, err error) {
	return book.newItem("", supporter)
}

func (book *Book) NewItemWithBarcode(barcode, supporter string) (item Item, err error) {
	return book.newItem(barcode, supporter)
}

func (book *Book) newItem(barcode, supporter string) (item Item, err error) {
	category, err := getCategoryAndCheckBarcode(barcode)
	if book.Category.Name == "" {
		book.Category = category
	} else if category.Name != book.Category.Name {
		err = fmt.Errorf(`model error: %s has wrong prefix. want %s have %s`, barcode, book.Category.Prefix, category.Prefix)
		return
	}
	item = Item{
		Barcode:   barcode,
		Book:      *book,
		Supporter: supporter,
	}
	book.Category.addAmountAndSave()
	return
}
