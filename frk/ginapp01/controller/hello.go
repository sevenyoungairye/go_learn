package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addHelloRoutes(group *gin.RouterGroup) {
	group.GET("/hello", SayHello)
}

func SayHello(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  200,
		"message": "hell world!",
	})
}
