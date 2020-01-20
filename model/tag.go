package model

type Tag struct {
	Model

	Name  string `gorm:"size:128" json:"name"`
	State int    `gorm:"type:tinyint(3)"json:"state"`
}
