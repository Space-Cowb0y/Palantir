// middleware

package http

import (
	"net/http"
)

// Example middleware function
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request here
		next.ServeHTTP(w, r)
	})
}
