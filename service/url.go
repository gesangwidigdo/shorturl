package service

import (
	"errors"
	"strings"

	"github.com/gesangwidigdo/go-shorturl/dto"
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gesangwidigdo/go-shorturl/model"
	"github.com/gesangwidigdo/go-shorturl/utils"
)

type urlService struct {
	urlRepository interfaces.UrlRepository
}

func NewUrlService(urlRepository interfaces.UrlRepository) interfaces.UrlService {
	return &urlService{urlRepository}
}

// CreateShortUrl implements interfaces.UrlService.
func (u *urlService) CreateShortUrl(request dto.UrlCreateRequest) (dto.UrlCreateResponse, error) {
	if strings.Contains(request.ShortUrl, " ") || strings.Contains(request.ShortUrl, "/") {
		return dto.UrlCreateResponse{}, errors.New("short url cannot contain spaces or slashes")
	}

	var shortUrl string

	var url model.Url
	if request.ShortUrl == "" {
		var err error
		shortUrl, err = utils.EncryptURL(request.OriginalUrl)
		if err != nil {
			return dto.UrlCreateResponse{}, err
		}
		
		url = model.Url{
			OriginalUrl: request.OriginalUrl,
			ShortUrl:    shortUrl,
		}

	} else if request.ShortUrl != "" {

		url = model.Url{
			OriginalUrl: request.OriginalUrl,
			ShortUrl:    request.ShortUrl,
		}
		
	}

	if err := u.urlRepository.Create(&url); err != nil {
		return dto.UrlCreateResponse{}, err
	}

	return dto.UrlCreateResponse{ShortUrl: url.ShortUrl}, nil
}

// GetOriginalUrl implements interfaces.UrlService.
func (u *urlService) GetOriginalUrl(shortUrl string) (string, error) {
	url, err := u.urlRepository.FindByShortUrl(shortUrl)
	if err != nil {
		return "", err
	}

	return url.OriginalUrl, nil
}

// IncrementClicks implements interfaces.UrlService.
func (u *urlService) IncrementClicks(shortUrl string) error {
	panic("unimplemented")
}
