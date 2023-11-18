package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ExtractApiKey(header http.Header) (string, error) {
	const ApiKeyAuthMethod = "ApiKey"

	authString := header.Get("Authorization")
	authFields := strings.Fields(authString)
	if len(authFields) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if authMethod := authFields[0]; authMethod != ApiKeyAuthMethod {
		return "", fmt.Errorf("%s Authorization method is not supported", authMethod)
	}

	return authFields[1], nil
}

type authedHandler func(http.ResponseWriter, *http.Request, User)

func (cfg *ApiConfig) WithAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := ExtractApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, err.Error())
			return
		}

		var user User
		if err := cfg.DB.Where("api_key = ?", apiKey).First(&user).Error; err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		h(w, r, user)
	}
}
