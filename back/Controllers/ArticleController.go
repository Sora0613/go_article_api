package Controllers

import (
	"github.com/gin-gonic/gin"
	"go_api/Database"
	"go_api/Models"
	"net/http"
)

// Articleには各項目のIDを入れるつもりなので必要な情報はArticleからアクセスできるようにしたい。

type ArticleController struct{}

func (ac ArticleController) GetAllArticle(c *gin.Context) {
	db := Database.GormConnect()
	var article []Models.Article
	db.Find(&article)

	c.JSON(http.StatusOK, article)
}
