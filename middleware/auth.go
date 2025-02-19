package middleware

import (
	"net/http"
	"time"

	"github.com/gesangwidigdo/go-shorturl/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token verification failed"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("username", claims["sub"])
		c.Set("id", claims["id"])
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
