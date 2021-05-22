package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

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
