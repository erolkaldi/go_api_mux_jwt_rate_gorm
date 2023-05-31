package middleware

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

func RateCheckLimit(next http.Handler) http.Handler {

	r := rate.Every(30 * time.Second)
	limit := rate.NewLimiter(r, 50)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !limit.Allow() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You exceeded your rate limit"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
