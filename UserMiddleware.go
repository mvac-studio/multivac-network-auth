package mvac

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UserMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("authorization")

			userHeader := r.Header.Get("user")
			user := &UserContext{}
			err := json.Unmarshal([]byte(userHeader), user)
			if err != nil {
				log.Println(err)
			}
			ctx := context.WithValue(r.Context(), "user", user)
			ctx = context.WithValue(ctx, "token", token)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
