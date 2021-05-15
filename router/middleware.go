package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

func connectToDatabase(pgDB *pg.DB, f func(w http.ResponseWriter, r *http.Request, pgDB *pg.DB)) http.Handler {
	funcToPass := func(w http.ResponseWriter, r *http.Request) {
		f(w, r, pgDB)
	}
	return http.HandlerFunc(funcToPass)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("%s: %s in %s\n", r.Method, r.URL.Path, elapsed)
	})
}
