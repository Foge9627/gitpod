package middleware

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// TODO write into bytesbuffer so we can access it
// create an instance of Logrus

type Middleware func(handler http.Handler) http.Handler

func NewLoggingMiddleware() Middleware {
	logInMemory := &bytes.Buffer{}
	log.SetOutput(logInMemory)

	return func(next http.Handler) http.Handler {
		logging := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			uri := r.RequestURI
			method := r.Method
			duration := time.Since(start)
			log.WithFields(log.Fields{
				"uri":      uri,
				"method":   method,
				"duration": duration,
			})
			fmt.Println(logInMemory)
			next.ServeHTTP(w, r)
		})

		return logging
	}
}
