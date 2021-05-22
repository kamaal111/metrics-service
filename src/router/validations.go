package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/kamaal111/metrics-service/src/utils"
)

func processAccessToken(headerAccessToken string, appAccessToken string) (int, error) {
	if headerAccessToken == "" {
		return http.StatusBadRequest, errors.New("access token not found")
	}
	hasValidToken, err := utils.CompareHashAndToken(appAccessToken, headerAccessToken)
	if !hasValidToken {
		return http.StatusUnauthorized, errors.New("unauthorized")
	} else if err != nil {
		utils.MLogger("something went wrong while comparing hash", http.StatusInternalServerError, err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func validateBundleIdentifier(bundleIdentifier string) (string, error) {
	if bundleIdentifier == "" {
		return "", errors.New("bundle_identifier is required")
	}
	splittedBundleIdentifier := strings.FieldsFunc(bundleIdentifier, func(c rune) bool {
		return c == '.'
	})
	if len(splittedBundleIdentifier) < 2 {
		return "", errors.New("invalid bundle identifier")
	}
	return bundleIdentifier, nil
}

func validateCollectPayload(body []byte) (collectionPayload, error) {
	var payload collectionPayload
	err := json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return collectionPayload{}, err
	}
	if payload.AppVersion == "" {
		return collectionPayload{}, errors.New("app_version field is required")
	}
	if payload.BundleIdentifier == "" {
		return collectionPayload{}, errors.New("bundle_identifier field is required")
	}
	if payload.Payload.MetaData.AppBuildVersion == "" {
		return collectionPayload{}, errors.New("payload.metaData.appBuildVersion field is required")
	}
	return payload, nil
}
