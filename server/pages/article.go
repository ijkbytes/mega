package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/service"
	"github.com/ijkbytes/mega/utils"
	"net/http"
	"strconv"
)

func Article(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if id <= 0 || err != nil {
		c.String(http.StatusNotFound, "article not found")
		return
	}

	article := service.Article.GetArticle(id)
	if article == nil {
		c.String(http.StatusNotFound, "article not found")
		return
	}

	article.ContentHTML = utils.MarkdownToHtml(article.ContentMD)

	c.HTML(http.StatusOK, "article.html", gin.H{
		"Title":   "Mega",
		"Article": article,
	})
}
