package main

import (
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App) {
	api := app.Group("/picus")

	api.Get("list", ListItemsHandler)
	api.Get("get", ListItemsHandler)

	api.Post("put", PutItemHandler)
	api.Get("get/:objectID", GetObjectHandler)

}
