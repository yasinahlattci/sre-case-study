package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/yasinahlattci/sre-case-study/api/internal/api/handler"
	"github.com/yasinahlattci/sre-case-study/api/internal/api/router"
	"github.com/yasinahlattci/sre-case-study/api/internal/config"
	"github.com/yasinahlattci/sre-case-study/api/internal/db"
	"github.com/yasinahlattci/sre-case-study/api/internal/service"
)

func main() {

	config, err := config.LoadConfig(".conf", os.Getenv("APP_ENV"))

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app := fiber.New()

	dynamoClient := db.GetClient(config.Database.Region)
	tableName := config.Database.TableName

	service := service.NewDynamoDBService(dynamoClient, tableName)
	handler := handler.NewHandler(service)

	router.SetupRoutes(app, handler)

	port := config.Server.Port
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
