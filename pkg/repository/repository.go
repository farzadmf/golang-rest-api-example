// Package repository contains types to talk to DB.
package repository

import (
	"assignment/pkg/config"
	"assignment/pkg/models"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// ErrTagNotFound is we try to tag a member, but the tag doesn't exist in the DB.
	ErrTagNotFound = errors.New("tag not found")
)

// Repository is the interface to talk to the DB.
type Repository interface {
	// Reset is a convenient function to reset the DB.
	Reset() error

	// ++++++++++++++++++++++++++ ROLES +++++++++++++++++++++++++++

	// CreateRole creates a role in the DB.
	CreateRole(name string) (*models.Role, error)

	// GetRole returns the role matching the name.
	GetRole(name string) (*models.Role, error)

	// GetRoles returns all the roles.
	GetRoles() ([]*models.Role, error)

	// UpdateRole updates a role in the DB.
	UpdateRole(roleID uint, name string) (*models.Role, error)

	// DeleteRole deletes a role in the DB.
	DeleteRole(roleID uint) error

	// +++++++++++++++++++++++++++ TAGS +++++++++++++++++++++++++++

	// CreateTag creates a tag in the DB.
	CreateTag(value string) (*models.Tag, error)

	// GetTag returns the tag matching the value.
	GetTag(value string) (*models.Tag, error)

	// GetTags returns all the tags.
	GetTags() ([]*models.Tag, error)

	// UpdateTag updates a tag in the DB.
	UpdateTag(tagID uint, value string) (*models.Tag, error)

	// DeleteTag deletes a tag in the DB.
	DeleteTag(tagID uint) error

	// +++++++++++++++++++++++++ MEMBERS ++++++++++++++++++++++++++

	// GetMember returns the member matching the ID.
	GetMember(memberID uint) (*models.Member, error)

	// DeleteMember deletes a member in the DB.
	DeleteMember(memberID uint) error

	// TagMember adds tags for a member.
	TagMember(memberID uint, tags []string) error

	// +++++++++++++++++++++++ CONTRACTORS ++++++++++++++++++++++++

	// CreateContractor creates a contractor in the DB.
	CreateContractor(name string, duration int) (*models.Member, error)

	// GetContractor returns the contractor (member) matching the ID.
	GetContractor(contractorID uint) (*models.Member, error)

	// GetContractors returns all the contractors.
	GetContractors() ([]*models.Member, error)

	// ++++++++++++++++++++++++ EMPLOYEES +++++++++++++++++++++++++

	// CreateEmployee creates an employee in the DB.
	CreateEmployee(name string, roleID uint) (*models.Member, error)

	// GetEmployee returns the employee (member) matching the ID.
	GetEmployee(employeeID uint) (*models.Member, error)

	// GetEmployees returns all the employees.
	GetEmployees() ([]*models.Member, error)
}

type repository struct {
	db *gorm.DB
}

// New creates a new repository.
func New(cfg *config.Config) (Repository, error) {
	r := repository{}

	db, err := connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("repository.New(): %w", err)
	}

	r.db = db

	if err := r.db.AutoMigrate(&models.Tag{}); err != nil {
		return nil, fmt.Errorf("AutoMigrate(Tag): %w", err)
	}
	if err := r.db.AutoMigrate(&models.Member{}); err != nil {
		return nil, fmt.Errorf("AutoMigrate(Member): %w", err)
	}
	if err := r.db.AutoMigrate(&models.Role{}); err != nil {
		return nil, fmt.Errorf("AutoMigrate(MemberRole): %w", err)
	}

	return &r, nil
}

func connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, fmt.Errorf("connect(): %w", err)
	}

	return db, nil
}

// Reset is a convenient function to reset the DB.
func (r *repository) Reset() error {
	for _, table := range []string{
		models.MemberTableName,
		models.TagTableName,
		models.RoleTableName,
	} {
		if err := r.db.Exec(
			fmt.Sprintf(
				"TRUNCATE TABLE %s CASCADE; ALTER SEQUENCE %s_id_seq RESTART;",
				table, table,
			),
		).Error; err != nil {
			return err
		}
	}

	return nil
}
