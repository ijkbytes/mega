package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/service"
)

func GetTags(c *gin.Context) {
	logger := log.Get("GetTags").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		Name  string `form:"name"`
		State int    `form:"state" binding:"oneof=-1 0 1"`
	}
	req.State = -1
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if req.Name != "" {
		maps["name"] = req.Name
	}
	if req.State != -1 {
		maps["state"] = req.State
	}

	data["lists"] = service.Tag.GetTags(0, 10, maps)
	data["total"] = service.Tag.GetTagTotal(maps)

	result.Data = data
}

func AddTag(c *gin.Context) {
	logger := log.Get("AddTag").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		Name  string `json:"name" binding:"required"`
		State int    `json:"state" binding:"oneof=0 1"`
	}
	if err := c.ShouldBind(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	service.Tag.AddTag(req.Name, req.State)
}

func EditTag(c *gin.Context) {
	logger := log.Get("EditTag").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	var req struct {
		Name  string `json:"name" binding:"required"`
		State int    `json:"state" binding:"oneof=0 1"`
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
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
		return
	}

	if !service.Tag.ExistTagById(uriParams.Id) {
		logger.Warnf("cannot find id: %v", uriParams.Id)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistTag
		return
	}

	data := make(map[string]interface{})
	data["name"] = req.Name
	data["state"] = req.State
	service.Tag.EditTag(uriParams.Id, data)
}

func DeleteTag(c *gin.Context) {
	logger := log.Get("DeleteTag").Sugar()
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

	if !service.Tag.ExistTagById(req.Id) {
		logger.Warnf("cannot find id: %v", req.Id)
		result.httpStatus = http.StatusNotFound
		result.Code = CodeNotExistTag
		return
	}
	service.Tag.DeleteTag(req.Id)
}
