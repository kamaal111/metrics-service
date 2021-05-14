package router

import (
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
)

func HandleRequests(pgDB *pg.DB, port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

// func connectToDatabase(pgDB *pg.DB, f func(w http.ResponseWriter, r *http.Request, pgDB *pg.DB)) http.Handler {
// 	funcToPass := func(w http.ResponseWriter, r *http.Request) {
// 		f(w, r, pgDB)
// 	}
// 	return http.HandlerFunc(funcToPass)
// }
