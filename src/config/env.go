package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// unwrap removes wrapping quotes if they exists
func unwrap(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}

	return s
}

// Getenv returns an environment variable value by key.
func Getenv(key string) string {
	val := unwrap(os.Getenv(key))
	if val == "" {
		panic(fmt.Sprintf("Failed to get <%s> environment variable", key))
	}

	return val
}

// GetenvInt returns an environment variable casted to int by key.
func GetenvInt(key string) int {
	val, err := strconv.Atoi(Getenv(key))
	if err != nil {
		panic(fmt.Sprintf("Failed to cast <%s> environment variable to int", key))
	}
	return val
}

// GetenvMap returns an environment variable casted to string map by key.
func GetenvMap(key string) map[string]string {
	var (
		m   = Getenv(key)
		res map[string]string
	)

	err := json.Unmarshal([]byte(m), &res)
	if err != nil {
		panic(fmt.Sprintf("Failed to cast <%s> environment variable to string map", key))
	}

	return res
}
