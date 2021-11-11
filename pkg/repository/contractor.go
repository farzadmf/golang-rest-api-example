package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"assignment/pkg/models"
)

// This file contains functions for contractors.

// CreateContractor creates a contractor in the DB.
func (r *repository) CreateContractor(name string, duration int) (*models.Member, error) {
	contractor := models.Member{
		Name:     name,
		Type:     models.Contractor,
		Duration: duration,
	}

	if err := r.db.Create(&contractor).Error; err != nil {
		return nil, fmt.Errorf("create contractor: %w", err)
	}

	return &contractor, nil
}

// GetContractor returns the contractor (member) matching the ID.
func (r *repository) GetContractor(contractorID uint) (*models.Member, error) {
	var contractor *models.Member
	err := r.db.
		Preload("Tags").
		Where("id = ? AND type = ?", contractorID, models.Contractor).
		First(&contractor).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get contractor: %w", err)
	}

	return contractor, nil
}

// GetContractors returns all the contractors.
func (r *repository) GetContractors() ([]*models.Member, error) {
	var contractors []*models.Member
	if err := r.db.
		Preload("Tags").
		Find(&contractors, "type = ?", models.Contractor).
		Error; err != nil {
		return nil, fmt.Errorf("get contractors: %w", err)
	}

	return contractors, nil
}
