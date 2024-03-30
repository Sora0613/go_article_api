package Controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"gorm.io/gorm"
	"net/http"
)

type ArticleController struct{}

func (ac ArticleController) GetArticle(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")

	// Articleを取得
	var article Models.Article
	if err := db.Preload("Title").Preload("Company").Preload("SelectionProcess").Preload("OBVisits").Preload("Offer").Preload("InterviewFeedback").First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Article with this id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// SelectionProcessIDを取得
	selectionProcessID := article.SelectionProcess.ID
	fmt.Println("selectionProcessID: ", selectionProcessID)

	// selectionProcessIDからStepsを取得
	var steps Models.Steps
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&steps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Steps not found for this article"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// selectionProcessIDからエントリーシートを取得
	var entrySheet Models.EntrySheet
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&entrySheet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EntrySheet not found for this article"})
		return
	}

	// 説明会を取得
	var jobFair Models.JobFair
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&jobFair).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "JobFair not found for this article"})
		return
	}

	// selectionProcessIDから試験の情報を取得
	var writtenTest Models.WrittenTest
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&writtenTest).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "WrittenTest not found for this article"})
		return
	}

	// selectionProcessIDからグループディスカッションを取得
	var groupDiscussion Models.GroupDiscussion
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&groupDiscussion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "GroupDiscussion not found for this article"})
		return
	}

	// selectionProcessIDからその他の選考を取得
	var otherSelection Models.OtherSelection
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&otherSelection).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OtherSelection not found for this article"})
		return
	}

	// selectionProcessIDから面接情報を取得
	var interviews []Models.Interview
	if err := db.Where("selection_process_id = ?", selectionProcessID).Find(&interviews).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Interviews not found for this article"})
		return
	}

	// 必要な情報を記事に追加
	article.SelectionProcess.Steps = steps
	article.SelectionProcess.EntrySheet = entrySheet
	article.SelectionProcess.JobFair = jobFair
	article.SelectionProcess.WrittenTest = writtenTest
	article.SelectionProcess.GroupDiscussion = groupDiscussion
	article.SelectionProcess.OtherSelection = otherSelection
	article.SelectionProcess.Interviews = interviews

	// selectionProcessIDからOB訪問を取得
	var OBVisits Models.OBVisits
	if err := db.Where("article_id = ?", id).First(&OBVisits).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OBVisits not found for this article"})
		return
	}

	// selectionProcessIDから内定情報を取得
	var Offer Models.Offer
	if err := db.Where("article_id = ?", id).First(&Offer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found for this article"})
		return
	}

	// selectionProcessIDから面接についての感想を取得
	var InterviewFeedback Models.InterviewFeedback
	if err := db.Where("article_id = ?", id).First(&InterviewFeedback).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "InterviewFeedback not found for this article"})
		return
	}

	article.OBVisits = OBVisits
	article.Offer = Offer
	article.InterviewFeedback = InterviewFeedback

	c.JSON(http.StatusOK, article)
}
