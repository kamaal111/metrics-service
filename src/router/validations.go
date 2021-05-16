package router

import (
	"encoding/json"
	"errors"

	"github.com/kamaal111/metrics-service/src/models"
)

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
	return payload, nil
}
