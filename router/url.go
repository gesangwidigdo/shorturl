package router

import (
	"github.com/gesangwidigdo/go-shorturl/controller"
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gesangwidigdo/go-shorturl/repository"
	"github.com/gesangwidigdo/go-shorturl/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func URLRouter(r *gin.RouterGroup, db *gorm.DB) {
	var urlRepository interfaces.UrlRepository = repository.NewUrlRepository(db)
	var urlService interfaces.UrlService = service.NewUrlService(urlRepository)
	var urlController interfaces.UrlController = controller.NewUrlController(urlService)

	r.POST("/shorten", urlController.CreateShortUrl)
	r.GET("/:short_url", urlController.RedirectToOriginal)
}
