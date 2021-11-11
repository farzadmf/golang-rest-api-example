package request

import (
	"errors"
)

// This file contains request types for contractors.

// CreateContractor is the request to create a contractor.
type CreateContractor struct {
	Name     string `json:"name,omitempty"`
	Duration int    `json:"duration,string,omitempty"`
}

// IsValid checks the request to see if it's valid.
func (c *CreateContractor) IsValid() (bool, error) {
	if c.Name == "" || c.Duration == 0 {
		return false, errors.New("name and duration are required")
	}

	return true, nil
}
