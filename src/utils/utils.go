package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/kamaal111/metrics-service/src/models"
	"golang.org/x/crypto/bcrypt"
)

func ParseStringToAPIVersion(versionString string) (models.APIVersion, error) {
	splittedVersionString := strings.FieldsFunc(versionString, func(c rune) bool {
		return c == '.'
	})
	var apiVersion models.APIVersion
	if len(splittedVersionString) < 2 {
		return apiVersion, errors.New("invalid version string")
	}
	major, err := strconv.Atoi(splittedVersionString[0])
	if err != nil {
		return apiVersion, errors.New("invalid major version")
	}
	apiVersion.Major = major
	minor, err := strconv.Atoi(splittedVersionString[1])
	if err != nil {
		return apiVersion, errors.New("invalid minor version")
	}
	apiVersion.Minor = minor
	if len(splittedVersionString) < 3 {
		return apiVersion, nil
	}
	patch, err := strconv.Atoi(splittedVersionString[2])
	if err != nil {
		return apiVersion, errors.New("invalid patch version")
	}
	apiVersion.Patch = patch
	return apiVersion, nil
}

func MLogger(message string, statusCode int, err error) {
	log.Printf("{message: %s, code: %d, error: %s}\n", message, statusCode, err.Error())
}

func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func HashAndSalt(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashAndToken(hashedToken string, plainToken string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(plainToken))
	if err != nil {
		return false, err
	}
	return true, nil
}
