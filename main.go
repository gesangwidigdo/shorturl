package main

import (
	"github.com/gesangwidigdo/go-shorturl/config"
	"github.com/gesangwidigdo/go-shorturl/model"
	"github.com/gesangwidigdo/go-shorturl/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app_name": "Short URL",
		})
	})

	db := config.ConnectDB()
	if err := model.Migrate(db); err != nil {
		panic(err.Error())
	}

	router.Router(r, db)

	r.Run()
}