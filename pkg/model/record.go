package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Record record the data of a lending event
type Record struct {
	gorm.Model
	Borrower    Account `gorm:"foreignkey:BorrowerID"`
	BorrowerID  int
	Item        Item
	ItemID      int
	LendingTime time.Time
	ReturnTime  time.Time
}
