package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kamaal111/metrics-service/src/models"
	"github.com/kamaal111/metrics-service/src/utils"
)

func processAccessToken(headerAccessToken string, appAccessToken string) (int, error) {
	if headerAccessToken == "" {
		return http.StatusBadRequest, errors.New("access token not found")
	}
	hasValidToken, err := compareHashAndToken(appAccessToken, []byte(headerAccessToken))
	if !hasValidToken {
		return http.StatusUnauthorized, errors.New("unauthorized")
	} else if err != nil {
		utils.MLogger("something went wrong while comparing hash", http.StatusInternalServerError, err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func validateCollectPayload(body []byte) (models.CollectionPayload, error) {
	var payload models.CollectionPayload
	err := json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return models.CollectionPayload{}, err
	}
	if payload.AppVersion == "" {
		return models.CollectionPayload{}, errors.New("app_version field is required")
	}
	if payload.BundleIdentifier == "" {
		return models.CollectionPayload{}, errors.New("bundle_identifier field is required")
	}
	if payload.Payload.MetaData.AppBuildVersion == "" {
		return models.CollectionPayload{}, errors.New("payload.metaData.appBuildVersion field is required")
	}
	return payload, nil
}
