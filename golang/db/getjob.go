package db

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dev-sareno/ginamus/dto"
)

func GetJob(jobId string) (*dto.Job, error) {
	svc := GetDynamoDbSession()

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(jobId),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find '" + jobId + "'"
		return nil, errors.New(msg)
	}

	var item dto.Job
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Record, %v", err)
	}
	return &item, nil
}
