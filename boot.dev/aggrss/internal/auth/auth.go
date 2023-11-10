package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API Key from the headers of an HTTP request
// Example:
// Authorization: ApiKey <insert apikey here>
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	fields := strings.Fields(val)
	if len(fields) != 2 {
		return "", errors.New("malformed authentication header")
	}

	if fields[0] != "ApiKey" {
		return "", errors.New("malformed first part of authentication header")
	}

	return fields[1], nil
}
