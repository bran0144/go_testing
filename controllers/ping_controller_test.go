package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTextContext(response)

	fmt.Println(response.Code)
	Ping(context)
	fmt.Println(response.Code)

	if response.Code != http.StatusOK {
		t.Error("response code should be 200")

	}
}
