package main

import (
	"fmt"
	"github.com/dev-sareno/ginamus/gin/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/", handler.Post)
	r.GET("/", handler.Get)

	if err := r.Run(":8000"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
