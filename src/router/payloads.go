package router

import (
	"github.com/kamaal111/metrics-service/src/models"
)

type registerPayload struct {
	BundleIdentifier string `json:"bundle_identifier"`
}

type collectionPayload struct {
	BundleIdentifier string                     `json:"bundle_identifier"`
	AppVersion       string                     `json:"app_version"`
	Payload          []models.CollectionMetrics `json:"payload"`
}
