package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type ArticleController struct{}

func (ac ArticleController) GetAllArticle(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")

	var article Models.Article
	article = Models.Article{
		Title:             Models.Title{},
		Company:           Models.Company{},
		OBVisits:          Models.OBVisits{},
		Offer:             Models.Offer{},
		InterviewFeedback: Models.InterviewFeedback{},
	}

	if err := db.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// TODO: all of errors return json

	var title Models.Title
	if err := db.Where("article_id = ?", id).First(&title).Error; err != nil {
		c.String(http.StatusNotFound, "Title not found")
		return
	}

	var company Models.Company
	if err := db.Where("article_id = ?", id).First(&company).Error; err != nil {
		c.String(http.StatusNotFound, "Company not found")
		return
	}

	var OBVisits Models.OBVisits
	if err := db.Where("article_id = ?", id).First(&OBVisits).Error; err != nil {
		c.String(http.StatusNotFound, "OBVisits not found")
		return
	}

	var Offer Models.Offer
	if err := db.Where("article_id = ?", id).First(&Offer).Error; err != nil {
		c.String(http.StatusNotFound, "Offer not found")
		return
	}

	var InterviewFeedback Models.InterviewFeedback
	if err := db.Where("article_id = ?", id).First(&InterviewFeedback).Error; err != nil {
		c.String(http.StatusNotFound, "InterviewFeedback not found")
		return
	}

	article.Title = title
	article.Company = company
	article.OBVisits = OBVisits
	article.Offer = Offer
	article.InterviewFeedback = InterviewFeedback

	c.JSON(http.StatusOK, article)
}
