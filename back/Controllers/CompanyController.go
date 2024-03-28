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

type CompanyController struct{}

func (cc CompanyController) GetAllCompanies(c *gin.Context) {
	db := Database.GormConnect()
	var company []Models.Company
	db.Find(&company)

	c.JSON(http.StatusOK, company)
}

func (cc CompanyController) GetCompany(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")
	var company Models.Company

	// 会社IDがない場合。
	if err := db.First(&company, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (cc CompanyController) PostCompany(c *gin.Context) {
	db := Database.GormConnect()
	company := Models.Company{}
	now := time.Now()
	company.CreatedAt = now
	company.CreatedAt = now

	err := c.BindJSON(&company)

	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	db.Create(&company)
	db.Save(&company)

	c.JSON(http.StatusOK, company)
}
