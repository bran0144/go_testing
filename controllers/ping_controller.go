package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluralsight/webservice/go_testing/services"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, services.HandlePing())
}
