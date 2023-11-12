package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const AuthApiKeyMethod = "ApiKey"

func GetApiKey(h http.Header) (string, error) {
	fields := strings.Fields(h.Get("Authorization"))
	if len(fields) != 2 {
		return "", errors.New("malformed Authorization request header")
	}

	authMethod := fields[0]
	if authMethod != AuthApiKeyMethod {
		return "", fmt.Errorf("Authorization method %q is not supported", authMethod)
	}

	return fields[1], nil
}
