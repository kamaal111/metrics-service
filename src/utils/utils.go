package utils

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func MLogger(message string, statusCode int, err error) {
	log.Printf("{message: %s, code: %d, error: %s}\n", message, statusCode, err.Error())
}

func GenerateSecureToken(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func HashAndSalt(bytes []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashAndToken(hashedToken string, plainToken []byte) (bool, error) {
	byteHash := []byte(hashedToken)
	err := bcrypt.CompareHashAndPassword(byteHash, plainToken)
	return err == nil, err
}
