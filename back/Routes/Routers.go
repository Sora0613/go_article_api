package Routes

import (
	"github.com/gin-gonic/gin"
	"go_api/Controllers"
)

func SetupRoutes(s *gin.Engine) {

	// インスタンス定義
	articleController := Controllers.ArticleController{}
	TitleController := Controllers.TitleController{}
	companyController := Controllers.CompanyController{}
	selectionProcessController := Controllers.SelectionProcessController{}
	obvisitController := Controllers.OBVisitController{}
	offerController := Controllers.OfferController{}
	interviewFeedbackController := Controllers.InterviewFeedbackController{}

	// TODO:IDは全てそれぞれのテーブルのID, ArticleIDでも撮ってこれるようにしようかな。
	articleRoutes := s.Group("/article")
	{
		// articleRoutes.GET("", articleController.GetAllArticle)
		articleRoutes.GET("/:id", articleController.GetArticle)
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
