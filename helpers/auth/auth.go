package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mariasilva795/go-api-rest/models"
	"github.com/mariasilva795/go-api-rest/server"
)

func ValidateToken(s server.Server, r *http.Request) (string, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	if tokenString == "" {
		return "", fmt.Errorf("no Authorization header provided")
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
