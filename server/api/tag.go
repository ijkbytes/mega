package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/service"
	"net/http"
)

func GetTags(c *gin.Context) {
	logger := log.Get("GetTags").Sugar()
	var req struct {
		Name  string `form:"name"`
		State int    `form:"state" binding:"oneof=-1 0 1"`
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

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	logger := log.Get("AddTag").Sugar()
	var req struct {
		Name  string `json:"name" binding:"required"`
		State int    `json:"state" binding:"oneof=0 1"`
	}
	if err := c.ShouldBind(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	service.Tag.AddTag(req.Name, req.State)

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}

func EditTag(c *gin.Context) {
	logger := log.Get("EditTag").Sugar()
	var req struct {
		Name  string `json:"name" binding:"required"`
		State int    `json:"state" binding:"oneof=0 1"`
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
	if err := c.ShouldBind(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	if !service.Tag.ExistTagById(uriParams.Id) {
		logger.Warnf("cannot find id: %v", uriParams.Id)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistTag,
			"msg":  getErrMsg(CodeNotExistTag),
		})
		return
	}

	data := make(map[string]interface{})
	data["name"] = req.Name
	data["state"] = req.State
	service.Tag.EditTag(uriParams.Id, data)

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}

func DeleteTag(c *gin.Context) {
	logger := log.Get("DeleteTag").Sugar()
	var req struct {
		Id int `uri:"id" bindig:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	if !service.Tag.ExistTagById(req.Id) {
		logger.Warnf("cannot find id: %v", req.Id)
		c.JSON(http.StatusNotFound, gin.H{
			"code": CodeNotExistTag,
			"msg":  getErrMsg(CodeNotExistTag),
		})
		return
	}
	service.Tag.DeleteTag(req.Id)

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
	})
}
