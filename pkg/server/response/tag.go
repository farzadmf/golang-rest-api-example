package response

import (
	"assignment/pkg/models"
)

// This file contains types for tag responses.

// Tag is the type we use to send tag responses.
type Tag struct {
	ID    uint   `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

// FromModel converts the DB tag model to response type.
func (t *Tag) FromModel(tag *models.Tag) {
	t.ID = tag.ID
	t.Value = tag.Value
}
