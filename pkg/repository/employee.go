package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"assignment/pkg/models"
)

// This file contains functions for employees.

// CreateEmployee creates an employee in the DB.
func (r *repository) CreateEmployee(name string, roleID uint) (*models.Member, error) {
	employee := models.Member{
		Name:   name,
		RoleID: &roleID,
		Type:   models.Employee,
	}

	if err := r.db.Create(&employee).Error; err != nil {
		return nil, fmt.Errorf("create employee: %w", err)
	}

	var role models.Role
	if err := r.db.First(&role, roleID).Error; err != nil {
		return nil, fmt.Errorf("create employee get role: %w", err)
	}

	employee.Role = &role

	return &employee, nil
}

// GetEmployee returns the employee (member) matching the ID.
func (r *repository) GetEmployee(employeeID uint) (*models.Member, error) {
	var employee *models.Member
	err := r.db.
		Preload("Role").
		Preload("Tags").
		Where("id = ? AND type = ?", employeeID, models.Employee).
		First(&employee).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get employee: %w", err)
	}

	return employee, nil
}

// GetEmployees returns all the employees.
func (r *repository) GetEmployees() ([]*models.Member, error) {
	var employees []*models.Member
	if err := r.db.
		Preload("Role").
		Preload("Tags").
		Find(&employees, "type = ?", models.Employee).
		Error; err != nil {
		return nil, fmt.Errorf("get employees: %w", err)
	}

	return employees, nil
}
