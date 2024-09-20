package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/kish1n/GiAuth/internal/config"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func JWTMiddleware(cfg config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt_token")
			if err != nil {
				if err == http.ErrNoCookie {
					ape.RenderErr(w, problems.Unauthorized())
					return
				}
				ape.RenderErr(w, problems.InternalError())
				return
			}

			tokenStr := cookie.Value

			claims := &jwt.StandardClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return cfg.JWT().SecretKey, nil
			})

			if err != nil || !token.Valid {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims.Subject)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
