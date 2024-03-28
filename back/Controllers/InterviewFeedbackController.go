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

type InterviewFeedbackController struct{}

func (ic InterviewFeedbackController) GetAllInterviewFeedbacks(c *gin.Context) {
	db := Database.GormConnect()
	var interviewFeedbacks []Models.InterviewFeedback
	db.Find(&interviewFeedbacks)

	c.JSON(http.StatusOK, interviewFeedbacks)
}

func (ic InterviewFeedbackController) GetInterviewFeedback(c *gin.Context) {
	db := Database.GormConnect()
	id := c.Param("id")
	var interviewFeedback Models.InterviewFeedback

	// 面接フィードバックIDがない場合。
	if err := db.First(&interviewFeedback, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	db.First(&interviewFeedback, id)

	c.JSON(http.StatusOK, interviewFeedback)
}

func (ic InterviewFeedbackController) PostInterviewFeedback(c *gin.Context) {
	db := Database.GormConnect()
	interviewFeedback := Models.InterviewFeedback{}
	now := time.Now()
	interviewFeedback.CreatedAt = now
	interviewFeedback.UpdatedAt = now

	err := c.BindJSON(&interviewFeedback)

	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	db.Create(&interviewFeedback)
	db.Save(&interviewFeedback)

	c.JSON(http.StatusOK, interviewFeedback)
}
