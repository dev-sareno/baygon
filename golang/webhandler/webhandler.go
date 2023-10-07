package webhandler

import (
	"fmt"
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if ok := mq.PublishToLookupCname(ch, &job); !ok {
		c.String(http.StatusInternalServerError, fmt.Sprint("Server error. Unable to publish job to mq"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"jobId": job.Id,
	})
}

func GetJobById(c *gin.Context, ch *amqp.Channel) {
	type Item struct {
		Domain string `json:"domain"`
		A      string `json:"a"`
		Cname  string `json:"cname"`
	}
	type Response struct {
		JobId     string  `json:"jobId"`
		Completed bool    `json:"completed"`
		Progress  float64 `json:"progress"`
		Data      []Item  `json:"data"`
	}

	jobId := c.Param("jobId")

	// validate uuid
	if _, err := uuid.Parse(jobId); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid job id %s", jobId))
		return
	}

	job, statusCode := db.GetJob(jobId)
	if statusCode != http.StatusOK {
		c.String(statusCode, fmt.Sprint("Server error."))
		return
	}

	response := Response{JobId: jobId}

	// simplify response
	response.Progress = (float64(len(job.Data.Outputs)) / 2) * 100 // divide two because we only have A and CNAME resolution
	response.Completed = response.Progress >= 100

	const CNAME = 0
	const A = 1
	response.Data = []Item{}
	for i, domain := range job.Data.Input.Domains {
		v := Item{Domain: domain}
		jo := job.Data.Outputs
		if len(jo) > 0 {
			// cname is available
			v.Cname = jo[CNAME].Data[i]
		}
		if len(jo) > 1 {
			// a is available
			v.A = jo[A].Data[i]
		}
		response.Data = append(response.Data, v)
	}

	c.JSON(http.StatusOK, response)
}
