package models

import (
	"errors"

	"gorm.io/gorm"
)

const (
	// MemberTableName is the Member table name in the DB.
	MemberTableName = "member"
)

var (
	// ErrDurationRequired is raised when the duration for a contractor is not specified.
	ErrDurationRequired = errors.New("duration required for contractors")

	// ErrRoleRequired is raised when the role for an employee is not specified.
	ErrRoleRequired = errors.New("role required for employees")
)

// Member represents a member of our team.
type Member struct {
	gorm.Model

	// Name is the name of the employee.
	Name string

	// Duration is the contract duration (for contractors).
	Duration int

	// Type is the member type (contractor or employee).
	Type int

	// Role is the role ID of the member (for employees).
	RoleID *uint

	// Role is the role of the member (for employees).
	Role *Role

	// Tags is the tags for an employee
	Tags []Tag `gorm:"many2many:member_tags"`
}

// TableName is the name of the Member table in our DB.
func (*Member) TableName() string {
	return MemberTableName
}

// IsValid checks whether this member type is valid.
func (m *Member) IsValid() (bool, error) {
	if m.Type == Contractor && m.Duration == 0 {
		return false, ErrDurationRequired
	}

	if m.Type == Employee && m.Role == nil && m.RoleID == nil {
		return false, ErrRoleRequired
	}

	return true, nil
}
