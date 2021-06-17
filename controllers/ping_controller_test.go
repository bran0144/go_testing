package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pluralsight/webservice/go_testing/services"
)

type pingServiceMock struct {
}

func (mock pingServiceMock) HandlePing() (string, error) {
	return "pong", nil
}
func TestPing(t *testing.T) {
	services.PingService = pingServiceMock{}

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
