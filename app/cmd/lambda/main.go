package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yasinahlattci/sre-case-study/app/internal/bootstrap"
)

func albResponse(statusCode int, success bool, message string) events.ALBTargetGroupResponse {
	b, _ := json.Marshal(map[string]interface{}{
		"success": success,
		"message": message,
	})
	return events.ALBTargetGroupResponse{
		StatusCode:        statusCode,
		StatusDescription: fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
		Body:              string(b),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func deleteHandler(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {

	deps, err := bootstrap.Bootstrap(os.Getenv("APP_ENV"))

	if err != nil {
		deps.Logger.Printf("Failed to bootstrap dependencies: %v", err)
		return albResponse(500, false, "Internal server error"), nil
	}

	path := request.Path
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(pathParts) != 2 || pathParts[0] != "picus" {
		return albResponse(400, false, "Invalid request path"), nil
	}

	objectID := pathParts[1]

	existingItem, err := deps.Service.GetItem(ctx, objectID)
	if err != nil {
		deps.Logger.Printf("Error checking item: %v", err)
		return albResponse(500, false, "Internal server error"), nil
	}

	if existingItem == nil {
		return albResponse(404, false, fmt.Sprintf("Item with objectID %s not found", objectID)), nil
	}

	if err := deps.Service.DeleteItem(ctx, objectID); err != nil {
		deps.Logger.Printf("Error deleting item: %v", err)
		return albResponse(500, false, "Internal server error"), nil
	}

	return albResponse(200, true, "Item deleted successfully"), nil
}

func main() {
	lambda.Start(deleteHandler)
}
