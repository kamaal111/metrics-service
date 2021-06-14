package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kamaal111/metrics-service/src/utils"
)

func withDeprecateEndpoint(versionString string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerVersionString := r.Header.Get("version")
		if headerVersionString == "" {
			// TODO: Deprecate this
			headerVersionString = "1.0"
		}
		headerVersion, err := utils.ParseStringToAPIVersion(headerVersionString)
		if err != nil {
			errorHandler(w, err.Error(), http.StatusBadRequest)
			return
		}
		version, err := utils.ParseStringToAPIVersion(versionString)
		if err != nil {
			errorHandler(w, err.Error(), http.StatusBadRequest)
			return
		}
		if headerVersion.IsHigherThan(version) {
			errorHandler(w, "this endpoint has been deprecated", http.StatusGone)
			return
		}
		next.ServeHTTP(w, r)
	})
}

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
		observer := &responseObserver{ResponseWriter: w}
		start := time.Now()
		next.ServeHTTP(observer, r)
		elapsed := time.Since(start)
		log.Printf("%d %s: %s in %s\n", observer.status, r.Method, r.URL.Path, elapsed)
	})
}

type responseObserver struct {
	http.ResponseWriter
	status int
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	o.status = code
}
