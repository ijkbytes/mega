package service

import "github.com/ijkbytes/mega/model"

type tagService struct {
}

var Tag = &tagService{}

func (srv *tagService) GetTags(offset int, limit int, maps interface{}) (tags []*model.Tag) {
	db.Where(maps).Offset(offset).Limit(limit).Find(&tags)
	return
}

func (srv *tagService) GetTagTotal(maps interface{}) (count int) {
	db.Model(&model.Tag{}).Where(maps).Count(&count)
	return
}

func (srv *tagService) ExistTagByName(name string) bool {
	var tag model.Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

func (srv *tagService) AddTag(name string, state int) bool {
	db.Create(&model.Tag{
		Name:  name,
		State: state,
	})

	return true
}

func (srv *tagService) ExistTagById(id int) bool {
	var tag model.Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

func (srv *tagService) DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&model.Tag{})

	return true
}

func (srv *tagService) EditTag(id int, data interface{}) bool {
	db.Model(&model.Tag{}).Where("id = ?", id).Update(data)

	return true
}
