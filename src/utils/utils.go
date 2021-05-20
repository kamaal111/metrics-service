package utils

import "log"

func MLogger(message string, statusCode int, err error) {
	log.Printf("{message: %s, code: %d, error: %s}", message, statusCode, err.Error())
}
