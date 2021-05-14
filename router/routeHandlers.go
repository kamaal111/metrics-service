package router

import (
	"encoding/json"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
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
