package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	route()
}

func route() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(middleware)
	r.GET("/ping", middleware, callPing)

	r.Run(":8080")
}

func middleware(c *gin.Context) {
	fmt.Println("I'm a middleware!")
	c.Next()
}

func callPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
