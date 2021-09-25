package handler

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func WithLogging(h http.HandlerFunc) http.HandlerFunc {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(rw, r) // serve the original request

		duration := time.Since(start)

		// log request details
		log.Info().
			Str("uri", uri).
			Str("method", method).
			Dur("duration", duration).
			Msg("Received request")
	}

	return http.HandlerFunc(logFn)
}
