package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type OfferController struct {
	db *gorm.DB
}

func OfferDatabase(db *gorm.DB) *OfferController {
	return &OfferController{db: db}
}

func (oc OfferController) GetAllOffer(c *gin.Context) {
	var offer []Models.Offer
	oc.db.Find(&offer)

	c.JSON(http.StatusOK, offer)
}

func (oc OfferController) GetOffer(c *gin.Context) {
	id := c.Params.ByName("id")
	var offer Models.Offer
	oc.db.First(&offer, id)

	// ArticleIDから取得
	if err := oc.db.Where("article_id = ?", id).First(&offer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Offer info not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, offer)
}

func (oc OfferController) PostOffer(c *gin.Context) {
	requestData := Models.Offer{}

	if err := c.BindJSON(&requestData); err != nil {
		c.String(http.StatusBadRequest, "Failed to parse JSON: "+err.Error())
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := oc.db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		article = Models.Article{
			Offer: Models.Offer{},
		}
		if err := oc.db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	// Articleに関連付けられたOfferが存在するかを確認
	var existingOffer Models.Offer
	if err := oc.db.Where("article_id = ?", requestData.ArticleID).First(&existingOffer).Error; err == nil {
		c.JSON(http.StatusConflict, "Offer already exists for this article.")
		return
	}

	// 内定情報の登録
	offer := Models.Offer{
		ArticleID:      article.ID,
		Offered:        requestData.Offered,
		OfferedAt:      requestData.OfferedAt,
		TaskAfterOffer: requestData.TaskAfterOffer,
		Constraints:    requestData.Constraints,
	}

	if err := oc.db.Create(&offer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create Offer: "+err.Error())
		return
	}

	// ArticleにOfferを関連付けて保存
	article.Offer = offer
	if err := oc.db.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to save Offer to article: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, offer)
}
