package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

const PORT = "127.0.0.1:8080"

func HandleRequests(pgDB *pg.DB) {
	getRoot := loggerMiddleware(connectToDatabase(pgDB, root))

	mux := http.NewServeMux()

	mux.Handle("/", getRoot)

	log.Printf("Listening on %s\n", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}

func root(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	payload := struct {
		Hello string `json:"Hello"`
	}{
		Hello: "Welcome",
	}
	output, err := json.Marshal(payload)
	if err != nil {
		errorHandler(w, err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

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

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func errorHandler(w http.ResponseWriter, error string, code int) {
	errorResponse := Error{
		Message: error,
		Status:  code,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(errorResponse.Status)
	json.NewEncoder(w).Encode(errorResponse)
}
