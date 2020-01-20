package model

import (
	"time"
)

type Model struct {
	Id       uint64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreateAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createAt"`
	UpdateAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updateAt"`
}

var AllModels = []interface{}{
	&User{}, &Article{}, &Tag{}, &Setting{},
}
