package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/service"
	"github.com/ijkbytes/mega/utils"
	"net/http"
)

func Index(c *gin.Context) {
	page := utils.GetPage(c)
	articles := service.Article.GetArticles(page-1, 1, make(map[string]interface{}))
	count := service.Article.GetArticleTotal(make(map[string]interface{}))
	pagination := utils.NewPagination(page, 1, count)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Articles":   articles,
		"Url":        "/",
		"Pagination": pagination,
	})
}
