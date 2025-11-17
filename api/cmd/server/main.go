package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
	"github.com/yasinahlattci/sre-case-study/api/internal/api/handler"
	"github.com/yasinahlattci/sre-case-study/api/internal/api/router"
	"github.com/yasinahlattci/sre-case-study/api/internal/db"
	"github.com/yasinahlattci/sre-case-study/api/internal/service"
)

func main() {

	app := fiber.New()

	dynamoClient := db.GetClient("eu-west-1")
	tableName := "picusv3"

	service := service.NewDynamoDBService(dynamoClient, tableName)
	handler := handler.NewHandler(service)
	
	router.SetupRoutes(app, handler)

	port := ":3000"
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}