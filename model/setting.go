package model

type Setting struct {
	Model

	Category string `gorm:"size:32;index;not null" json:"category"`
	Name     string `gorm:"size:64;index;not null" json:"name"`
	Value    string `gorm:"type:text" json:"value"`
}

const (
	SettingCategoryBasic = "basic"

	SettingNameBasicBlogTitle  = "basicBlogTitle"
	SettingNameBasicFaviconURL = "basicFaviconURL"
)
