package middleware

import (
	"errors"
	"forum-site-app/internal/configs"
	"forum-site-app/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Authorization header is missing"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource.",
			})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Authorization header is missing"))
			return
		}

		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource.",
			})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
