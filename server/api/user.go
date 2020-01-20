package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/model"
	"github.com/ijkbytes/mega/service"
	"github.com/ijkbytes/mega/utils"
	"net/http"
	"strings"
	"time"
)

func Register(c *gin.Context) {
	logger := log.Get("Register").Sugar()
	var req struct {
		UserName string `json:"user_name" binding:"required,gte=4,lte=32"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=8,lte=32,printascii"`
		Confirm  string `json:"confirm" binding:"eqfield=Password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	// Only one user is currently allowed for the blog site
	// todo support for multiple users
	if count, err := service.User.UsersCount(); err != nil || count > 0 {
		logger.Warnf("only one user is currently allowed for the blog site")
		c.JSON(http.StatusForbidden, gin.H{
			"code": CodeUserExist,
			"msg":  getErrMsg(CodeUserExist),
		})
		return
	}

	salt := utils.Salt()
	newUser := &model.User{
		UserName:  req.UserName,
		Email:     req.Email,
		AvatarUrl: "",
		Salt:      salt,
		Password:  utils.EncryptPassword(req.Password, salt),
	}
	if err := service.User.AddUser(newUser); err != nil {
		logger.Warnf("register user err: %v", err)
		c.JSON(http.StatusForbidden, gin.H{
			"code": CodeUserExist,
			"msg":  getErrMsg(CodeUserExist),
		})
		return
	}

	logger.Infof("create user `%s` success", newUser.UserName)
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}

func Login(c *gin.Context) {
	logger := log.Get("Login").Sugar()
	var req struct {
		UserName string `json:"user_name" binding:"required,gte=4,lte=32"`
		Password string `json:"password" binding:"required,gte=8,lte=32,printascii"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	user := service.User.GetUserByName(req.UserName)
	if user == nil {
		logger.Warnf("user `%s` not found: %v", req.UserName)
		c.JSON(http.StatusForbidden, gin.H{
			"code": CodeUserLoginError,
			"msg":  getErrMsg(CodeUserLoginError),
		})
		return
	}

	encrypted := utils.EncryptPassword(req.Password, user.Salt)
	if strings.Compare(encrypted, user.Password) != 0 {
		logger.Warnf("user's password error")
		c.JSON(http.StatusForbidden, gin.H{
			"code": CodeUserLoginError,
			"msg":  getErrMsg(CodeUserLoginError),
		})
		return
	}

	session := sessions.Default(c)
	ttl := config.Mega.Session.TTL
	session.Options(sessions.Options{
		MaxAge:   int(ttl),
		Path:     "/",
		HttpOnly: true,
	})
	session.Set("userName", user.UserName)
	session.Set("expireAt", time.Now().Add(time.Duration(ttl)*time.Second).Unix())
	err := session.Save()
	if err != nil {
		logger.Errorf("session save err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": CodeError,
			"msg":  getErrMsg(CodeError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}

func Logout(c *gin.Context) {
	logger := log.Get("Logout").Sugar()
	session := sessions.Default(c)
	userName, ok := session.Get("userName").(string)
	if !ok {
		logger.Warnf("session dose not exist")
	} else {
		session.Options(sessions.Options{
			MaxAge:   -1,
			Path:     "/",
			HttpOnly: true,
		})
		session.Clear()
		if err := session.Save(); err != nil {
			logger.Errorf("`%s` session save err: %v", userName, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": CodeError,
				"msg":  getErrMsg(CodeError),
			})
			return
		}

		logger.Infof("`%s` logout success", userName)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}
