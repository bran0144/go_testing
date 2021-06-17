package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pluralsight/webservice/go_testing/controllers"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.Ping)
	router.RUN(":8080")
}
