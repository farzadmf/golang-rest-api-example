package models

import (
	"gorm.io/gorm"
)

const (
	// RoleTableName is the Role table name in the DB.
	RoleTableName = "role"
)

// Role is a tag in our system.
type Role struct {
	gorm.Model

	// Name is the name of the role.
	Name string `gorm:"unique"`
}

// TableName is the name of the MemberRole table in our DB.
func (*Role) TableName() string {
	return RoleTableName
}
