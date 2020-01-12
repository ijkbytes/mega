package model

import (
	"time"
)

type Model struct {
	Id       uint64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreateAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}
