package handler

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yasinahlattci/sre-case-study/api/internal/service"
)

type Handler struct {
	dbService *service.DynamoDBService
}

func NewHandler(s *service.DynamoDBService) *Handler {
	return &Handler{
		dbService: s,
	}
}

func (h *Handler) GetItemHandler(c *fiber.Ctx) error {
	objectID := c.Params("objectID")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	value, err := h.dbService.GetItem(ctx, objectID)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(408).JSON(fiber.Map{
				"error": "Request timeout",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if value == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Item not found",
		})
	}
	return c.Status(200).JSON(value)
}

func (h *Handler) PutItemHandler(c *fiber.Ctx) error {
	objectID := uuid.New().String()

	var req struct {
		Data interface{} `json:"data"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Data == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Data field is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := h.dbService.PutItem(ctx, objectID, req.Data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to put item",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"objectID": objectID,
	})
}

func (h *Handler) ListItemsHandler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	items, err := h.dbService.ListItems(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to list items",
		})
	}
	return c.Status(200).JSON(items)
}
