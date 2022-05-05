package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router = gin.Default()

func Run() {
	getRoutes()
	//gin.SetMode(gin.ReleaseMode)
	err := Router.Run(":8080")
	if err != nil {
		return
	}
}

func getRoutes() {

	v1 := Router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	// say hello...
	addHelloRoutes(v1)

	v2 := Router.Group("/v2")
	addUserRoutes(v2)
	addPingRoutes(v2)

	addHelloRoutes(v2)

	getFavicon()
}

func getFavicon() {
	Router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "okay")
	})

	Router.StaticFile("/favicon.ico", "./resources/favicon.ico")
}
