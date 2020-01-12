package model

import "html/template"

type Article struct {
	Model

	TagID int `gorm:"index" json:"tag_id"`
	Tag   Tag `json:"tag"`

	Title       string        `json:"title"`
	Desc        string        `json:"desc"`
	ContentMD   string        `json:"content_md"`
	ContentHTML template.HTML `gorm:"-" json:"-"`
	State       int           `json:"state"`
}
