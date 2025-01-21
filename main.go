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
	r.POST("/post", callPost)
	r.POST("/post-form-data", callPostFormData)

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

func callPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

type FormData struct {
	Name string `form:"name"`
	Age  uint   `form:"age"`
}

func callPostFormData(c *gin.Context) {
	var form FormData
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("form: ", form)
	c.JSON(200, gin.H{
		"message": "post form data " + form.Name,
	})
}
