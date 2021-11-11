package server

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"assignment/pkg/models"
	"assignment/pkg/server/request"
	"assignment/pkg/server/response"
)

// This file contains routes for contractors.

func (s *apiServer) addContractorRoutes() {
	s.app.Get("/contractor", s.getContractors)
	s.app.Get("/contractor/:id", s.getContractor)
	s.app.Post("/contractor", s.createContractor)
	s.app.Delete("/contractor/:id", s.deleteContractor)
}

// @Summary List all contractors
// @Tags contractors
// @Accept json
// @Produce json
// @Success 200 {array} response.Contractor
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /contractor [get]
func (s *apiServer) getContractors(c *fiber.Ctx) error {
	contractors, err := s.repo.GetContractors()
	if err != nil {
		return s.serverError(c)
	}

	return contractorsResponse(c, contractors)
}

// @Summary Get a contractor by ID
// @Tags contractors
// @Accept json
// @Produce json
// @Failure 500 {object} string "When there's an issue on the server side"
// @Param id path int true "Contractor ID" Format(int)
// @Success 200 {object} response.Contractor
// @Router /contractor/{id} [get]
func (s *apiServer) getContractor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	member, err := s.repo.GetContractor(uint(id))
	if err != nil {
		return s.serverError(c)
	}

	if member == nil {
		return s.notFound(c)
	}

	return contractorResponse(c, member, fiber.StatusOK)
}

// @Summary Create a contractor
// @Tags contractors
// @Accept json
// @Produce json
// @Param contractor body request.CreateContractor true "Add contractor"
// @Success 200 {object} response.Contractor
// @Failure 400 {object} string "If either name or duration is empty"
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /contractor [post]
func (s *apiServer) createContractor(c *fiber.Ctx) error {
	var req request.CreateContractor
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return s.badRequest(c, err)
	}

	if ok, err := req.IsValid(); !ok {
		return s.badRequest(c, err)
	}

	member, err := s.repo.CreateContractor(req.Name, req.Duration)
	if err != nil {
		return s.serverError(c)
	}

	return contractorResponse(c, member, fiber.StatusCreated)
}

// @Summary Delete a contractor by ID
// @Tags contractors
// @Accept json
// @Produce json
// @Param id path int true "Contractor ID" Format(int)
// @Success 200 {object} string
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /contractor/{id} [delete]
func (s *apiServer) deleteContractor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	err = s.repo.DeleteMember(uint(id))
	if err != nil {
		return s.badRequest(c, err)
	}

	return nil
}

func contractorResponse(
	c *fiber.Ctx,
	member *models.Member,
	statusCode int,
) error {
	var resp response.Contractor
	resp.FromModel(member)

	err := c.JSON(resp)
	if err != nil {
		return err
	}
	return c.SendStatus(statusCode)
}

func contractorsResponse(c *fiber.Ctx, members []*models.Member) error {
	var resp []response.Contractor
	for _, m := range members {
		var contractor response.Contractor
		contractor.FromModel(m)

		resp = append(resp, contractor)
	}

	return c.JSON(resp)
}
