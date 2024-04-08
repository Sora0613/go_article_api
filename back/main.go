package main

import (
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Routes"
)

func main() {
	db := Database.GormConnect()
	s := gin.Default()

	Routes.SetupRoutes(s, db)

	s.Run(":8080")
}
