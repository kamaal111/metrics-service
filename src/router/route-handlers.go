package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-pg/pg/v10"

	"github.com/kamaal111/metrics-service/src/db"
	"github.com/kamaal111/metrics-service/src/models"
	"github.com/kamaal111/metrics-service/src/utils"
)

func registerHandler(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	accessToken, err := generateSecureToken(32)
	if err != nil {
		utils.MLogger("something went wrong while generating secure token", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedToken, err := hashAndSalt(accessToken)
	if err != nil {
		utils.MLogger("something went wrong while hashing and salting access token", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: hashedToken,
	}
	output, err := json.Marshal(response)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func metricsHandler(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	bundleIdentifier, err := getBundleIdentifierFromURLPath(r.URL.Path)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	if bundleIdentifier == "" {
		errorHandler(w, "invalid bundle_identifier in path", http.StatusBadRequest)
		return
	}

	app, err := db.GetAppByBundleIdentifier(pgDB, bundleIdentifier)
	if err == pg.ErrNoRows {
		errorHandler(w, fmt.Sprintf("%s not found", bundleIdentifier), http.StatusNotFound)
		return
	} else if err != nil {
		utils.MLogger("something went wrong while getting app by bundle identifier", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessTokenCode, err := processAccessToken(r.Header.Get("access_token"), app.AccessToken)
	if err != nil {
		errorHandler(w, err.Error(), accessTokenCode)
		return
	}

	metrics, err := app.GetMetrics(pgDB)
	if err == pg.ErrNoRows {
		errorHandler(w, "metrics not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.MLogger("something went wrong while getting metrics from app", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(metrics)
	if err != nil {
		utils.MLogger("something went wrong while marshalling metrics", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
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

	app, err := db.GetAppByBundleIdentifier(pgDB, payload.BundleIdentifier)
	if err == pg.ErrNoRows {
		errorHandler(w, "app not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.MLogger("something went wrong while getting app with bundle identifier", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessTokenCode, err := processAccessToken(r.Header.Get("access_token"), app.AccessToken)
	if err != nil {
		errorHandler(w, err.Error(), accessTokenCode)
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
		utils.MLogger("something went wrong while saving metrics", http.StatusInternalServerError, err)
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
		utils.MLogger("something went wrong while marshaling response", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
