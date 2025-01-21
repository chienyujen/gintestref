package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCallPing(t *testing.T) {
	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the function
	callPing(c)
	fmt.Println(w.Body.String())
	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "pong"}`, w.Body.String())
}

func TestMiddleware(t *testing.T) {
	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the function
	middleware(c)

	// Check the response
	assert.Empty(t, w.Body.String())
}

func TestCallPost(t *testing.T) {
	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the function
	callPost(c)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "post"}`, w.Body.String())
}

func TestCallPostFormData(t *testing.T) {
	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the request body
	c.Request = httptest.NewRequest("POST", "/post-form-data", nil)
	c.Request.Header.Set("Content-Type", "application/form-data")
	c.Request.Form = make(map[string][]string)
	c.Request.Form.Add("name", "John")
	c.Request.Form.Add("age", "30")

	// Call the function
	callPostFormData(c)
	fmt.Println(w.Body.String())
	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "post form data John"}`, w.Body.String())
}
