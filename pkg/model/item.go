package model

import (
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
