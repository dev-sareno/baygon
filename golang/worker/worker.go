package worker

import (
	"github.com/dev-sareno/ginamus/context"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/dev-sareno/ginamus/workerhandler"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func Run() {
	ch, mqClose, ok := mq.GetChannel()
	if !ok {
		log.Println("failed to connect to RabbitMQ.")
		return
	}
	defer mqClose()

	var q *amqp.Queue
	lt := os.Getenv("WORKER_DNS_LOOKUP_TYPE")
	switch lt {
	case "A":
		log.Println("worker mode for lookup A")
		q = mq.GetLookupAQueue(ch)
		break
	case "CNAME":
		log.Println("worker mode for lookup CNAME")
		q = mq.GetLookupCnameQueue(ch)
		break
	default:
		log.Printf("unsupported lookup type %s\n", lt)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("failed to register a consumer")
		return
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			jobCtx := context.WorkerContext{
				Db:        db.GetDynamoDbSession(),
				MqChannel: ch,
				Job:       nil,
			}
			workerhandler.HandleJob(&jobCtx, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
