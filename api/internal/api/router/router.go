package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yasinahlattci/sre-case-study/api/internal/api/handler"
)

func SetupRoutes(app *fiber.App, h *handler.Handler) {

	api := app.Group("/picus")

	api.Get("list", h.ListItemsHandler)
	api.Get("get", h.ListItemsHandler)

	api.Post("put", h.PutItemHandler)
	api.Get("get/:objectID", h.GetItemHandler)
}