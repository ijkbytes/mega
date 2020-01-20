package model

import "html/template"

type Article struct {
	Model

	TagID int `gorm:"index" json:"tagId"`
	Tag   Tag `json:"tag"`

	Title       string        `gorm:"size:64;not null" json:"title"`
	Desc        string        `gorm:"size:100" json:"desc"`
	ContentMD   string        `gorm:"type:text;not null" json:"contentMD"`
	ContentHTML template.HTML `gorm:"-" json:"-"`
	State       int           `gorm:"type:tinyint(3)" json:"state"`
}
