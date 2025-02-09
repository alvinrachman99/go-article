package handler

import (
	"article/config"
	"article/models"
	"article/repository"
	"article/service"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostsHandler struct {
	PostsService *service.PostService
}

func NewPostsHandler(db *sql.DB, cfg config.Config) *PostsHandler {
	postsRepo := repository.NewPostsRepository(db)
	postsService := service.NewPostsService(postsRepo)

	return &PostsHandler{PostsService: postsService}
}

func (h *PostsHandler) Create(c *fiber.Ctx) error {
	var posts models.Posts
	if err := c.BodyParser(&posts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := h.PostsService.Create(&posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create Success", "status": fiber.StatusOK})
}

func (h *PostsHandler) GetPosts(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
	}

	posts, err := h.PostsService.GetPosts(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": fiber.StatusOK, "payload": posts})
}

func (h *PostsHandler) GetPostsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	posts, err := h.PostsService.GetPostsById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": fiber.StatusOK, "payload": posts})
}

func (h *PostsHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var posts *models.Posts
	if err := c.BodyParser(&posts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}
	posts.Id = id

	if err := h.PostsService.Update(posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Update Success", "status": fiber.StatusOK})
}

func (h *PostsHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.PostsService.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Delete Success", "status": fiber.StatusOK})
}
