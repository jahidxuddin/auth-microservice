package main

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Method: %s, Path: %s, Time: %s\n", r.Method, r.URL.Path, start.Format(time.RFC1123))
		next(w, r)
		duration := time.Since(start)
		fmt.Printf("Completed in %v\n", duration)
	}
}
