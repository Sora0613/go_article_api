package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ArticleController struct{}

func (ac ArticleController) GetAllArticle(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10 // デフォルトのページサイズを10に。
	}

	// データベースへの接続
	db := Database.GormConnect()

	var articles []Models.Article

	offset := (page - 1) * pageSize
	db.Preload("Title").
		Preload("Company").
		Preload("SelectionProcess").
		Preload("OBVisits").
		Preload("Offer").
		Preload("InterviewFeedback").
		Offset(offset).
		Limit(pageSize).
		Find(&articles)

	// forで回してarticleに各IDを設定
	for i := range articles {
		articles[i].Title.ArticleID = articles[i].ID
		articles[i].Company.ArticleID = articles[i].ID
		articles[i].SelectionProcess.ArticleID = articles[i].ID
		articles[i].OBVisits.ArticleID = articles[i].ID
		articles[i].Offer.ArticleID = articles[i].ID
		articles[i].InterviewFeedback.ArticleID = articles[i].ID
	}

	c.JSON(http.StatusOK, articles)
}

func (ac ArticleController) GetArticle(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")

	var article Models.Article
	if err := db.Preload("Title").
		Preload("Company").
		Preload("SelectionProcess").
		Preload("SelectionProcess.Steps").
		Preload("SelectionProcess.EntrySheet").
		Preload("SelectionProcess.JobFair").
		Preload("SelectionProcess.WrittenTest").
		Preload("SelectionProcess.GroupDiscussion").
		Preload("SelectionProcess.OtherSelection").
		Preload("SelectionProcess.Interviews").
		Preload("OBVisits").
		Preload("Offer").
		Preload("InterviewFeedback").
		First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Article with this id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// 必要な情報がArticle内に含まれているため、返すだけでOK
	c.JSON(http.StatusOK, article)
}

func (ac ArticleController) PostArticle(c *gin.Context) {
	db := Database.GormConnect()
	tx := db.Begin()

	var requestData Models.Article
	if err := c.BindJSON(&requestData); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON: " + err.Error()})
		return
	}

	article := Models.Article{
		Title:             requestData.Title,
		Company:           requestData.Company,
		SelectionProcess:  requestData.SelectionProcess,
		OBVisits:          requestData.OBVisits,
		Offer:             requestData.Offer,
		InterviewFeedback: requestData.InterviewFeedback,
	}

	// articleに関連して紐付け。
	article.Title.ArticleID = article.ID
	article.Company.ArticleID = article.ID
	article.SelectionProcess.ArticleID = article.ID
	article.OBVisits.ArticleID = article.ID
	article.Offer.ArticleID = article.ID
	article.InterviewFeedback.ArticleID = article.ID

	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article: " + err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully"})
}
