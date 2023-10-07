package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dev-sareno/ginamus/dto"
	"log"
	"net/http"
)

func GetJob(jobId string) (*dto.Job, int) {
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
		log.Printf("got error calling GetItem: %s", err)
		return nil, http.StatusInternalServerError
	}

	if result.Item == nil {
		return nil, http.StatusNotFound
	}

	var item dto.Job
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Printf("failed to unmarshal Record, %v", err)
		return nil, http.StatusInternalServerError
	}
	return &item, http.StatusOK
}
