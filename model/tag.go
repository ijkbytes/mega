package model

type Tag struct {
	Model

	Name  string `gorm:"size:128" json:"name"`
	State int    `json:"state"`
}
