package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"assignment/pkg/models"
)

// This file contains functions for member operations.

// CreateRole creates a role in the DB.
func (r *repository) CreateRole(name string) (*models.Role, error) {
	role := models.Role{
		Name: name,
	}

	if err := r.db.Create(&role).Error; err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	return &role, nil
}

// GetRoles returns the role matching the name.
func (r *repository) GetRole(name string) (*models.Role, error) {
	var role models.Role

	err := r.db.First(&role, "name = ?", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get role: %w", err)
	}

	return &role, nil
}

// GetRoles returns all the roles.
func (r *repository) GetRoles() ([]*models.Role, error) {
	var roles []*models.Role

	if err := r.db.Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("get roles: %w", err)
	}

	return roles, nil
}

// UpdateRole updates a role in the DB.
func (r *repository) UpdateRole(roleID uint, name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, roleID).Error; err != nil {
		return nil, fmt.Errorf("update role: %w", err)
	}

	role.Name = name

	if err := r.db.Save(&role).Error; err != nil {
		return nil, fmt.Errorf("update role: %w", err)
	}

	return &role, nil
}

// DeleteRole deletes a role in the DB.
func (r *repository) DeleteRole(roleID uint) error {
	if err := r.db.Delete(&models.Role{}, roleID).Error; err != nil {
		return fmt.Errorf("delete role: %w", err)
	}

	return nil
}
