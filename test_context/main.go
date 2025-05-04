package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const (
	ctxUserID ctxKey = "userID"
)

func main() {
	r := gin.Default()

	r.GET("", testContext)

	r.Run(":8080")
}

func testContext(c *gin.Context) {
	userID := "abcd1234"

	ctxWithUserID := context.WithValue(c.Request.Context(), ctxUserID, userID)
	newService(ctxWithUserID)
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
func newService(c context.Context) {
	// Extract the userID from the context
	userID, ok := c.Value(ctxUserID).(string)
	if !ok {
		// Handle the case where userID is not present in the context
		return
	}
	fmt.Println(userID)
	// Use the userID for your service logic
}
