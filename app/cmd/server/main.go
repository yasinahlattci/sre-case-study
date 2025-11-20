package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/yasinahlattci/sre-case-study/app/internal/api/handler"
	"github.com/yasinahlattci/sre-case-study/app/internal/api/router"
	"github.com/yasinahlattci/sre-case-study/app/internal/bootstrap"
)

func main() {

	deps, err := bootstrap.Bootstrap(os.Getenv("APP_ENV"))
	if err != nil {
		log.Fatalf("Failed to bootstrap dependencies: %v", err)
	}

	app := fiber.New()

	handler := handler.NewHandler(deps.Service, deps.Logger)

	router.SetupRoutes(app, handler)

	port := deps.Config.Server.Port
	if err := app.Listen(port); err != nil {
		deps.Logger.Fatalf("Failed to start server: %v", err)
	}

}
