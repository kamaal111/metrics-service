package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-pg/pg/v10"
)

func collectHandler(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = validateCollectPayload(body)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}
	output, err := json.Marshal(response)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Hello string `json:"Hello"`
	}{
		Hello: "Welcome",
	}
	output, err := json.Marshal(response)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
