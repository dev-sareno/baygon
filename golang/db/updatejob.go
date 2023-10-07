package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/dev-sareno/ginamus/dto"
	"log"
)

func UpdateJob(job *dto.Job) bool {
	svc := GetDynamoDbSession()

	update := expression.
		Set(expression.Name("lastActivityId"), expression.Value(job.LastActivityId)).
		Set(expression.Name("lastActivityIsOk"), expression.Value(job.LastActivityIsOk)).
		Set(expression.Name("lastActivityMessage"), expression.Value(job.LastActivityMessage)).
		Set(expression.Name("data.outputs"), expression.Value(job.Data.Outputs))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
		return false
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(job.Id),
			},
		},
		TableName:                 aws.String(TableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err = svc.UpdateItem(input)
	if err != nil {
		log.Printf("got error calling UpdateItem: %s\n", err)
		return false
	}

	log.Println("successfully updated.")
	return true
}
