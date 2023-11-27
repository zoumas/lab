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
	// The hostname:post the server binds to and listens for incoming requests
	Addr string
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
		return nil, EnvVarNotSet("ADDR")
	}

	return &Env{
		Addr: addr,
	}, nil
}

func EnvVarNotSet(name string) error {
	return fmt.Errorf("%s environment variable is not set", name)
}

// Empty returns a empty Env struct.
// This is used in testing.
func Empty() *Env {
	return &Env{}
}
