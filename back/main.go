package main

import (
	"github.com/gin-gonic/gin"
	"go_api/Routes"
)

func main() {
	// Database.GormConnect()
	s := gin.Default()

	Routes.SetupRoutes(s)

	s.Run(":8080")
}
