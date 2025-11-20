package handler

import (
	"context"
	"errors"
	"time"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	GetItem(ctx context.Context, objectID string) (interface{}, error)
	PutItem(ctx context.Context, objectID string, data interface{}) error
	ListItems(ctx context.Context) ([]map[string]interface{}, error)
}
type Handler struct {
	dbService ServiceInterface
	logger   *log.Logger
}

func NewHandler(s ServiceInterface, logger *log.Logger) *Handler {
	return &Handler{
		dbService: s,
		logger:    logger,
	}
}

func (h *Handler) GetItemHandler(c *fiber.Ctx) error {
	objectID := c.Params("objectID")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	value, err := h.dbService.GetItem(ctx, objectID)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.logger.Printf("Request timeout for objectID %s: %v", objectID, err)
			return c.Status(408).JSON(fiber.Map{
				"error": "Request timeout",
			})
		}
		h.logger.Printf("Error retrieving item %s: %v", objectID, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if value == nil {
		h.logger.Printf("Item not found for objectID %s", objectID)
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
		h.logger.Printf("Invalid request body: %v", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Data == nil {
		h.logger.Printf("Data field is required")
		return c.Status(400).JSON(fiber.Map{
			"error": "Data field is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := h.dbService.PutItem(ctx, objectID, req.Data)
	if err != nil {
		h.logger.Printf("Failed to put item %s: %v", objectID, err)
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
		h.logger.Printf("Failed to list items: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to list items",
		})
	}
	return c.Status(200).JSON(items)
}
