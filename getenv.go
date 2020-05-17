package getenv

import (
	"fmt"
	"log"
	"os"
)

// DefaultString -
type DefaultString *string

// GetEnv - main function to fetch for environment variable or return default
func GetEnv(key string, d DefaultString) string {
	value := os.Getenv(key)
	if len(value) == 0 && d != nil {
		log.Print(fmt.Sprintf("%s is not found. fallback to default value.", key))
		return Value(d)
	} else if len(value) == 0 && d == nil {
		log.Panicf("%s is a required environment variable", key)
	}

	return value
}

// Nil - custom fn to return nil
func Nil() *string {
	return nil
}

// String - custom fn to return the pointer of a string
func String(s string) *string {
	newS := s
	return &newS
}

// Value - returns the default value as a string
func Value(defaultValue DefaultString) string {
	return *defaultValue
}
