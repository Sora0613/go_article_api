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

	obVisitRoutes := s.Group("/obvisit")
	{
		obVisitRoutes.GET("/:id", getOBVisits)
		obVisitRoutes.GET("", getAllOBVisits)
		obVisitRoutes.POST("", postOBVisits)
	}

	s.Run(":8080")
}
