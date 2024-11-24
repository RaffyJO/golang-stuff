package middleware

import (
	"music-app/internal/configs"
	"music-app/internal/models/response"
	"music-app/pkg/jwt"
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
			c.JSON(http.StatusUnauthorized, response.WebResponse{
				Status:  "Unauthorized",
				Message: "Authorization header is missing",
				Data:    nil,
			})
			c.Abort()
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.WebResponse{
				Status:  "Unauthorized",
				Message: "You are not authorized to access this resource",
				Data:    nil,
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
			c.JSON(http.StatusUnauthorized, response.WebResponse{
				Status:  "Unauthorized",
				Message: "Authorization header is missing",
				Data:    nil,
			})
			c.Abort()
			return
		}

		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.WebResponse{
				Status:  "Unauthorized",
				Message: "You are not authorized to access this resource",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
