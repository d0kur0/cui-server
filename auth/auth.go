package auth

import (
	"context"
	"net/http"

	"github.com/d0kur0/cui-server/graph/model"

	"github.com/d0kur0/cui-server/database"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userToken := r.Header.Get("USER_TOKEN")

			if len(userToken) == 0 {
				next.ServeHTTP(w, r)
				return
			}

			user := database.ValidateAndGetUser(userToken)
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
