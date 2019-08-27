package model

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
