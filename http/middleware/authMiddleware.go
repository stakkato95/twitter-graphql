package middleware

import (
	"context"
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
	"github.com/stakkato95/twitter-service-graphql/http/service"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Auth(service service.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			// for auth endpoints
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			user, err := service.Authorize(token)
			if err != nil {
				logger.Error("can not authorize by token: " + err.Error())
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *dto.UserDto {
	raw, _ := ctx.Value(userCtxKey).(*dto.UserDto)
	return raw
}
