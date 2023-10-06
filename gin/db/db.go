package db

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"log"
	"time"
)

const TableName = "Jobs"

func GetDynamoDbSession() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return svc
}

func Init() {
	svc := GetDynamoDbSession()

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("CreatedAt"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("CreatedAt"),
				KeyType:       aws.String("RANGE"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   aws.String(TableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		var aerr awserr.Error
		if errors.As(err, &aerr) {
			if aerr.Code() == dynamodb.ErrCodeResourceInUseException {
				log.Println("Table already exists. Skip.")
				return
			} else {
				log.Fatalf("Got error calling CreateTable: %s", err)
			}
		}
	}

	fmt.Println("Created the table", TableName)
}

func CreateJob(domains []string) error {
	svc := GetDynamoDbSession()

	item := Job{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Format(time.RFC3339),
		Data: JobData{
			Type: 0,
			Input: JobInput{
				Domains: domains,
				Filler:  [][]string{}, // TODO: implement
			},
		},
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("got error calling PutItem: %s", err)
	}

	fmt.Printf("ID: %s, CreatedAt: %s\n", item.ID, item.CreatedAt)
	return nil
}
