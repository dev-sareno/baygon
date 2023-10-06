package handler

import (
	"context"
	"github.com/dev-sareno/ginamus/gin/mq"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"time"
)

func Post(c *gin.Context, ch *amqp091.Channel) {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	mq.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	mq.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	c.String(http.StatusOK, "Ok")
}

func Get(c *gin.Context, ch *amqp091.Channel) {

}
