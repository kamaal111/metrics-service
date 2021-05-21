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
	return validateBundleIdentifier(splittedURLPath[1])
}
