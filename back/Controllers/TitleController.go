package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type TitleController struct {
	db *gorm.DB
}

func TitleDatabase(db *gorm.DB) *TitleController {
	return &TitleController{db: db}
}

func (tc TitleController) GetAllTitle(c *gin.Context) {
	var title []Models.Title
	tc.db.Find(&title)

	c.JSON(http.StatusOK, title)
}

func (tc TitleController) GetTitle(c *gin.Context) {
	id := c.Param("id")
	var title Models.Title

	// ArticleIDからタイトルを取得
	if err := tc.db.Where("article_id = ?", id).First(&title).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Title info not found for this article"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, title)
}

func (tc TitleController) PostTitle(c *gin.Context) {
	requestData := Models.Title{}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON: " + err.Error()})
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := tc.db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		article = Models.Article{
			Title: Models.Title{},
		}
		if err := tc.db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article: " + err.Error()})
			return
		}
	}

	// Articleに関連付けられたタイトルが既に存在するかを確認
	var existingTitle Models.Title
	if err := tc.db.Where("article_id = ?", requestData.ArticleID).First(&existingTitle).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Title already exists for this article."})
		return
	}

	title := Models.Title{
		ArticleID: article.ID,
		Title:     requestData.Title,
	}

	// タイトルをデータベースに作成
	if err := tc.db.Create(&title).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create title: " + err.Error()})
		return
	}

	// Articleにタイトルを関連付けて保存
	article.Title = title
	if err := tc.db.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save title to article: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, title)
}
