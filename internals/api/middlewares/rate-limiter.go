package middlewares

import (
	"net/http"
	"time"
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Limit the number of requests to 5 per second
		time.Sleep(200 * time.Millisecond)
		next.ServeHTTP(w, r)
	})
}
