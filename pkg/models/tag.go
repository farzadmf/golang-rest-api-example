package models

import (
	"gorm.io/gorm"
)

const (

	// TagTableName is the Tag table name in the DB.
	TagTableName = "tag"
)

// Tag is a tag in our system.
type Tag struct {
	gorm.Model

	Value string `gorm:"unique"`
}

// TableName is the name of the Tag table in our DB.
func (*Tag) TableName() string {
	return TagTableName
}
