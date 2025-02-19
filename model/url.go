package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string `json:"original_url" gorm:"not null; type:varchar(255)"`
	ShortUrl    string `json:"short_url" gorm:"unique" gorm:"not null"`
	Clicks      int    `json:"clicks" gorm:"default:0"`
	UserID			uint   `json:"user_id" gorm:"not null"`

	User User `json:"user"`
}
