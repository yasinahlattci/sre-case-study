package db

import (
	"context"
	"sync"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	client *dynamodb.Client
	once         sync.Once
)

func GetClient(region string) *dynamodb.Client {

	once.Do(func() {
		
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
		)
		if err != nil {
			log.Fatal("Failed to load AWS config:", err)
		}
		
		client = dynamodb.NewFromConfig(cfg)
	})
	return client
}