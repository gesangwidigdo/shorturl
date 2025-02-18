package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, db *gorm.DB) {
	urlRoutes := r.Group("/api/url")
	URLRouter(urlRoutes, db)
}
