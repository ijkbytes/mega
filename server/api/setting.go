package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/model"
	"github.com/ijkbytes/mega/service"
	"net/http"
)

func GetBasicSettings(c *gin.Context) {
	settings := service.Setting.GetCategorySettings(model.SettingCategoryBasic)
	data := map[string]interface{}{}
	for _, setting := range settings {
		data[setting.Name] = setting.Value
	}

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  getErrMsg(CodeSuccess),
		"data": data,
	})
}

func UpdateBasicSettings(c *gin.Context) {
	logger := log.Get("UpdateBasicSettings").Sugar()

	args := make(map[string]interface{})
	if err := c.ShouldBindJSON(&args); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeInvalidParams,
			"msg":  getErrMsg(CodeInvalidParams),
		})
		return
	}

	var basics []*model.Setting
	for k, v := range args {
		basic := &model.Setting{
			Category: model.SettingCategoryBasic,
			Name:     k,
			Value:    v.(string),
		}
		basics = append(basics, basic)
	}

	if err := service.Setting.UpdateSettings(model.SettingCategoryBasic, basics); err != nil {
		logger.Errorf("update settings err: %v", err)
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
