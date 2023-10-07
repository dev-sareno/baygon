package context

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dev-sareno/ginamus/dto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerContext struct {
	Db        *dynamodb.DynamoDB
	MqChannel *amqp.Channel
	Job       *dto.Job
}
