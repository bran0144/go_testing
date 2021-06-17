package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pluralsight/webservice/go_testing/services"
)

type pingServiceMock struct {
	handlePingFn func() (string, error)
}

func (mock pingServiceMock) HandlePing() (string, error) {
	return mock.handlePingFn()
}

func TestPingWithError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.handlePingFn = func() (string, error) {
		return "", errors.New("error executing ping")
	}
	services.PingService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTextContext(response)

	Ping(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("response code should be 500")

	}
	if response.Body.String() != "error executing ping" {
		t.Error("error calling ping")
	}
}
func TestPingNoError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.handlePingFn = func() (string, error) {
		return "pong", nil
	}
	services.PingService = serviceMock

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
