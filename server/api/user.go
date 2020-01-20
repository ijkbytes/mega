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
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		UserName string `json:"userName" binding:"required,gte=4,lte=32"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=8,lte=32,printascii"`
		Confirm  string `json:"confirm" binding:"eqfield=Password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	// Only one user is currently allowed for the blog site
	// todo support for multiple users
	if count, err := service.User.UsersCount(); err != nil || count > 0 {
		logger.Warnf("only one user is currently allowed for the blog site")
		result.httpStatus = http.StatusForbidden
		result.Code = CodeUserExist
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
		result.httpStatus = http.StatusForbidden
		result.Code = CodeUserExist
		return
	}

	logger.Infof("create user `%s` success", newUser.UserName)
}

func Login(c *gin.Context) {
	logger := log.Get("Login").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		UserName string `json:"userName" binding:"required,gte=4,lte=32"`
		Password string `json:"password" binding:"required,gte=8,lte=32,printascii"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	user := service.User.GetUserByName(req.UserName)
	if user == nil {
		logger.Warnf("user `%s` not found: %v", req.UserName)
		result.httpStatus = http.StatusForbidden
		result.Code = CodeUserLoginError
		return
	}

	encrypted := utils.EncryptPassword(req.Password, user.Salt)
	if strings.Compare(encrypted, user.Password) != 0 {
		logger.Warnf("user's password error")
		result.httpStatus = http.StatusForbidden
		result.Code = CodeUserLoginError
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
		result.httpStatus = http.StatusInternalServerError
		result.Code = CodeError
		return
	}
}

func Logout(c *gin.Context) {
	logger := log.Get("Logout").Sugar()
	result := &Result{}
	defer returnResult(c, result)

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
			result.httpStatus = http.StatusInternalServerError
			result.Code = CodeError
			return
		}

		logger.Infof("`%s` logout success", userName)
	}
}
