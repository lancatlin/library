package model

import (
	"errors"
	"regexp"

	"github.com/jinzhu/gorm"
)

var (
	// ErrInvalidClassNum is threw when regexp not match
	ErrInvalidClassNum = errors.New("utils: invalid classification number")
)

// ClassNum records the classification number
type ClassNum struct {
	gorm.Model
	Value string
	Books []Book `gorm:"many2many:book_class_nums"`
}

func (c ClassNum) String() string {
	return c.Value
}

// NewClassNum return a ClassNum and an error
func NewClassNum(s string) (c ClassNum, err error) {
	if !regexp.MustCompile(`^\d+\.?\d*$`).MatchString(s) {
		err = ErrInvalidClassNum
		return
	}
	c = ClassNum{
		Value: s,
	}
	return
}
