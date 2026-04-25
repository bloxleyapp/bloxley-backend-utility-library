package utilslibs

import (
	"log"
	"os"
	"strconv"
)

// getEnvOrFail is a helper that crashes if an env variable is missing or empty.
func GetEnvOrFail(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return val
}

// getIntEnvOrFail is a helper that crashes if an env variable is missing or empty.
func GetIntEnvOrFail(key string) int {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}

	// Convert string to int
	intVal, err := strconv.Atoi(val)

	if err != nil {
		log.Fatalf("failed to convert environment variable %s to int", key)
	}

	return intVal
}

// GetBoolEnvOrFail is a helper that crashes if an env variable is missing, empty, or not a valid boolean.
func GetBoolEnvOrFail(key string) bool {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}

	// Convert string to bool (accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False)
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("failed to convert environment variable %s to bool", key)
	}

	return boolVal
}

// getEnvOrEmpty is a helper that returns an env variable or an empty string if it's missing.
func GetEnvOrEmpty(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return ""
	}
	return val
}
