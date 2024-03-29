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

	requestData := Models.Company{}

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
			Company: Models.Company{},
		}
		article.CreatedAt = now
		article.UpdatedAt = now
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	// Articleに関連付けられたCompanyが存在するかを確認
	var existingCompany Models.Company
	if err := db.Where("article_id = ?", requestData.ArticleID).First(&existingCompany).Error; err == nil {
		c.JSON(http.StatusConflict, "Company already exists for this article.")
		return
	}

	// 会社情報を登録
	now := time.Now()
	company := Models.Company{
		ArticleID:         article.ID,
		Name:              requestData.Name,
		Department:        requestData.Department,
		RecruitmentPeriod: requestData.RecruitmentPeriod,
	}
	company.CreatedAt = now
	company.UpdatedAt = now

	db.Create(&company)

	article.Company = company
	db.Save(&article)

	c.JSON(http.StatusOK, company)
}
