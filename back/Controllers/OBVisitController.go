package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type OBVisitController struct {
	db *gorm.DB
}

func OBVisitDatabase(db *gorm.DB) *OBVisitController {
	return &OBVisitController{db: db}
}

func (oc OBVisitController) GetAllOBVisits(c *gin.Context) {
	var obVisits []Models.OBVisits
	oc.db.Find(&obVisits)

	c.JSON(http.StatusOK, obVisits)
}

func (oc OBVisitController) GetOBVisits(c *gin.Context) {
	articleID := c.Param("id")
	var obVisits []Models.OBVisits

	// ArticleIDからOB訪問を検索
	if err := oc.db.Where("article_id = ?", articleID).Find(&obVisits).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "OBVisits not found for this article"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, obVisits)
}

func (oc OBVisitController) PostOBVisits(c *gin.Context) {
	requestData := Models.OBVisits{}

	if err := c.BindJSON(&requestData); err != nil {
		c.String(http.StatusBadRequest, "Failed to parse JSON: "+err.Error())
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := oc.db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		article = Models.Article{
			OBVisits: Models.OBVisits{},
		}
		if err := oc.db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	// Articleに関連付けられたOBVisitsが存在するかを確認
	var existingOBVisits Models.OBVisits
	if err := oc.db.Where("article_id = ?", requestData.ArticleID).First(&existingOBVisits).Error; err == nil {
		c.JSON(http.StatusConflict, "OBVisits already exists for this article.")
		return
	}

	// OBVisitsを登録
	obvisits := Models.OBVisits{
		ArticleID: article.ID,
		Visited:   requestData.Visited,
	}

	if err := oc.db.Create(&obvisits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create OBVisits: "+err.Error())
		return
	}

	// ArticleにOBVisitsを関連付けて保存
	article.OBVisits = obvisits
	if err := oc.db.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to save OBVisits to article: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, obvisits)
}
