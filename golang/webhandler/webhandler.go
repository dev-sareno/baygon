package webhandler

import (
	"fmt"
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)

func Post(c *gin.Context, ch *amqp.Channel) {
	type Request struct {
		Domains []string `json:"domains" binding:"required"`
	}

	var body Request
	if err := c.ShouldBind(&body); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Bad request %s", err.Error()))
		return
	}

	fmt.Println(body.Domains)

	// create job in db
	job, ok := db.CreateJob(body.Domains)
	codec.Encode(&job)
	if !ok {
		c.String(http.StatusInternalServerError, fmt.Sprint("Server error. Unable to create job in db"))
		return
	}

	// publish job to mq
	if ok := mq.PublishToLookupA(ch, &job); !ok {
		c.String(http.StatusInternalServerError, fmt.Sprint("Server error. Unable to publish job to mq"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"jobId": job.Id,
	})
}

func Get(c *gin.Context, ch *amqp.Channel) {

}
