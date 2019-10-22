package model

import "github.com/jinzhu/gorm"

// Role define the role of a user
type Role int

const (
	// RoleUnknown is anyone who hasn't logged in
	RoleUnknown Role = iota
	// RoleUser can booking books and access their personal data
	RoleUser
	// RoleMember has all permissions user has
	RoleMember
	// RoleAdmin has the permission to access the system
	RoleAdmin
)

// Account is the struct of account
type Account struct {
	gorm.Model
	Name      string
	Email     string
	Phone     string
	Role      Role
	Login     bool
	Records   []Record `gorm:"foreignkey:BorrowerID"`
	Donations []Item   `gorm:"foreignkey:SupporterID"`
}

// GetAccountBySession gets the user by session id
// In testing, always return a test admin
func GetAccountBySession(session string) Account {
	return Account{
		Name:  "test",
		Login: true,
		Role:  RoleAdmin,
	}
}
