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

type OBVisitController struct{}

func (oc OBVisitController) GetAllOBVisits(c *gin.Context) {
	db := Database.GormConnect()
	var obVisits []Models.OBVisits
	db.Find(&obVisits)

	c.JSON(http.StatusOK, obVisits)
}

func (oc OBVisitController) GetOBVisits(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")
	var obVisits []Models.OBVisits

	// OB訪問IDがない場合。
	if err := db.First(&obVisits, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, obVisits)
}

func (oc OBVisitController) PostOBVisits(c *gin.Context) {
	db := Database.GormConnect()

	requestData := Models.OBVisits{}

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
			OBVisits: Models.OBVisits{},
		}
		article.CreatedAt = now
		article.UpdatedAt = now
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	var existingOBVisits Models.OBVisits
	if err := db.Where("article_id = ?", requestData.ArticleID).First(&existingOBVisits).Error; err == nil {
		c.JSON(http.StatusConflict, "OBVisits already exists for this article.")
		return
	}

	// 会社情報を登録
	now := time.Now()
	obvisits := Models.OBVisits{
		ArticleID: article.ID,
		Visited:   requestData.Visited,
	}
	obvisits.CreatedAt = now
	obvisits.UpdatedAt = now

	db.Create(&obvisits)

	article.OBVisits = obvisits
	db.Save(&article)

	c.JSON(http.StatusOK, obvisits)
}
