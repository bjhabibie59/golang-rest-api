package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"go-rest-api/pkg/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSON(w, http.StatusUnauthorized, "Authorization header is missing", nil)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			response.JSON(w, http.StatusUnauthorized, "Invalid authorization format", nil)
			return
		}

		tokenString := tokenParts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default_secret"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			response.JSON(w, http.StatusUnauthorized, "Invalid or expired token", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
