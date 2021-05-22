package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func apiKeyRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("api_key") != os.Getenv("SECRET_TOKEN") && os.Getenv("SECRET_TOKEN") != "" {
			errorHandler(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func restrictToHttpMethod(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			errorHandler(w, fmt.Sprintf("%s does not allow %s. use %s instead", r.URL.String(), r.Method, method), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("%s: %s in %s\n", r.Method, r.URL.Path, elapsed)
	})
}
