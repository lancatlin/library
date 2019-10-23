package model

// Role define the role of a user
type Role int

const (
	// RoleUnknown is anyone who hasn't logged in
	RoleUnknown Role = iota
	// RoleAdmin has the permission to access the system
	RoleAdmin
	// RoleMember has all permissions user has
	RoleMember
	// RoleUser can booking books and access their personal data
	RoleUser
)

// Account is the struct of account
type Account struct {
	ID        uint
	Name      string
	Phone     string
	Role      Role
	Login     bool     `gorm:"-"`
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

func (a Account) Equal(obj interface{}) bool {
	if acct, ok := obj.(Account); ok {
		return a.ID == acct.ID
	}
	return false
}
