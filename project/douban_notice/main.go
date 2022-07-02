package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"top.lel.dn/main/pkg/yaml"
)

func main() {

	r := gin.Default()
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

}
