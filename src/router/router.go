package router

import (
	"log"
	"net/http"
)

func HandleRequests(port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))
	mux.Handle("/register/", loggerMiddleware(restrictToHttpMethod(http.MethodPost, apiKeyRequired(http.HandlerFunc(registerHandler)))))
	mux.Handle("/collect/", loggerMiddleware(restrictToHttpMethod(http.MethodPost, http.HandlerFunc(collectHandler))))
	mux.Handle("/metrics/", loggerMiddleware(restrictToHttpMethod(http.MethodGet, http.HandlerFunc(metricsHandler))))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
