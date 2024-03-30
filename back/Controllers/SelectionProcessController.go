package Controllers

import (
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"net/http"
)

type SelectionProcessController struct{}

func (sc SelectionProcessController) GetSelectionProcess(c *gin.Context) {
	db := Database.GormConnect()
	article_id := c.Params.ByName("id")
	selectionProcess := Models.SelectionProcess{
		Steps:           Models.Steps{},
		EntrySheet:      Models.EntrySheet{},
		JobFair:         Models.JobFair{},
		WrittenTest:     Models.WrittenTest{},
		GroupDiscussion: Models.GroupDiscussion{},
		OtherSelection:  Models.OtherSelection{},
		Interviews:      []Models.Interview{},
	}

	if err := db.Where("article_id = ?", article_id).First(&selectionProcess).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}

	// get selectionProcess's id
	selectionProcessID := selectionProcess.ID

	// get steps by selectionProcess's id
	var steps Models.Steps
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&steps).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}
	// get related tables by selectionProcess's id

	// ここから
	// get entrySheet by selectionProcess's id
	var entrySheet Models.EntrySheet
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&entrySheet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}

	// get jobFair by selectionProcess's id
	var jobFair Models.JobFair
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&jobFair).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}

	// get writtenTest by selectionProcess's id
	var writtenTest Models.WrittenTest
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&writtenTest).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}

	// get groupDiscussion by selectionProcess's id
	var groupDiscussion Models.GroupDiscussion
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&groupDiscussion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}

	// get otherSelection by selectionProcess's id
	var otherSelection Models.OtherSelection
	if err := db.Where("selection_process_id = ?", selectionProcessID).First(&otherSelection).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}

	// get interviews by selectionProcess's id and interviewType. Then, get interviews interviewType times.
	var interviews []Models.Interview
	if err := db.Where("selection_process_id = ?", selectionProcessID).Find(&interviews).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This id not found"})
		return
	}
	// ここまでnullableにしたい。

	selectionProcess.Steps = steps
	selectionProcess.EntrySheet = entrySheet
	selectionProcess.JobFair = jobFair
	selectionProcess.WrittenTest = writtenTest
	selectionProcess.GroupDiscussion = groupDiscussion
	selectionProcess.OtherSelection = otherSelection
	selectionProcess.Interviews = interviews

	c.JSON(http.StatusOK, selectionProcess)
}

func (sc SelectionProcessController) PostSelectionProcess(c *gin.Context) {
	db := Database.GormConnect()

	requestData := Models.SelectionProcess{}

	if err := c.BindJSON(&requestData); err != nil {
		c.String(http.StatusBadRequest, "Failed to parse JSON: "+err.Error())
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
			c.JSON(http.StatusInternalServerError, "Failed to create article: "+err.Error())
			return
		}
	}

	// SelectionProcessを作成
	selectionProcess := Models.SelectionProcess{
		ArticleID:       requestData.ArticleID,
		Steps:           requestData.Steps,
		EntrySheet:      requestData.EntrySheet,
		JobFair:         requestData.JobFair,
		WrittenTest:     requestData.WrittenTest,
		GroupDiscussion: requestData.GroupDiscussion,
		OtherSelection:  requestData.OtherSelection,
		Interviews:      requestData.Interviews,
	}
	if err := db.Create(&selectionProcess).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create selectionProcess: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, selectionProcess)
}
