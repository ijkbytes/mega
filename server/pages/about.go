package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}
