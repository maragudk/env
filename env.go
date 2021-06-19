// Package env provides helpers functions to load different variable types from the environment.
package env

import (
	"os"
	"strconv"
	"time"
)

// GetStringOrDefault value.
func GetStringOrDefault(name, defaultV string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	return v
}

// GetIntOrDefault value. Also returns the default value if the value is not an int.
func GetIntOrDefault(name string, defaultV int) int {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	vAsInt, err := strconv.Atoi(v)
	if err != nil {
		return defaultV
	}
	return vAsInt
}

// GetBoolOrDefault value. Also returns the default value if the value is not a boolean.
func GetBoolOrDefault(name string, defaultV bool) bool {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	vAsBool, err := strconv.ParseBool(v)
	if err != nil {
		return defaultV
	}
	return vAsBool
}

// GetDurationOrDefault value. Also returns the default value if the value is not a time.Duration.
func GetDurationOrDefault(name string, defaultV time.Duration) time.Duration {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	vAsDuration, err := time.ParseDuration(v)
	if err != nil {
		return defaultV
	}
	return vAsDuration
}
