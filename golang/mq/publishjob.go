package mq

import (
	"context"
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/dto"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func getQueue(ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		"jobs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Println("failed to declare a queue")
		return nil
	}
	return &q
}

func PublishJob(ch *amqp.Channel, job *dto.Job) bool {
	q := getQueue(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encodedJob := codec.Encode(job)
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(encodedJob),
		})
	if err != nil {
		log.Println("failed to publish a message")
		return false
	}
	log.Printf(" [x] Sent %s\n", encodedJob)
	return true
}
