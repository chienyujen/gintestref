package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route()
}

func route() {
	r := gin.Default()
	r.GET("/ping", callPing)

	r.Run(":8080")
}

func callPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
