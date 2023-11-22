package middleware

import (
	"net/http"
)

const AuthToken = "MyCoolStarJeton"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != AuthToken {
			http.Error(w, "Access not authorised", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
