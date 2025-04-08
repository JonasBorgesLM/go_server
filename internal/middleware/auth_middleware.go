// Package middleware provides HTTP middleware for handling authentication.
package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/JonasBorgesLM/go_server/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware verifies the JWT token in the Authorization header.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header is required", "Authorization")
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token format", "token")
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token", "token")
			return
		}

		next.ServeHTTP(w, r)
	}
}
