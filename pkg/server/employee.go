package server

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"assignment/pkg/models"
	"assignment/pkg/server/request"
	"assignment/pkg/server/response"
)

// This file contains routes for employees.

func (s *apiServer) addEmployeeRoutes() {
	s.app.Get("/employee", s.getEmployees)
	s.app.Get("/employee/:id", s.getEmployee)
	s.app.Post("/employee", s.createEmployee)
	s.app.Delete("/employee/:id", s.deleteEmployee)
}

// @Summary List all employees
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {array} response.Employee
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /employee [get]
func (s *apiServer) getEmployees(c *fiber.Ctx) error {
	employees, err := s.repo.GetEmployees()
	if err != nil {
		return s.serverError(c)
	}

	return employeesResponse(c, employees)
}

// @Summary Get an employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID" Format(int)
// @Success 200 {object} response.Employee
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /employee/{id} [get]
func (s *apiServer) getEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	member, err := s.repo.GetEmployee(uint(id))
	if err != nil {
		return s.serverError(c)
	}

	if member == nil {
		return s.notFound(c)
	}

	return employeeResponse(c, member, fiber.StatusOK)
}

// @Summary Create an employee
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body request.CreateEmployee true "Add employee"
// @Success 200 {object} response.Employee
// @Failure 400 {object} string "If either name or role is empty"
// @Failure 404 {object} string "If the specified role name cannot be found in the DB"
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /employee [post]
func (s *apiServer) createEmployee(c *fiber.Ctx) error {
	var req request.CreateEmployee
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return s.badRequest(c, err)
	}

	if ok, err := req.IsValid(); !ok {
		return s.badRequest(c, err)
	}

	role, err := s.repo.GetRole(req.Role)
	if err != nil {
		return s.serverError(c)
	}

	if role == nil {
		return s.notFound(c)
	}

	member, err := s.repo.CreateEmployee(req.Name, role.ID)
	if err != nil {
		return s.serverError(c)
	}

	return employeeResponse(c, member, fiber.StatusCreated)
}

// @Summary Delete an employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID" Format(int)
// @Success 200 {object} string
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /employee/{id} [delete]
func (s *apiServer) deleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	err = s.repo.DeleteMember(uint(id))
	if err != nil {
		return s.serverError(c)
	}

	return nil
}

func employeeResponse(
	c *fiber.Ctx,
	member *models.Member,
	statusCode int,
) error {
	var resp response.Employee
	resp.FromModel(member)

	err := c.JSON(resp)
	if err != nil {
		return err
	}
	return c.SendStatus(statusCode)
}

func employeesResponse(c *fiber.Ctx, members []*models.Member) error {
	var resp []response.Employee
	for _, m := range members {
		var employee response.Employee
		employee.FromModel(m)

		resp = append(resp, employee)
	}

	return c.JSON(resp)
}
