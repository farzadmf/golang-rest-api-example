package server

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"assignment/pkg/models"
	"assignment/pkg/repository"
	"assignment/pkg/server/request"
	"assignment/pkg/server/response"
)

// This file contains routes for roles.

func (s *apiServer) addTagRoutes() {
	s.app.Get("/tag", s.getTags)
	s.app.Get("/role/:value", s.getTag)
	s.app.Post("/tag", s.createTag)
	s.app.Delete("/tag/:id", s.deleteTag)
	s.app.Post("/member/:id/tags", s.memberTags)
}

// @Summary List all tags
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {array} response.Tag
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /tag [get]
func (s *apiServer) getTags(c *fiber.Ctx) error {
	tags, err := s.repo.GetTags()
	if err != nil {
		return s.serverError(c)
	}

	return tagsResponse(c, tags)
}

// @Summary Get a tag by value
// @Tags tags
// @Accept json
// @Produce json
// @Param name path string true "Tag value" Format(string)
// @Success 200 {object} response.Tag
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /tag/{value} [get]
func (s *apiServer) getTag(c *fiber.Ctx) error {
	value := c.Params("value")

	tag, err := s.repo.GetTag(value)
	if err != nil {
		return s.serverError(c)
	}

	if tag == nil {
		return s.notFound(c)
	}

	return tagResponse(c, tag, fiber.StatusOK)
}

// @Summary Create a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param tag body request.CreateTag true "Add tag"
// @Success 200 {object} response.Tag
// @Failure 400 {object} string "If value is empty"
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /tag [post]
func (s *apiServer) createTag(c *fiber.Ctx) error {
	var req request.CreateTag
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return s.badRequest(c, err)
	}

	if ok, err := req.IsValid(); !ok {
		return s.badRequest(c, err)
	}

	tag, err := s.repo.CreateTag(req.Value)
	if err != nil {
		return s.serverError(c)
	}

	return tagResponse(c, tag, fiber.StatusCreated)
}

// @Summary Delete a tag by ID
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID" Format(int)
// @Success 200 {object} string
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /tag/{id} [delete]
func (s *apiServer) deleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	err = s.repo.DeleteTag(uint(id))
	if err != nil {
		return s.badRequest(c, err)
	}

	return nil
}

// @Summary Add tags to a member
// @Tags tags
// @Accept json
// @Produce json
// @Failure 400 {object} string "If no tag is specified"
// @Param id path int true "Member ID" Format(int)
// @Param tags body request.MemberTag true "Member Tags"
// @Success 200 {object} string
// @Failure 400 {object} string "When not all the specified tags exist in the DB"
// @Failure 500 {object} string "When there's an issue on the server side"
// @Router /member/{id}/tags [post]
func (s *apiServer) memberTags(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return s.badRequest(c, err)
	}

	var req request.MemberTag
	err = json.Unmarshal(c.Body(), &req)
	if err != nil {
		return s.badRequest(c, err)
	}

	if ok, err := req.IsValid(); !ok {
		return s.badRequest(c, err)
	}

	m, err := s.repo.GetMember(uint(id))
	if err != nil {
		return s.serverError(c)
	}

	if m == nil {
		return s.notFound(c)
	}

	err = s.repo.TagMember(m.ID, req.Values)
	if errors.Is(err, repository.ErrTagNotFound) {
		return s.badRequest(c, errors.New("all the specified tags should be present in the database"))
	}

	if err != nil {
		return s.serverError(c)
	}

	return nil
}

func tagResponse(
	c *fiber.Ctx,
	tag *models.Tag,
	statusCode int,
) error {
	var resp response.Tag
	resp.FromModel(tag)

	err := c.JSON(resp)
	if err != nil {
		return err
	}
	c.SendStatus(statusCode)

	return c.SendStatus(fiber.StatusOK)
}

func tagsResponse(c *fiber.Ctx, roles []*models.Tag) error {
	var resp []response.Tag
	for _, r := range roles {
		var role response.Tag
		role.FromModel(r)

		resp = append(resp, role)
	}

	return c.JSON(resp)
}
