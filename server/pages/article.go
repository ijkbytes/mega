package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/service"
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

	c.HTML(http.StatusOK, "article.html", gin.H{
		"Title":   "Mega",
		"Article": article,
	})
}
