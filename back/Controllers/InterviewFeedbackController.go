package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type InterviewFeedbackController struct{}

func (ic InterviewFeedbackController) GetAllInterviewFeedbacks(c *gin.Context) {
	db := Database.GormConnect()
	var interviewFeedbacks []Models.InterviewFeedback
	db.Find(&interviewFeedbacks)

	c.JSON(http.StatusOK, interviewFeedbacks)
}

func (ic InterviewFeedbackController) GetInterviewFeedback(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")
	var interviewFeedback Models.InterviewFeedback

	// ArticleIDから取得
	if err := db.Where("article_id = ?", id).First(&interviewFeedback).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, interviewFeedback)
}

func (ic InterviewFeedbackController) PostInterviewFeedback(c *gin.Context) {
	db := Database.GormConnect()

	requestData := Models.InterviewFeedback{}

	if err := c.BindJSON(&requestData); err != nil {
		c.String(http.StatusBadRequest, "Failed to parse JSON: "+err.Error())
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		article = Models.Article{
			InterviewFeedback: Models.InterviewFeedback{},
		}
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	// Articleに関連付けられたInterviewFeedbackが存在するかを確認
	var existingInterviewFeedback Models.InterviewFeedback
	if err := db.Where("article_id = ?", requestData.ArticleID).First(&existingInterviewFeedback).Error; err == nil {
		c.JSON(http.StatusConflict, "InterviewFeedback already exists for this article.")
		return
	}

	// InterviewFeedbackを登録
	interviewFeedback := Models.InterviewFeedback{
		ArticleID:         article.ID,
		MainFocus:         requestData.MainFocus,
		MemorableQuestion: requestData.MemorableQuestion,
		Preparation:       requestData.Preparation,
		KeyPoints:         requestData.KeyPoints,
		ResearchSource:    requestData.ResearchSource,
		Impressions:       requestData.Impressions,
		AdviceForFuture:   requestData.AdviceForFuture,
	}

	if err := db.Create(&interviewFeedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create interview feedback: "+err.Error())
		return
	}

	// ArticleにInterviewFeedbackを関連付けて保存
	article.InterviewFeedback = interviewFeedback
	if err := db.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to save interview feedback to article: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, interviewFeedback)
}
