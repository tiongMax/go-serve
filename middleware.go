package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("Request: %s %s - %v", r.Method, r.URL.Path, duration)
	}
}

func panicRecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC recovered: %v | Method: %s | Path: %s", err, r.Method, r.URL.Path)
				writeError(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		next(w, r)
	}
}

func WithDefaults(next http.HandlerFunc) http.HandlerFunc {
	return panicRecoveryMiddleware(loggingMiddleware(next))
}
