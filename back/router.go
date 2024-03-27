package main

import (
	"example/go/Models"
	"github.com/gin-gonic/gin"
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
	db.First(&company, id)

	c.JSON(200, company)
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
