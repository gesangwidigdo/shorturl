package controller

import (
	"github.com/gesangwidigdo/go-shorturl/dto"
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gin-gonic/gin"
)

type urlController struct {
	urlService interfaces.UrlService
}

func NewUrlController(urlService interfaces.UrlService) interfaces.UrlController {
	return &urlController{urlService}
}

// CreateShortUrl implements interfaces.UrlController.
func (u *urlController) CreateShortUrl(ctx *gin.Context) {
	var urlReq dto.UrlCreateRequest
	if err := ctx.ShouldBindJSON(&urlReq); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	urlRes, err := u.urlService.CreateShortUrl(urlReq)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"short_url": urlRes.ShortUrl})
}

// GetOriginalUrl implements interfaces.UrlController.
func (u *urlController) RedirectToOriginal(ctx *gin.Context) {
	shortUrl := ctx.Param("short_url")
	originalUrl, err := u.urlService.Redirect(shortUrl)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(301, originalUrl)
}
