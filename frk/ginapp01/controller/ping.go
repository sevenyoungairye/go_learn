package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addPingRoutes(rg *gin.RouterGroup) {
	// ping链接组
	pingGroup := rg.Group("/ping")
	// ping下面的url
	pingGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
}
