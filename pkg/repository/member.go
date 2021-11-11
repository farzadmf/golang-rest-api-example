package repository

import (
	"assignment/pkg/models"
	"fmt"
)

// This file contains generic functions for members.

// GetMember returns the member matching the ID.
func (r *repository) GetMember(memberID uint) (*models.Member, error) {
	var member models.Member
	if err := r.db.First(&member, memberID).Error; err != nil {
		return nil, fmt.Errorf("get member: %w", err)
	}

	return &member, nil
}

// DeleteMember deletes a member in the DB.
func (r *repository) DeleteMember(memberID uint) error {
	if err := r.db.Delete(&models.Member{}, memberID).Error; err != nil {
		return fmt.Errorf("delete member: %w", err)
	}

	return nil
}
