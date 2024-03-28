package Routes

import (
	"github.com/gin-gonic/gin"
	"go_api/Controllers"
)

func SetupRoutes(s *gin.Engine) {

	// インスタンス定義
	articleController := Controllers.ArticleController{}
	companyController := Controllers.CompanyController{}
	obvisitController := Controllers.OBVisitController{}
	offerController := Controllers.OfferController{}

	articleRoutes := s.Group("/article")
	{
		articleRoutes.GET("", articleController.GetAllArticle)
	}

	companyRoutes := s.Group("/company")
	{
		companyRoutes.GET("", companyController.GetAllCompanies)
		companyRoutes.GET("/:id", companyController.GetCompany)
		companyRoutes.POST("", companyController.PostCompany)
	}

	obVisitRoutes := s.Group("/obvisit")
	{
		obVisitRoutes.GET("", obvisitController.GetAllOBVisits)
		obVisitRoutes.GET("/:id", obvisitController.GetOBVisits)
		obVisitRoutes.POST("", obvisitController.PostOBVisits)
	}

	offerRoutes := s.Group("/offer")
	{
		offerRoutes.GET("", offerController.GetAllOffer)
		offerRoutes.GET("/:id", offerController.GetOffer)
		offerRoutes.POST("", offerController.PostOffer)
	}

}
