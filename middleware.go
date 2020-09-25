package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func (app *App) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("%v %v %v ms", r.Method, r.RequestURI, time.Since(now).Milliseconds())
	})
}
