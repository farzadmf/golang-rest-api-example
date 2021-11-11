package response

import "assignment/pkg/models"

// This file contains types for contractor responses.

// Contractor is the type we use to send contractor responses.
type Contractor struct {
	ID       uint     `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Duration int      `json:"duration,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

// FromModel converts the DB role model to response type.
func (c *Contractor) FromModel(member *models.Member) {
	c.ID = member.ID
	c.Name = member.Name
	c.Duration = member.Duration
	for _, t := range member.Tags {
		c.Tags = append(c.Tags, t.Value)
	}
}
