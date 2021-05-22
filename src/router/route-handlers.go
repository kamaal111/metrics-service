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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var payload registerPayload
	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bundleIdentifier, err := validateBundleIdentifier(payload.BundleIdentifier)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.GetAppByBundleIdentifier(db.PGDatabase, bundleIdentifier)
	if err != pg.ErrNoRows {
		errorHandler(w, fmt.Sprintf("app with %s as bundle_identifier already exists", bundleIdentifier), http.StatusConflict)
		return
	}

	accessToken, err := utils.GenerateSecureToken(32)
	if err != nil {
		utils.MLogger("something went wrong while generating secure token", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedToken, err := utils.HashAndSalt(accessToken)
	if err != nil {
		utils.MLogger("something went wrong while hashing and salting access token", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app := models.AppsTable{
		BundleIdentifier: bundleIdentifier,
		AccessToken:      hashedToken,
	}
	err = app.Save(db.PGDatabase)
	if err != nil {
		utils.MLogger("something went wrong while saving app", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}

	output, err := json.Marshal(response)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	bundleIdentifier, err := getBundleIdentifierFromURLPath(r.URL.Path)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	if bundleIdentifier == "" {
		errorHandler(w, "invalid bundle_identifier in path", http.StatusBadRequest)
		return
	}

	app, err := db.GetAppByBundleIdentifier(db.PGDatabase, bundleIdentifier)
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

	metrics, err := app.GetMetrics(db.PGDatabase)
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

func collectHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload, errorCode, err := validateCollectPayload(body)
	if err != nil {
		errorHandler(w, err.Error(), errorCode)
		return
	}

	app, err := db.GetAppByBundleIdentifier(db.PGDatabase, payload.BundleIdentifier)
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

	for _, metricsPayload := range payload.Payload {
		metrics := models.MetricsTable{
			AppVersion:      payload.AppVersion,
			AppBuildVersion: metricsPayload.MetaData.AppBuildVersion,
			Payload:         metricsPayload,
			AppID:           app.ID,
		}
		err = metrics.Save(db.PGDatabase)
		if err != nil {
			utils.MLogger("something went wrong while saving metrics", http.StatusInternalServerError, err)
			errorHandler(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
