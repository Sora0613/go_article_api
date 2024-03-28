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

	c.JSON(http.StatusOK, company)
}

func getCompany(c *gin.Context) {
	db := gormConnect()
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

	c.JSON(http.StatusOK, company)
}

// Routes for OBVisits
func getAllOBVisits(c *gin.Context) {
	db := gormConnect()
	var obVisits []Models.OBVisits
	db.Find(&obVisits)

	c.JSON(http.StatusOK, obVisits)
}

func getOBVisits(c *gin.Context) {
	db := gormConnect()
	id := c.Param("id")
	var obVisits []Models.OBVisits

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

func postOBVisits(c *gin.Context) {
	db := gormConnect()
	obVisits := Models.OBVisits{}
	now := time.Now()
	obVisits.CreatedAt = now
	obVisits.CreatedAt = now

	err := c.BindJSON(&obVisits)

	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	db.Create(&obVisits)
	db.Save(&obVisits)

	c.JSON(http.StatusOK, obVisits)
}
