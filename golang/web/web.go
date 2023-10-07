package web

import (
	"fmt"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/mq"
	"github.com/dev-sareno/ginamus/webhandler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Run() {
	ch, mqClose, ok := mq.GetChannel()
	if !ok {
		log.Println("failed to connect to RabbitMQ.")
		return
	}
	defer mqClose()

	if ok := db.PrepareTable(); !ok {
		log.Println("failed to prepare db.")
		return
	}

	// init Gin web server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/", func(c *gin.Context) {
		webhandler.Post(c, ch)
	})

	r.GET("/", func(c *gin.Context) {
		webhandler.Get(c, ch)
	})

	if err := r.Run(":8000"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
