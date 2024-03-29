package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type TitleController struct{}

func (oc TitleController) GetAllTitle(c *gin.Context) {
	db := Database.GormConnect()
	var title []Models.Title
	db.Find(&title)

	c.JSON(http.StatusOK, title)
}

func (oc TitleController) GetTitle(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Params.ByName("id")
	var title Models.Title
	db.First(&title, id)

	if err := db.First(&title, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, title)
}

func (oc TitleController) PostTitle(c *gin.Context) {
	db := Database.GormConnect()

	requestData := Models.Title{}

	if err := c.BindJSON(&requestData); err != nil {
		c.String(http.StatusBadRequest, "Failed to parse JSON: "+err.Error())
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		now := time.Now()
		article = Models.Article{
			Title: Models.Title{},
		}
		article.CreatedAt = now
		article.UpdatedAt = now
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	var existingTitle Models.Title
	if err := db.Where("article_id = ?", requestData.ArticleID).First(&existingTitle).Error; err == nil {
		c.JSON(http.StatusConflict, "Title already exists for this article.")
		return
	}

	now := time.Now()
	title := Models.Title{
		ArticleID: article.ID,
		Title:     requestData.Title,
	}
	title.CreatedAt = now
	title.UpdatedAt = now

	db.Create(&title)

	article.Title = title
	db.Save(&article)

	c.JSON(http.StatusOK, title)
}
