package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "automatic/backEnd/config"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.Run(":8080")
}
