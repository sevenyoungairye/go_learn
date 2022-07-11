// gin app, pls see: https://pkg.go.dev/github.com/gin-gonic/gin#readme-quick-start

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"top.lel.dn/main/internal/corn"
	"top.lel.dn/main/internal/services/episode"
	"top.lel.dn/main/pkg/yaml"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	port := fmt.Sprintf(":%d", yaml.GetServer().Server.Port)
	err := r.Run(port)
	if err != nil {
		return
	}

	corn.Demo()
	episode.Run()
}
