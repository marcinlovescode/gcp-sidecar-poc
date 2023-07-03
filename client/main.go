package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingHandler() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"userId": c.GetHeader("X-User-Id"),
		})
	}
}

func main() {
	router := gin.Default()
	router.GET("/ping", pingHandler())
	router.Run(":8090")
}