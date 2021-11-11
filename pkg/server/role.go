package server

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"assignment/pkg/models"
	"assignment/pkg/server/request"
	"assignment/pkg/server/response"
)

// This file contains routes for roles.

func (s *apiServer) addRoleRoutes() {
	s.app.Get("/role", s.getRoles)
	s.app.Get("/role/:name", s.getRole)
	s.app.Post("/role", s.createRole)
	s.app.Delete("/role/:id", s.deleteRole)
}

// @Summary List all roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {array} response.Role
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /role [get]
func (s *apiServer) getRoles(c *fiber.Ctx) error {
	roles, err := s.repo.GetRoles()
	if err != nil {
		return s.serverError(c)
	}

	return rolesResponse(c, roles)
}

// @Summary Get a role by name
// @Tags roles
// @Accept json
// @Produce json
// @Param name path string true "Role name" Format(string)
// @Success 200 {object} response.Role
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /role/{name} [get]
func (s *apiServer) getRole(c *fiber.Ctx) error {
	name := c.Params("name")

	role, err := s.repo.GetRole(name)
	if err != nil {
		return s.serverError(c)
	}

	if role == nil {
		return s.notFound(c)
	}

	return roleResponse(c, role, fiber.StatusOK)
}

// @Summary Create a role
// @Tags roles
// @Accept json
// @Produce json
// @Param contractor body request.CreateRole true "Add role"
// @Success 200 {object} response.Role
// @Failure 400 {object} string "If name is empty"
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /role [post]
func (s *apiServer) createRole(c *fiber.Ctx) error {
	var req request.CreateRole
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return s.badRequest(c, err)
	}

	if ok, err := req.IsValid(); !ok {
		return s.badRequest(c, err)
	}

	role, err := s.repo.CreateRole(req.Name)
	if err != nil {
		return s.serverError(c)
	}

	return roleResponse(c, role, fiber.StatusCreated)
}

// @Summary Delete a role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID" Format(int)
// @Success 200 {object} string
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /role/{id} [delete]
func (s *apiServer) deleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	err = s.repo.DeleteRole(uint(id))
	if err != nil {
		return s.serverError(c)
	}

	return nil
}

func roleResponse(
	c *fiber.Ctx,
	role *models.Role,
	statusCode int,
) error {
	var resp response.Role
	resp.FromModel(role)

	err := c.JSON(resp)
	if err != nil {
		return err
	}
	return c.SendStatus(statusCode)
}

func rolesResponse(c *fiber.Ctx, roles []*models.Role) error {
	var resp []response.Role
	for _, r := range roles {
		var role response.Role
		role.FromModel(r)

		resp = append(resp, role)
	}

	return c.JSON(resp)
}
