package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluralsight/webservice/go_testing/services"
)

func Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, result)
}
