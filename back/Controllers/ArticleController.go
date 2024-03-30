package Controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type ArticleController struct{}

func (ac ArticleController) GetAllArticle(c *gin.Context) {
	db := Database.GormConnect()
	var article []Models.Article
	db.Find(&article)

	c.JSON(http.StatusOK, article)
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
