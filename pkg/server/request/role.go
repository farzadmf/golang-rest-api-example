package request

import (
	"errors"
)

// This file contains request types for roles.

// CreateRole is the request to create a role.
type CreateRole struct {
	Name string `json:"name,omitempty"`
}

// IsValid checks the request to see if it's valid.
func (c *CreateRole) IsValid() (bool, error) {
	if c.Name == "" {
		return false, errors.New("name is required")
	}

	return true, nil
}
