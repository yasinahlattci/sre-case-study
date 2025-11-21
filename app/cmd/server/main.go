package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	handler := handler.NewHandler(deps.Service, deps.Logger, deps.Config.Server.RequestTimeout)

	router.SetupRoutes(app, handler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		deps.Logger.Println("Gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(ctx); err != nil {
			deps.Logger.Printf("Error during shutdown: %v", err)
		}
	}()

	port := deps.Config.Server.Port
	deps.Logger.Printf("Server starting on port %s", port)

	if err := app.Listen(port); err != nil {
		deps.Logger.Fatalf("Failed to start server: %v", err)
	}
}
