package interfaces

import (
	"github.com/gesangwidigdo/go-shorturl/dto"
	"github.com/gesangwidigdo/go-shorturl/model"
	"github.com/gin-gonic/gin"
)

type UrlRepository interface {
	Create(url *model.Url) error
	FindByShortUrl(shortUrl string) (*model.Url, error)
	IncrementClicks(url *model.Url) error
}

type UrlService interface {
	CreateShortUrl(request dto.UrlCreateRequest) (dto.UrlCreateResponse, error)
	Redirect(shortUrl string) (string, error)
}

type UrlController interface {
	CreateShortUrl(ctx *gin.Context)
	RedirectToOriginal(ctx *gin.Context)
}