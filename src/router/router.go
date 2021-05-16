package router

import (
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
)

func HandleRequests(pgDB *pg.DB, port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))
	mux.Handle("/collect", loggerMiddleware(restrictToHttpMethod(http.MethodPost, connectToDatabase(pgDB, collectHandler))))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
