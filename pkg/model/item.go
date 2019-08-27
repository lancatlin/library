package model

import "time"

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

// Status return the status of a book
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

// Record return the undone record of an item.
// if not exist, return nil.
func (i *Item) Record() *Record {
	return nil
}
