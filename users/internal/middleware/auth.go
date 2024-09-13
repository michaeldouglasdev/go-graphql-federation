package middleware

import (
	"context"
	"fmt"
	"go-graphql-apollo-federation/users/graph/model"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	User model.User `json:"user"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerAuth := r.Header.Get("Authorization")
		if headerAuth == "" {
			next.ServeHTTP(w, r)
		} else {
			user, err := ParseToken(headerAuth)
			if err != nil {
				next.ServeHTTP(w, r)
			} else {
				if err != nil {
					panic("Unknown Error")
				}
				ctx := context.WithValue(r.Context(), "user", user)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
	})
}

func ParseToken(tokenAuth string) (*model.User, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	key := []byte(jwtSecret)

	// Crie uma instância da estrutura de claims
	claims := &UserClaims{}

	// Parse o token e associe-o à estrutura de claims
	token, err := jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	// Verifique se o token é válido e se as claims podem ser convertidas
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return &claims.User, nil
	} else {
		return nil, fmt.Errorf("invalid token or claims")
	}
}
