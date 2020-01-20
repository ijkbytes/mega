package service

import (
	"github.com/ijkbytes/mega/model"
	"sync"
)

type settingService struct {
	mutex *sync.Mutex
}

var Setting = &settingService{
	mutex: &sync.Mutex{},
}

func (srv *settingService) GetSetting(category, name string) *model.Setting {
	ret := &model.Setting{}
	if err := db.Where("`category` = ? AND `name` = ?", category, name).Find(ret).Error; err != nil {
		return nil
	}

	return ret
}

func (srv *settingService) GetCategorySettings(category string) []*model.Setting {
	var ret []*model.Setting

	if err := db.Where("`category` = ?", category).Find(&ret).Error; err != nil {
		return nil
	}

	return ret
}

func (srv *settingService) AddSetting(setting *model.Setting) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	if nil != srv.GetSetting(setting.Category, setting.Name) {
		return nil
	}

	tx := db.Begin()
	if err := tx.Create(setting).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *settingService) UpdateSettings(category string, settings []*model.Setting) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	for _, setting := range settings {
		if err := tx.Model(&model.Setting{}).
			Where("`category` = ? AND `name` = ?", category, setting.Name).
			Select("value").Updates(map[string]interface{}{"value": setting.Value}).Error; nil != err {

			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	return nil
}
