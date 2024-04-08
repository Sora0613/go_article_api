package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CompanyController struct {
	db *gorm.DB
}

func CompanyDatabase(db *gorm.DB) *CompanyController {
	return &CompanyController{db: db}
}

func (cc CompanyController) GetAllCompanies(c *gin.Context) {
	var company []Models.Company
	cc.db.Find(&company)

	c.JSON(http.StatusOK, company)
}

func (cc CompanyController) GetCompany(c *gin.Context) {
	id := c.Param("id")
	var company Models.Company

	// ArticleIDから関連付けられたCompanyを取得
	if err := cc.db.Where("article_id = ?", id).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company info not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (cc CompanyController) PostCompany(c *gin.Context) {
	requestData := Models.Company{}

	if err := c.BindJSON(&requestData); err != nil {
		c.String(http.StatusBadRequest, "Failed to parse JSON: "+err.Error())
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := cc.db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		now := time.Now()
		article = Models.Article{
			Company: Models.Company{},
		}
		article.CreatedAt = now
		article.UpdatedAt = now
		if err := cc.db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article: " + err.Error()})
			return
		}
	}

	// Articleに関連付けられたCompanyが存在するかを確認
	var existingCompany Models.Company
	if err := cc.db.Where("article_id = ?", requestData.ArticleID).First(&existingCompany).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Company already exists for this article"})
		return
	}

	// 会社情報を登録
	company := Models.Company{
		ArticleID:         article.ID,
		Name:              requestData.Name,
		Department:        requestData.Department,
		RecruitmentPeriod: requestData.RecruitmentPeriod,
	}

	if err := cc.db.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: " + err.Error()})
		return
	}

	// ArticleにCompanyを関連付けて保存
	article.Company = company
	if err := cc.db.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save company to article: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}
