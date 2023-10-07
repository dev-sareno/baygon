package mq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func getQueue(ch *amqp.Channel, queueName string) *amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Println("failed to declare a queue")
		return nil
	}
	return &q
}

func publish(ch *amqp.Channel, queueName string, data string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(data),
		})
	if err != nil {
		log.Println("failed to publish a message")
		return false
	}
	log.Printf(" [x] data has been sent to the queue %s\n", queueName)
	return true
}
