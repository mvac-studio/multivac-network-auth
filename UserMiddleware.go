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
			userHeader := r.Header.Get("User")
			user := &UserContext{}
			err := json.Unmarshal([]byte(userHeader), user)
			if err != nil {
				log.Println(err)
			}
			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
