package Routes

import (
	"github.com/gin-gonic/gin"
	"go_api/Controllers"
	"gorm.io/gorm"
)

func SetupRoutes(s *gin.Engine, db *gorm.DB) {

	// インスタンス定義
	articleController := Controllers.ArticleDatabase(db)
	TitleController := Controllers.TitleDatabase(db)
	companyController := Controllers.CompanyDatabase(db)
	selectionProcessController := Controllers.SelectionProcessDatabase(db)
	obvisitController := Controllers.OBVisitDatabase(db)
	offerController := Controllers.OfferDatabase(db)
	interviewFeedbackController := Controllers.InterviewFeedbackDatabase(db)

	// ID指定の場合は全てArticleのIDを指定

	articleRoutes := s.Group("/article")
	{
		articleRoutes.GET("", articleController.GetAllArticle)
		articleRoutes.GET("/:id", articleController.GetArticle)
		articleRoutes.POST("", articleController.PostArticle)
	}

	titleRoutes := s.Group("/title")
	{
		titleRoutes.GET("", TitleController.GetAllTitle)
		titleRoutes.GET("/:id", TitleController.GetTitle)
		titleRoutes.POST("", TitleController.PostTitle)
	}

	companyRoutes := s.Group("/company")
	{
		companyRoutes.GET("", companyController.GetAllCompanies)
		companyRoutes.GET("/:id", companyController.GetCompany)
		companyRoutes.POST("", companyController.PostCompany)
	}

	selectionProcessRoutes := s.Group("/selection-process")
	{
		selectionProcessRoutes.GET("", selectionProcessController.GetAllSelectionProcesses)
		selectionProcessRoutes.GET("/:id", selectionProcessController.GetSelectionProcess)
		selectionProcessRoutes.POST("", selectionProcessController.PostSelectionProcess)
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

	interviewFeedbackRoutes := s.Group("/interview-feedback")
	{
		interviewFeedbackRoutes.GET("", interviewFeedbackController.GetAllInterviewFeedbacks)
		interviewFeedbackRoutes.GET("/:id", interviewFeedbackController.GetInterviewFeedback)
		interviewFeedbackRoutes.POST("", interviewFeedbackController.PostInterviewFeedback)
	}

}
