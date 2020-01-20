package model

type User struct {
	Model

	UserName  string `gorm:"not null;unique_index;size:32" json:"userName"`
	Email     string `gorm:"not null;unique_index;size:64" json:"email"`
	AvatarUrl string `gorm:"type:varchar(255)"`
	Salt      string `gorm:"not null;type:varchar(32)"`
	Password  string `gorm:"not null;type:varchar(32)"`
}
