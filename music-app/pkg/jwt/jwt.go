package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint, username string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       userID,
			"username": username,
			"exp":      time.Now().Add(time.Minute * 60).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string) (uint, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("Invalid token")
	}
	return uint(claims["id"].(float64)), claims["username"].(string), nil
}

func ValidateTokenWithoutExpiry(tokenString string, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("Invalid token")
	}
	return int64(claims["id"].(float64)), claims["username"].(string), nil
}
