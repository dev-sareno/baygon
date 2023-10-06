package main

import (
	"fmt"
	"github.com/dev-sareno/ginamus/gin/handler"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"os"
)

func main() {
	// setup RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", err, "Failed to connect to RabbitMQ")
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	// init Gin web server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/", func(c *gin.Context) {
		handler.Post(c)
	})

	r.GET("/", func(c *gin.Context) {
		handler.Get(c)
	})

	if err := r.Run(":8000"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
