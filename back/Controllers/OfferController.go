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

type OfferController struct{}

func (oc OfferController) GetAllOffer(c *gin.Context) {
	db := Database.GormConnect()
	var offer []Models.Offer
	db.Find(&offer)

	c.JSON(http.StatusOK, offer)
}

func (oc OfferController) GetOffer(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Params.ByName("id")
	var offer Models.Offer
	db.First(&offer, id)

	if err := db.First(&offer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, offer)
}

func (oc OfferController) PostOffer(c *gin.Context) {
	db := Database.GormConnect()

	requestData := Models.Offer{}

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
			Offer: Models.Offer{},
		}
		article.CreatedAt = now
		article.UpdatedAt = now
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	var existingOffer Models.Offer
	if err := db.Where("article_id = ?", requestData.ArticleID).First(&existingOffer).Error; err == nil {
		c.JSON(http.StatusConflict, "Offer already exists for this article.")
		return
	}

	now := time.Now()
	offer := Models.Offer{
		ArticleID:      article.ID,
		Offered:        requestData.Offered,
		OfferedAt:      requestData.OfferedAt,
		TaskAfterOffer: requestData.TaskAfterOffer,
		Constraints:    requestData.Constraints,
	}
	offer.CreatedAt = now
	offer.UpdatedAt = now

	db.Create(&offer)

	article.Offer = offer
	db.Save(&article)

	c.JSON(http.StatusOK, offer)
}
