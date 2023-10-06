package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Post(c *gin.Context) {

}

func Get(c *gin.Context) {

}
