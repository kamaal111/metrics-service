package router

import (
	"errors"
	"strings"
)

func getBundleIdentifierFromURLPath(path string) (string, error) {
	splittedURLPath := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})
	if len(splittedURLPath) < 3 {
		return "", errors.New("no bundle identifier defined")
	}
	return validateBundleIdentifier(splittedURLPath[2])
}
