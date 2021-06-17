package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTextContext(response)

	Ping(context)

	if response.Code != http.StatusOK {
		t.Error("response code should be 200")

	}
	if response.Body.String() != "pong" {
		t.Error("response body should be 'pong'")
	}
}
