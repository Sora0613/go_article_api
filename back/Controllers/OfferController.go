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
	offer := Models.Offer{}
	now := time.Now()
	offer.CreatedAt = now
	offer.UpdatedAt = now

	err := c.BindJSON(&offer)

	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	db.Create(&offer)
	db.Save(&offer)

	c.JSON(http.StatusOK, offer)
}
