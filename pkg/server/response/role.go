package response

import (
	"assignment/pkg/models"
)

// This file contains types for role responses.

// Role is the type we use to send role responses.
type Role struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// FromModel converts the DB role model to response type.
func (r *Role) FromModel(role *models.Role) {
	r.ID = role.ID
	r.Name = role.Name
}
