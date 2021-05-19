package router

import (
	"crypto/rand"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func generateSecureToken(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func hashAndSalt(bytes []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareApiToken(hashedApiToken string, plainApiToken []byte) (bool, error) {
	byteHash := []byte(hashedApiToken)
	err := bcrypt.CompareHashAndPassword(byteHash, plainApiToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getBundleIdentifierFromURLPath(path string) (string, error) {
	splittedURLPath := strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})
	if len(splittedURLPath) < 2 {
		return "", errors.New("use app bundle identifier at the end of this url")
	}
	bundleIdentifier := splittedURLPath[1]
	splittedBundleIdentifier := strings.FieldsFunc(bundleIdentifier, func(c rune) bool {
		return c == '.'
	})
	if len(splittedBundleIdentifier) < 2 {
		return "", errors.New("invalid bundle identifier")
	}
	return bundleIdentifier, nil
}
