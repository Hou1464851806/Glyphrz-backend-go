package main

import (
	"Glyphrz-go/global"
	"Glyphrz-go/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitConfig()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}
