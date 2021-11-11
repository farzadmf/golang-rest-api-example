package request

import (
	"errors"
)

// This file contains request types for employees.

// CreateEmployee is the request to create a employee.
type CreateEmployee struct {
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

// IsValid checks the request to see if it's valid.
func (c *CreateEmployee) IsValid() (bool, error) {
	if c.Name == "" || c.Role == "" {
		return false, errors.New("name and role are required")
	}

	return true, nil
}
