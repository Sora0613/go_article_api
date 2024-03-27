package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gormConnect()

	s := gin.Default()

	companyRoutes := s.Group("/company")
	{
		companyRoutes.GET("", getAllCompanies)
		companyRoutes.GET("/:id", getCompany)
		companyRoutes.POST("", postCompany)
	}

	s.Run(":8080")
}
