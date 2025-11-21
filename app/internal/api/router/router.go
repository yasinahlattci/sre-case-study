package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yasinahlattci/sre-case-study/app/internal/api/handler"
)

func SetupRoutes(app *fiber.App, h *handler.Handler) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	api := app.Group("/picus")

	api.Get("list", h.ListItemsHandler)

	api.Post("put", h.PutItemHandler)
	api.Get("get/:objectID", h.GetItemHandler)
}
