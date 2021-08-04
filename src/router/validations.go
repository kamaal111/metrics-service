package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/kamaal111/metrics-service/src/models"
	"github.com/kamaal111/metrics-service/src/utils"
)

func processAccessToken(headerAccessToken string, appAccessToken string, apiVersion models.APIVersion) (int, error) {
	if headerAccessToken == "" {
		return http.StatusBadRequest, errors.New("access token not found")
	}
	if apiVersion.IsLessThan(models.VERSION_2_0_0) {
		hasValidToken, err := utils.CompareHashAndToken(appAccessToken, headerAccessToken)
		if !hasValidToken {
			return http.StatusUnauthorized, errors.New("unauthorized")
		} else if err != nil {
			utils.MLogger("something went wrong while comparing hash", http.StatusInternalServerError, err)
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	} else {
		if headerAccessToken != appAccessToken {
			return http.StatusUnauthorized, errors.New("unauthorized")
		}
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

func validateCollectPayload(body []byte) (collectionPayload, int, error) {
	var payload collectionPayload
	err := json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return collectionPayload{}, http.StatusInternalServerError, errors.New("something went wrong")
	}
	if payload.AppVersion == "" {
		return collectionPayload{}, http.StatusBadRequest, errors.New("app_version field is required")
	}
	appVersion, err := utils.ParseStringToAPIVersion(payload.AppVersion)
	if err != nil {
		return collectionPayload{}, http.StatusBadRequest, err
	}
	payload.AppVersion = appVersion.ToString()
	if payload.BundleIdentifier == "" {
		return collectionPayload{}, http.StatusBadRequest, errors.New("bundle_identifier field is required")
	}
	if len(payload.Payload) == 0 {
		return collectionPayload{}, http.StatusBadRequest, errors.New("payload field is required")
	}
	if payload.Payload[0].MetaData.AppBuildVersion == "" {
		return collectionPayload{}, http.StatusBadRequest, errors.New("payload.metaData.appBuildVersion field is required")
	}

	return payload, http.StatusOK, nil
}
