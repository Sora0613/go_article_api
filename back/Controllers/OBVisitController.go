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
	obVisits := Models.OBVisits{}
	now := time.Now()
	obVisits.CreatedAt = now
	obVisits.UpdatedAt = now

	err := c.BindJSON(&obVisits)

	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	db.Create(&obVisits)
	db.Save(&obVisits)

	c.JSON(http.StatusOK, obVisits)
}
