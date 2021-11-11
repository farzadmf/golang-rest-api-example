package response

import (
	"assignment/pkg/models"
)

// This file contains types for employee responses.

// Employee is the type we use to send contractor responses.
type Employee struct {
	ID   uint     `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Role string   `json:"role,omitempty"`
	Tags []string `json:"tags"`
}

// FromModel converts the DB role model to response type.
func (e *Employee) FromModel(member *models.Member) {
	e.ID = member.ID
	e.Name = member.Name
	e.Role = member.Role.Name
	e.Tags = []string{}
	for _, t := range member.Tags {
		e.Tags = append(e.Tags, t.Value)
	}
}
