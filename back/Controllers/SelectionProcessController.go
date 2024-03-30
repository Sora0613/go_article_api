package Controllers

import (
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"net/http"
)

type SelectionProcessController struct{}

func (sc SelectionProcessController) GetAllSelectionProcesses(c *gin.Context) {
	db := Database.GormConnect()
	var selectionProcesses []Models.SelectionProcess
	db.Preload("Steps").Preload("EntrySheet").Preload("JobFair").Preload("WrittenTest").Preload("GroupDiscussion").Preload("OtherSelection").Preload("Interviews").Find(&selectionProcesses)

	c.JSON(http.StatusOK, selectionProcesses)
}

func (sc SelectionProcessController) GetSelectionProcess(c *gin.Context) {
	db := Database.GormConnect()
	articleID := c.Param("id")

	var selectionProcess Models.SelectionProcess

	if err := db.Preload("Steps").Preload("EntrySheet").Preload("JobFair").Preload("WrittenTest").Preload("GroupDiscussion").Preload("OtherSelection").Preload("Interviews").Where("article_id = ?", articleID).First(&selectionProcess).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SelectionProcess not found for this article"})
		return
	}

	c.JSON(http.StatusOK, selectionProcess)
}

func (sc SelectionProcessController) PostSelectionProcess(c *gin.Context) {
	db := Database.GormConnect()

	requestData := Models.SelectionProcess{}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON: " + err.Error()})
		return
	}

	// Articleを検索して存在しない場合は新規作成
	var article Models.Article
	if err := db.Where("id = ?", requestData.ArticleID).First(&article).Error; err != nil {
		// Articleが見つからない場合は新しいArticleを作成
		article = Models.Article{
			SelectionProcess: Models.SelectionProcess{},
		}
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article: " + err.Error()})
			return
		}
	}

	// SelectionProcessを作成
	selectionProcess := Models.SelectionProcess{
		ArticleID:       article.ID,
		Steps:           requestData.Steps,
		EntrySheet:      requestData.EntrySheet,
		JobFair:         requestData.JobFair,
		WrittenTest:     requestData.WrittenTest,
		GroupDiscussion: requestData.GroupDiscussion,
		OtherSelection:  requestData.OtherSelection,
		Interviews:      requestData.Interviews,
	}

	if err := db.Create(&selectionProcess).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create selectionProcess: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, selectionProcess)
}
