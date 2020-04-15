package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}
