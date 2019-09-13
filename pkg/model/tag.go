package model

import "github.com/jinzhu/gorm"

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
