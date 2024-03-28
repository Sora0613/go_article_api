package Routes

import (
	"github.com/gin-gonic/gin"
	"go_api/Controllers"
)

func SetupRoutes(s *gin.Engine) {

	companyRoutes := s.Group("/company")
	{
		companyRoutes.GET("", Controllers.GetAllCompanies)
		companyRoutes.GET("/:id", Controllers.GetCompany)
		companyRoutes.POST("", Controllers.PostCompany)
	}

	obVisitRoutes := s.Group("/obvisit")
	{
		obVisitRoutes.GET("/:id", Controllers.GetOBVisits)
		obVisitRoutes.GET("", Controllers.GetAllOBVisits)
		obVisitRoutes.POST("", Controllers.PostOBVisits)
	}
}
