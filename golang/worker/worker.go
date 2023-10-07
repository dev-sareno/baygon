package worker

import (
	"fmt"
	"github.com/dev-sareno/ginamus/context"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/dev-sareno/ginamus/workerhandler"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func Run() {
	// setup RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	mq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	ch, err := conn.Channel()
	mq.FailOnError(err, "Failed to open a channel")
	defer func() {
		if err := ch.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	q, err := ch.QueueDeclare(
		"jobs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	mq.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	mq.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			jobCtx := context.WorkerContext{
				Db:        db.GetDynamoDbSession(),
				MqChannel: ch,
				Job:       nil,
			}
			workerhandler.HandleJob(jobCtx, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
