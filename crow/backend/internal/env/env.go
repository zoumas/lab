// Package env provides a way to load and group environment variables into a struct.
package env

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env groups all the environment variables the application needs to run.
// This is always implementation specific and so can't be generalized into a library.
type Env struct {
	// The hostname:post the server binds to and listens for incoming requests.
	Addr string
	// the single URL that will be allowed to make requests to the server. This is to be the frontend.
	CorsOrigin string
}

func Load() (*Env, error) {
	local := flag.Bool("local", false, "Depend on a .env file for local development")
	flag.Parse()

	if *local {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading local .env : %s", err)
		}
	}

	addr, ok := os.LookupEnv("ADDR")
	if !ok {
		return nil, envVarNotSet("ADDR")
	}

	corsOrigin, ok := os.LookupEnv("CORS_ORIGIN")
	if !ok {
		return nil, envVarNotSet("CORS_ORIGIN")
	}

	return &Env{
		Addr:       addr,
		CorsOrigin: corsOrigin,
	}, nil
}

func envVarNotSet(name string) error {
	return fmt.Errorf("%s environment variable is not set", name)
}
