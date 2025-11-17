package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBService struct {
	client *dynamodb.Client
	table  string
}

func NewDynamoDBService(client *dynamodb.Client, table string) *DynamoDBService {
	return &DynamoDBService{
		client: client,
		table:  table,
	}
}

func (s *DynamoDBService) PutItem(ctx context.Context, objectID string, data interface{}) error {

	av, err := attributevalue.Marshal(data)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: &s.table,
		Item: map[string]types.AttributeValue{
			"objectID": &types.AttributeValueMemberS{Value: objectID},
			"data":     av,
		},
	}

	_, err = s.client.PutItem(ctx, input)
	return err
}

func (s *DynamoDBService) GetItem(ctx context.Context, objectID string) (interface{}, error) {
	input := &dynamodb.GetItemInput{
		TableName: &s.table,
		Key: map[string]types.AttributeValue{
			"objectID": &types.AttributeValueMemberS{Value: objectID},
		},
	}
	result, err := s.client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	valueAttr, exists := result.Item["data"]
	if !exists {
		return nil, nil
	}
	var value interface{}
	err = attributevalue.Unmarshal(valueAttr, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *DynamoDBService) ListItems(ctx context.Context) ([]map[string]interface{}, error) {
	input := &dynamodb.ScanInput{
		TableName: &s.table,
	}
	result, err := s.client.Scan(ctx, input)
	if err != nil {
		return nil, err
	}
	var items []map[string]interface{}
	for _, item := range result.Items {
		keyAttr, exists := item["objectID"].(*types.AttributeValueMemberS)
		if !exists {
			continue
		}
		valueAttr, exists := item["data"]
		if !exists {
			continue
		}
		var value interface{}
		err = attributevalue.Unmarshal(valueAttr, &value)
		if err != nil {
			continue
		}
		items = append(items, map[string]interface{}{
			"objectID": keyAttr.Value,
			"data":     value,
		})
	}
	return items, nil
}
