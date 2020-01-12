package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/service"
	"net/http"
)

func GetArticle(c *gin.Context) {
	logger := log.Get("GetArticle").Sugar()
	var req struct {
		Id int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	article := service.Article.GetArticle(req.Id)
	if article == nil {
		logger.Warnf("cannot find id: %v", req.Id)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistArticle,
			"msg":  getErrMsg(CodeNotExistArticle),
			"data": article,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": article,
	})
}

func GetArticles(c *gin.Context) {
	logger := log.Get("GetArticles").Sugar()
	var req struct {
		State int `form:"state" binding:"oneof=-1 0 1"`
		TagId int `form:"tag_id" binding:"gte=0"`
	}
	req.State = -1
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	if req.State != -1 {
		maps["state"] = req.State
	}
	if req.TagId != 0 {
		maps["tag_id"] = req.TagId
	}

	data["lists"] = service.Article.GetArticles(0, 10, maps)
	data["total"] = service.Article.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	logger := log.Get("AddArticle").Sugar()
	var req struct {
		TagID   int    `json:"tag_id" binding:"required"`
		Title   string `json:"title" binding:"required"`
		Desc    string `json:"desc"`
		Content string `json:"content" binding:"required"`
		State   int    `json:"state" binding:"oneof=0 1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	if !service.Tag.ExistTagById(req.TagID) {
		logger.Warnf("cannot find tag id: %v", req.TagID)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistTag,
			"msg":  getErrMsg(CodeNotExistTag),
		})
		return
	}

	data := make(map[string]interface{})
	data["tag_id"] = req.TagID
	data["title"] = req.Title
	data["desc"] = req.Desc
	data["content_md"] = req.Content
	data["state"] = req.State

	service.Article.AddArticle(data)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": make(map[string]interface{}),
	})
}

func EditArticle(c *gin.Context) {
	logger := log.Get("EditArticle").Sugar()
	var req struct {
		TagID   int    `json:"tag_id" binding:"required"`
		Title   string `json:"title" binding:"required"`
		Desc    string `json:"desc"`
		Content string `json:"content" binding:"required"`
		State   int    `json:"state" binding:"oneof=0 1"`
	}
	var uriParams struct {
		Id int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
		return
	}

	if !service.Article.ExistArticleById(uriParams.Id) {
		logger.Warnf("cannot find article id: %v", uriParams.Id)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistArticle,
			"msg":  getErrMsg(CodeNotExistArticle),
		})
		return
	}
	if !service.Tag.ExistTagById(req.TagID) {
		logger.Warnf("cannot find tag id: %v", uriParams.Id)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistTag,
			"msg":  getErrMsg(CodeNotExistTag),
		})
		return
	}

	data := make(map[string]interface{})
	data["tag_id"] = req.TagID
	data["title"] = req.Title
	data["desc"] = req.Desc
	data["content_md"] = req.Content
	service.Article.EditArticle(uriParams.Id, data)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": make(map[string]string),
	})
}

func DeleteArticle(c *gin.Context) {
	logger := log.Get("DeleteArticle").Sugar()
	var req struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	if !service.Article.ExistArticleById(req.Id) {
		logger.Warnf("cannot find article id: %v", req.Id)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistArticle,
			"msg":  getErrMsg(CodeNotExistArticle),
		})
		return
	}

	service.Article.DeleteArticle(req.Id)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": make(map[string]string),
	})
}
