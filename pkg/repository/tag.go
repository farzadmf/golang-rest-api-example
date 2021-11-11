package repository

import (
	"assignment/pkg/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// This file contains functions for tags.

// CreateTag creates a tag in the DB.
func (r *repository) CreateTag(value string) (*models.Tag, error) {
	tag := models.Tag{
		Value: value,
	}

	if err := r.db.Create(&tag).Error; err != nil {
		return nil, fmt.Errorf("create tag: %w", err)
	}

	return &tag, nil
}

// GetTag returns the tag matching the value.
func (r *repository) GetTag(value string) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.First(&tag, "value = ?", value).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get tag: %w", err)
	}

	return &tag, nil
}

// GetTags returns all the tags.
func (r *repository) GetTags() ([]*models.Tag, error) {
	var tags []*models.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, fmt.Errorf("get tags: %w", err)
	}

	return tags, nil
}

// UpdateTag updates a tag in the DB.
func (r *repository) UpdateTag(tagID uint, value string) (*models.Tag, error) {
	var tag models.Tag
	if err := r.db.First(&tag, tagID).Error; err != nil {
		return nil, fmt.Errorf("update tag: %w", err)
	}

	tag.Value = value
	if err := r.db.Save(&tag).Error; err != nil {
		return nil, fmt.Errorf("update tag: %w", err)
	}

	return &tag, nil
}

// DeleteTag deletes a tag in the DB.
func (r *repository) DeleteTag(tagID uint) error {
	if err := r.db.Delete(&models.Tag{}, tagID).Error; err != nil {
		return fmt.Errorf("delete tag: %w", err)
	}

	return nil
}

// TagMember adds tags for a member.
func (r *repository) TagMember(memberID uint, tags []string) error {
	m, err := r.GetMember(memberID)
	if err != nil {
		return fmt.Errorf("tag member | get member: %w", err)
	}

	for _, tag := range tags {
		t, err := r.GetTag(tag)
		if t == nil {
			return ErrTagNotFound
		}

		if err != nil {
			return fmt.Errorf("tag member | get tags: %w", err)
		}

		m.Tags = append(m.Tags, *t)
	}

	if err := r.db.Save(&m).Error; err != nil {
		return fmt.Errorf("tag member | save: %w", err)
	}

	return nil
}
