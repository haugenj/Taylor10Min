package utils

import (
	"os"
	"strconv"
)

// WithDefaultString returns the string value of the supplied environ variable or, if not
// present, the supplied default value
func WithDefaultString(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

// WithDefaultInt returns the int value of the supplied environ variable or, if not present,
// the supplied default value. If the int conversion fails, returns the default
func WithDefaultInt(key string, def int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

// WithDefaultBool returns the boolvalue of the supplied environ variable or, if not present,
// the supplied default value. If the conversion fails, returns the default
func WithDefaultBool(key string, def bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return b
}
