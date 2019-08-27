package model

import "github.com/jinzhu/gorm"

// User is the struct of account
type User struct {
	gorm.Model
	Name      string
	Email     string
	Phone     string
	Role      Role
	Login     bool
	Records   []Record `gorm:"foreignkey:BorrowerID"`
	Donations []Item   `gorm:"foreignkey:SupporterID"`
	Password  []byte
}

// GetUser gets the user by session id
func GetUser(session string) User {
	return User{
		Name:  "test",
		Login: true,
		Role:  RoleAdmin,
	}
}
