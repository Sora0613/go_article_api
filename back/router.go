package main

import (
	"errors"
	"example/go/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Routes for company
func getAllCompanies(c *gin.Context) {
	db := gormConnect()
	var company []Models.Company
	db.Find(&company)

	c.JSON(200, company)
}

func getCompany(c *gin.Context) {
	db := gormConnect()
	id := c.Param("id")
	var company Models.Company

	// 会社がない場合。
	if err := db.First(&company, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func postCompany(c *gin.Context) {
	db := gormConnect()
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

	c.JSON(200, company)
}
