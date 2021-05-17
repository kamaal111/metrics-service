package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/kamaal111/metrics-service/src/db"
	"github.com/kamaal111/metrics-service/src/models"
)

func metricsHandler(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	splittedURLPath := strings.FieldsFunc(r.URL.Path, func(c rune) bool {
		return c == '/'
	})
	if len(splittedURLPath) < 2 {
		errorHandler(w, "use app bundle identifier at the end of this url", http.StatusBadRequest)
		return
	}
	bundleIdentifier := splittedURLPath[1]
	splittedBundleIdentifier := strings.FieldsFunc(bundleIdentifier, func(c rune) bool {
		return c == '.'
	})
	if len(splittedBundleIdentifier) < 2 {
		errorHandler(w, "invalid bundle identifier user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func collectHandler(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload, err := validateCollectPayload(body)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}

	app, err := db.GetOrCreateAppByBundleIdentifier(pgDB, payload.BundleIdentifier)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metrics := models.MetricsTable{
		AppVersion:      payload.AppVersion,
		AppBuildVersion: payload.Payload.MetaData.AppBuildVersion,
		Payload:         payload.Payload,
		AppID:           app.ID,
	}
	err = metrics.Save(pgDB)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
