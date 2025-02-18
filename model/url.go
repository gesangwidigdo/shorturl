package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string `json:"original_url" gorm:"not null; type:varchar(255)"`
	ShortUrl    string `json:"short_url" gorm:"unique"`
	Clicks      int    `json:"clicks" gorm:"default:0"`
}