package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutItemRequest represents the request body for PUT operations
type PutItemRequest struct {
	Data interface{} `json:"data"`
}

func ListItemsHandler(c *fiber.Ctx) error {
	items, err := ListItems()

	if err != nil {
		log.Printf("Error listing items from DynamoDB: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to list items from database",
		})
	}
	return c.Status(200).JSON(items)
}

func PutItemHandler(c *fiber.Ctx) error {
	// Initialize DynamoDB client

	// Parse request body
	var req PutItemRequest

	requestKey := uuid.New().String()

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.Data == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Value is a required field",
		})
	}

	// Put item to DynamoDB directly without unnecessary JSON marshaling
	err := PutItem(requestKey, req.Data)
	if err != nil {
		log.Printf("Error putting item to DynamoDB: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to put item to database",
		})
	}

	log.Printf("Successfully put item with objectID: %s", requestKey)

	return c.Status(201).JSON(fiber.Map{
		"objectID": requestKey,
	})
}

func GetObjectHandler(c *fiber.Ctx) error {
	objectID := c.Params("objectID")
	if objectID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "objectID parameter is required",
		})
	}

	value, err := GetItem(objectID)
	if err != nil {
		log.Printf("Error getting item from DynamoDB: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get item from database",
		})
	}

	return c.Status(200).JSON(value)
}
