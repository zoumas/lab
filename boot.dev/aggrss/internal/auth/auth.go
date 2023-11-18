package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const authByApiKey = "ApiKey"

func GetApiKey(h http.Header) (string, error) {
	authorizationHeader := h.Get("Authorization")

	authorizationFields := strings.Fields(authorizationHeader)
	if len(authorizationFields) != 2 {
		return "", errors.New("malformed request body")
	}

	authorizationMethod := authorizationFields[0]
	if authorizationMethod != authByApiKey {
		return "", fmt.Errorf("%s authorization method is not supported", authorizationMethod)
	}

	return authorizationFields[1], nil
}
