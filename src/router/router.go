package router

import (
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
)

func HandleRequests(pgDB *pg.DB, port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))
	mux.Handle("/register/", loggerMiddleware(restrictToHttpMethod(http.MethodPost, connectToDatabase(pgDB, registerHandler))))
	mux.Handle("/collect/", loggerMiddleware(restrictToHttpMethod(http.MethodPost, connectToDatabase(pgDB, collectHandler))))
	mux.Handle("/metrics/", loggerMiddleware(restrictToHttpMethod(http.MethodGet, connectToDatabase(pgDB, metricsHandler))))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
