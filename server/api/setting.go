package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/model"
	"github.com/ijkbytes/mega/service"
)

func GetBasicSettings(c *gin.Context) {
	result := &Result{}
	defer returnResult(c, result)

	settings := service.Setting.GetCategorySettings(model.SettingCategoryBasic)
	data := map[string]interface{}{}
	for _, setting := range settings {
		data[setting.Name] = setting.Value
	}

	result.Data = data
}

func UpdateBasicSettings(c *gin.Context) {
	logger := log.Get("UpdateBasicSettings").Sugar()
	result := &Result{}
	defer returnResult(c, result)

	args := make(map[string]interface{})
	if err := c.ShouldBindJSON(&args); err != nil {
		logger.Warnf("invalid params: %v", err.Error())
		result.httpStatus = http.StatusBadRequest
		result.Code = CodeInvalidParams
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
		result.httpStatus = http.StatusInternalServerError
		result.Code = CodeError
		return
	}
}
