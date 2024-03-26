package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gormConnect()

	s := gin.Default()

	s.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.Run(":8080")
}
