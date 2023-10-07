package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dev-sareno/ginamus/dto"
	"github.com/google/uuid"
	"log"
	"time"
)

func CreateJob(domains []string) (dto.Job, bool) {
	svc := GetDynamoDbSession()

	item := dto.Job{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().Format(time.RFC3339),
		Data: dto.JobData{
			Type: 0, // we only support type 0 for now
			Input: dto.JobInput{
				Domains: domains,
				Filler:  [][]string{}, // TODO: implement
			},
		},
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Printf("got error marshalling item: %s\n", err)
		return dto.Job{}, false
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("got error calling PutItem: %s\n", err)
		return dto.Job{}, false
	}

	fmt.Printf("job has been created in the db. id: %s, createdAt: %s\n", item.Id, item.CreatedAt)
	return item, true
}
