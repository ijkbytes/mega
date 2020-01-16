package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	logger := log.Get("Login").Sugar()
	session := sessions.Default(c)

	// todo check user info

	ttl := config.Mega.Session.TTL
	session.Options(sessions.Options{
		MaxAge: int(ttl),
		Path:   "/",
	})

	session.Set("userName", "ijkbytes")
	session.Set("expireAt", time.Now().Add(time.Duration(ttl)*time.Second).Unix())
	err := session.Save()
	if err != nil {
		logger.Errorf("err: %v", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}
