package util

import (
	"log"
	"os"
)

// MustGetEnv returns the value of the environment variable with the given name. If the variable is not present, it will log a fatal error.
func MustGetEnv(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Panicf("Environment variable %s not set", key)
	}

	return value
}
