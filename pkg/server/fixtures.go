package server

import (
	"assignment/pkg/models"
	"assignment/pkg/server/response"
	"encoding/json"

	"gorm.io/gorm"
)

var (
	defaultRoles = map[uint]*models.Role{
		1: {
			Model: gorm.Model{ID: 1},
			Name:  "engineer",
		},
	}
	defaultMembers = map[uint]*models.Member{
		1: {
			Model:    gorm.Model{ID: 1},
			Name:     "contractor",
			Duration: 5,
			Tags: []models.Tag{
				*defaultTags[1],
			},
		},
		2: {
			Model: gorm.Model{ID: 2},
			Name:  "employee",
			Role:  defaultRoles[1],
			Tags: []models.Tag{
				*defaultTags[1],
			},
		},
	}
	defaultTags = map[uint]*models.Tag{
		1: {
			Model: gorm.Model{ID: 1},
			Value: "tag",
		},
	}
)

var (
	defaultRoleResp = response.Role{
		ID:   1,
		Name: "engineer",
	}
	defaultRoleRespStr string

	defaultRolesResp = []response.Role{
		defaultRoleResp,
	}
	defaultRolesRespStr string

	defaultTagResp = response.Tag{
		ID:    1,
		Value: "tag",
	}
	defaultTagRespStr string

	defaultTagsResp = []response.Tag{
		defaultTagResp,
	}
	defaultTagsRespStr string

	defaultContractorResp = response.Contractor{
		ID:       1,
		Name:     "contractor",
		Duration: 5,
		Tags:     []string{"tag"},
	}
	defaultContractorRespStr string

	defaultEmployeeResp = response.Employee{
		ID:   1,
		Name: "employee",
		Role: "engineer",
		Tags: []string{"tag"},
	}
	defaultEmployeeRespStr string

	defaultEmployeesResp = []response.Employee{
		defaultEmployeeResp,
	}
	defaultEmployeesRespStr string

	defaultContractorsResp = []response.Contractor{
		defaultContractorResp,
	}
	defaultContractorsRespStr string
)

var (
	members    = defaultMembers
	roles      = defaultRoles
	memberTags = defaultTags

	notFoundStr    = "Not Found"
	serverErrorStr = "server error"
	badRequestStr  = "invalid request"
)

func init() {
	// Ignoring errors; everything is fine!
	b, _ := json.Marshal(defaultRoleResp)
	defaultRoleRespStr = string(b)
	b, _ = json.Marshal(defaultRolesResp)
	defaultRolesRespStr = string(b)

	b, _ = json.Marshal(defaultTagResp)
	defaultTagRespStr = string(b)
	b, _ = json.Marshal(defaultTagsResp)
	defaultTagsRespStr = string(b)

	b, _ = json.Marshal(defaultContractorResp)
	defaultContractorRespStr = string(b)
	b, _ = json.Marshal(defaultContractorsResp)
	defaultContractorsRespStr = string(b)

	b, _ = json.Marshal(defaultEmployeeResp)
	defaultEmployeeRespStr = string(b)
	b, _ = json.Marshal(defaultEmployeesResp)
	defaultEmployeesRespStr = string(b)
}
