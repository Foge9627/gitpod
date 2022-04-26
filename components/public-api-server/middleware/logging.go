package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	logging := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method
		duration := time.Since(start)
		fmt.Println("Called", uri, "with", method, "took", duration)
		next.ServeHTTP(w, r)
	})

	return logging
}