package webhandler

import (
	"fmt"
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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
	if ok := mq.PublishToLookupCname(ch, &job); !ok {
		c.String(http.StatusInternalServerError, fmt.Sprint("Server error. Unable to publish job to mq"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"jobId": job.Id,
	})
}

func GetJobById(c *gin.Context, ch *amqp.Channel) {
	jobId := c.Param("jobId")

	// validate uuid
	if _, err := uuid.Parse(jobId); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid job id %s", jobId))
		return
	}

	item, err := db.GetJob(jobId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Server error. %s", err))
		return
	}
	log.Printf("%v\n", item)

	c.JSON(http.StatusOK, gin.H{})
}
