package request

import (
	"errors"
)

// This file contains request types for tags.

// CreateTag is the request to create a tag.
type CreateTag struct {
	Value string `json:"value,omitempty"`
}

// IsValid checks the request to see if it's valid.
func (c *CreateTag) IsValid() (bool, error) {
	if c.Value == "" {
		return false, errors.New("value is required")
	}

	return true, nil
}

// MemberTag is the request to create a tag.
type MemberTag struct {
	Values []string `json:"values,omitempty"`
}

// IsValid checks the request to see if it's valid.
func (mt *MemberTag) IsValid() (bool, error) {
	if len(mt.Values) == 0 {
		return false, errors.New("at least one tag is required")
	}

	return true, nil
}
