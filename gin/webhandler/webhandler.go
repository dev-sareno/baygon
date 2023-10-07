package webhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"time"
)

func Post(c *gin.Context, ch *amqp091.Channel) {
	type Request struct {
		Domains []string `json:"domains" binding:"required"`
	}

	var body Request
	if err := c.ShouldBind(&body); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Bad request %s", err.Error()))
		return
	}

	fmt.Println(body.Domains)

	d, err := json.Marshal(&body.Domains)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Server error %s", err))
		return
	}

	if err := db.CreateJob(body.Domains); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Server error %s", err))
		return
	}

	q, err := ch.QueueDeclare(
		"jobs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	mq.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mqbody := string(d)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(mqbody),
		})
	mq.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	c.String(http.StatusOK, "Ok")
}

func Get(c *gin.Context, ch *amqp091.Channel) {

}
