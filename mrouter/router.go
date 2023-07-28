package mrouter

import (
	"github.com/dev-sareno/ginamus/mhtml"
	"github.com/dev-sareno/ginamus/mtypus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", getBase)
	r.POST("/do", postDo)
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Page not found")
	})
	return r
}

func getBase(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, mhtml.GetMainPage())
}

func postDo(c *gin.Context) {
	var body mtypus.RequestPostDo
	if err := c.ShouldBind(&body); err != nil {
		c.String(http.StatusBadRequest, "Invalid body")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": body.Input,
	})
}
