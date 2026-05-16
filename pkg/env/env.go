package env

import (
	"fmt"
	"log/slog"
	"os"
)

func Get(key, default_value string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	slog.Warn("env variable not found", "key", key, "default", default_value)

	return default_value
}

func GetRequired(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("required env variable %s is not set", key)
	}

	return value, nil
}
