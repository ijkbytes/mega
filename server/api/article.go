package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/service"
)

func GetArticle(c *gin.Context) {
	logger := log.Get("GetArticle").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		Id int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	article := service.Article.GetArticle(req.Id)
	if article == nil {
		logger.Warnf("cannot find id: %v", req.Id)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistArticle
		return
	}

	result.Data = article
}

func GetArticles(c *gin.Context) {
	logger := log.Get("GetArticles").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		State int `form:"state" binding:"oneof=-1 0 1"`
		TagId int `form:"tagId" binding:"gte=0"`
	}
	req.State = -1
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	if req.State != -1 {
		maps["state"] = req.State
	}
	if req.TagId != 0 {
		maps["tagId"] = req.TagId
	}

	data["lists"] = service.Article.GetArticles(0, 10, maps)
	data["total"] = service.Article.GetArticleTotal(maps)

	result.Data = data
}

func AddArticle(c *gin.Context) {
	logger := log.Get("AddArticle").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		TagID   int    `json:"tagId" binding:"required"`
		Title   string `json:"title" binding:"required"`
		Desc    string `json:"desc"`
		Content string `json:"content" binding:"required"`
		State   int    `json:"state" binding:"oneof=0 1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	if !service.Tag.ExistTagById(req.TagID) {
		logger.Warnf("cannot find tag id: %v", req.TagID)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistTag
		return
	}

	data := make(map[string]interface{})
	data["tagId"] = req.TagID
	data["title"] = req.Title
	data["desc"] = req.Desc
	data["contentMD"] = req.Content
	data["state"] = req.State

	service.Article.AddArticle(data)
}

func EditArticle(c *gin.Context) {
	logger := log.Get("EditArticle").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		TagID   int    `json:"tagId" binding:"required"`
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
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	if !service.Article.ExistArticleById(uriParams.Id) {
		logger.Warnf("cannot find article id: %v", uriParams.Id)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistArticle
		return
	}
	if !service.Tag.ExistTagById(req.TagID) {
		logger.Warnf("cannot find tag id: %v", uriParams.Id)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistTag
		return
	}

	data := make(map[string]interface{})
	data["tagId"] = req.TagID
	data["title"] = req.Title
	data["desc"] = req.Desc
	data["contentMD"] = req.Content
	service.Article.EditArticle(uriParams.Id, data)
}

func DeleteArticle(c *gin.Context) {
	logger := log.Get("DeleteArticle").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	if !service.Article.ExistArticleById(req.Id) {
		logger.Warnf("cannot find article id: %v", req.Id)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistArticle
		return
	}

	service.Article.DeleteArticle(req.Id)
}
