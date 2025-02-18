package repository

import (
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gesangwidigdo/go-shorturl/model"
	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}


func NewUrlRepository(db *gorm.DB) interfaces.UrlRepository {
	return &urlRepository{db}
}


// Create implements interfaces.UrlRepository.
func (u *urlRepository) Create(url *model.Url) error {
	if err := u.db.Create(url).Error; err != nil {
		return err
	}

	return nil
}

// FindByShortUrl implements interfaces.UrlRepository.
func (u *urlRepository) FindByShortUrl(shortUrl string) (*model.Url, error) {
	var url model.Url
	if err := u.db.Where("short_url = ?", shortUrl).First(&url).Error; err != nil {
		return nil, err
	}

	return &url, nil
}

// IncrementClicks implements interfaces.UrlRepository.
func (u *urlRepository) IncrementClicks(url *model.Url) error {
	if err := u.db.Model(url).Update("clicks", url.Clicks+1).Error; err != nil {
		return err
	}

	return nil
}