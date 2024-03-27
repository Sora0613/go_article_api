package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gormConnect()

	s := gin.Default()

	s.GET("/company", getAllCompanies)

	s.GET("/company/:id", getCompany)

	s.POST("/company", postCompany)

	s.Run(":8080")
}
