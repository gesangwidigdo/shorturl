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
		shortUrl = utils.GenerateRandomString(6)
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
func (u *urlService) Redirect(shortUrl string) (string, error) {
	url, err := u.urlRepository.FindByShortUrl(shortUrl)
	if err != nil {
		return "", err
	}

	if err := u.urlRepository.IncrementClicks(url); err != nil {
		return "", err
	}

	return url.OriginalUrl, nil
}
