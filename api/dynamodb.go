package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	dynamoDBClient *dynamodb.Client
	tableName string = "picusv3"
)

func NewDynamoDBClient() *dynamodb.Client {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		fmt.Println("Region not set, using default region")
		region = "eu-west-1"
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatal("Failed to load AWS config:", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return client
}

func PutItem(objectID string, data interface{}) error {
	client := NewDynamoDBClient()

	// Convert the value directly to DynamoDB attribute value
	av, err := attributevalue.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal value to attribute value: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"objectID": &types.AttributeValueMemberS{Value: objectID},
			"data":     av,
		},
	}

	_, err = client.PutItem(context.TODO(), input)
	return err
}

func GetItem(objectID string) (interface{}, error) {
	client := NewDynamoDBClient()
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"objectID": &types.AttributeValueMemberS{Value: objectID},
		},
	}

	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, fmt.Errorf("item not found")
	}

	valueAttr, exists := result.Item["data"]
	if !exists {
		return nil, fmt.Errorf("data attribute not found")
	}

	// Unmarshal the attribute value directly to interface{}
	var value interface{}
	err = attributevalue.Unmarshal(valueAttr, &value)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal attribute value: %w", err)
	}

	return value, nil
}

func ListItems() ([]map[string]interface{}, error) {
	client := NewDynamoDBClient()
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := client.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var items []map[string]interface{}
	for _, item := range result.Items {
		keyAttr, okKey := item["objectID"].(*types.AttributeValueMemberS)
		if !okKey {
			continue
		}

		valueAttr := item["data"]
		var value interface{}
		err := attributevalue.Unmarshal(valueAttr, &value)
		if err != nil {
			continue
		}

		items = append(items, map[string]interface{}{
			"objectID":  keyAttr.Value,
			"data": value,
		})
	}

	return items, nil
}

func DeleteItem(objectID string) error {
	client := NewDynamoDBClient()
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"objectID": &types.AttributeValueMemberS{Value: objectID},
		},
	}

	_, err := client.DeleteItem(context.TODO(), input)
	return err
}
