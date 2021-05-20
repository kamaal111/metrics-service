package router

import (
	"errors"
	"strings"
)

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
