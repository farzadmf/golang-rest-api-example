package server

import (
	"assignment/pkg/models"
)

type stubRepository struct{}

// ++++++++++++++++++++++++++ ROLES +++++++++++++++++++++++++++
// CreateRole creates a role in the DB.
func (s *stubRepository) CreateRole(name string) (*models.Role, error) {
	role := models.Role{
		Name: name,
	}
	role.ID = uint(len(roles)) + 1

	roles[role.ID] = &role
	return &role, nil
}

// GetRole returns the role matching the name.
func (s *stubRepository) GetRole(name string) (*models.Role, error) {
	for _, role := range roles {
		if role.Name == name {
			return role, nil
		}
	}

	return nil, nil
}

// GetRoles returns all the roles.
func (s *stubRepository) GetRoles() ([]*models.Role, error) {
	var r []*models.Role
	for _, role := range roles {
		r = append(r, role)
	}

	return r, nil
}

// UpdateRole updates a role in the DB.
func (s *stubRepository) UpdateRole(roleID uint, name string) (*models.Role, error) {
	roles[roleID].Name = name
	return roles[roleID], nil
}

// DeleteRole deletes a role in the DB.
func (s *stubRepository) DeleteRole(roleID uint) error {
	delete(roles, roleID)
	return nil
}

// +++++++++++++++++++++++++++ TAGS +++++++++++++++++++++++++++
// CreateTag creates a tag in the DB.
func (s *stubRepository) CreateTag(value string) (*models.Tag, error) {
	tag := models.Tag{
		Value: value,
	}
	tag.ID = uint(len(memberTags)) + 1

	memberTags[tag.ID] = &tag
	return &tag, nil
}

// GetTag returns the tag matching the value.
func (s *stubRepository) GetTag(value string) (*models.Tag, error) {
	for _, t := range memberTags {
		if t.Value == value {
			return t, nil
		}
	}

	return nil, nil
}

// GetTags returns all the tags.
func (s *stubRepository) GetTags() ([]*models.Tag, error) {
	var t []*models.Tag
	for _, tag := range memberTags {
		t = append(t, tag)
	}
	return t, nil
}

// UpdateTag updates a tag in the DB.
func (s *stubRepository) UpdateTag(tagID uint, value string) (*models.Tag, error) {
	memberTags[tagID].Value = value
	return memberTags[tagID], nil
}

// DeleteTag deletes a tag in the DB.
func (s *stubRepository) DeleteTag(tagID uint) error {
	delete(memberTags, tagID)
	return nil
}

// +++++++++++++++++++++++++ MEMBERS ++++++++++++++++++++++++++
// DeleteMember deletes a member in the DB.
func (s *stubRepository) DeleteMember(memberID uint) error {
	delete(members, memberID)
	return nil
}

// +++++++++++++++++++++++ CONTRACTORS ++++++++++++++++++++++++
// CreateContractor creates a contractor in the DB.
func (s *stubRepository) CreateContractor(name string, duration int) (*models.Member, error) {
	c := models.Member{
		Name:     name,
		Duration: duration,
	}
	c.ID = uint(len(members)) + 1

	members[c.ID] = &c
	return &c, nil
}

// GetContractor returns the contractor (member) matching the ID.
func (s *stubRepository) GetContractor(contractorID uint) (*models.Member, error) {
	if members[contractorID] != nil && members[contractorID].Type == models.Contractor {
		return members[contractorID], nil
	}

	return nil, nil
}

// GetContractors returns all the contractors.
func (s *stubRepository) GetContractors() ([]*models.Member, error) {
	var c []*models.Member
	for _, m := range members {
		if m.Type == models.Contractor {
			c = append(c, m)
		}
	}
	return c, nil
}

// ++++++++++++++++++++++++ EMPLOYEES +++++++++++++++++++++++++
// CreateEmployee creates an employee in the DB.
func (s *stubRepository) CreateEmployee(name string, roleID uint) (*models.Member, error) {
	e := models.Member{
		Name:   name,
		RoleID: &roleID,
	}
	e.ID = uint(len(members)) + 1

	members[e.ID] = &e
	return &e, nil
}

// GetEmployee returns the employee (member) matching the ID.
func (s *stubRepository) GetEmployee(employeeID uint) (*models.Member, error) {
	if members[employeeID] != nil && members[employeeID].Type == models.Employee {
		return members[employeeID], nil
	}

	return nil, nil
}

// GetEmployees returns all the employees.
func (s *stubRepository) GetEmployees() ([]*models.Member, error) {
	var c []*models.Member
	for _, m := range members {
		if m.Type == models.Contractor {
			c = append(c, m)
		}
	}
	return c, nil
}

// GetMember returns the member matching the ID.
func (s *stubRepository) GetMember(memberID uint) (*models.Member, error) {
	return members[memberID], nil
}

// TagMember adds tags for a member.
func (s *stubRepository) TagMember(memberID uint, tags []string) error {
	member := members[memberID]

	// Doing brute force here!
	for _, tag := range memberTags {
		for _, t := range tags {
			if tag.Value == t {
				member.Tags = append(member.Tags, *tag)
			}
		}
	}

	return nil
}

// Reset is a convenient function to reset the DB.
func (s *stubRepository) Reset() error {
	members = defaultMembers
	roles = defaultRoles
	memberTags = defaultTags

	return nil
}
