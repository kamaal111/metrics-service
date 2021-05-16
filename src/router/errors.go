package router

import (
	"encoding/json"
	"net/http"
)

func errorHandler(w http.ResponseWriter, error string, code int) {
	errorResponse := Error{
		Message: error,
		Status:  code,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(errorResponse.Status)
	json.NewEncoder(w).Encode(errorResponse)
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
