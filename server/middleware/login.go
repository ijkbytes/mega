package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NeedLogin(c *gin.Context) {
	session := sessions.Default(c)
	userName, _ := session.Get("userName").(string)
	expireAt, _ := session.Get("expireAt").(int64)

	if len(userName) > 0 && time.Unix(expireAt, 0).After(time.Now()) {
		c.Next()
		return
	}

	session.Options(sessions.Options{
		MaxAge:   -1, // delete cookie
		Path:     "/",
		HttpOnly: true,
	})
	session.Clear()
	_ = session.Save()

	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"code": http.StatusForbidden,
		"msg":  "no login",
	})
}
