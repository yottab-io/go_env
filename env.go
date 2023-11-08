package env

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

var ErrRequiredVariable = errors.New("error bad Env")

// panicNotFallback check exist default value,
// if not exist, this variable is Required
func panicNotFallback(key string, defaultValueLen int) {
	if defaultValueLen == 0 {
		panicAndMassage(key)
	}
}
func panicAndMassage(key string) {
	log.Printf("Err Environment Variable <%s> is Required and not have value.", key)
	panic(ErrRequiredVariable)
}

// Get returns the string value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func Get(key string, defaultValue ...string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}
	return val
}

// GetInt returns the int value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetInt(key string, defaultValue ...int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Warning, Environment Variable <%s> must be INT, value=%s", key, val)
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}

	return intVal
}

// GetInt64 returns the int64 value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetInt64(key string, defaultValue ...int64) int64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}

	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Printf("Warning, Environment Variable <%s> must be INT, value=%s", key, val)
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}

	return intVal
}

// GetFloat returns the float64 value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetFloat(key string, defaultValue ...float64) float64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}

	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Printf("Warning, Environment Variable <%s> must be FLOAT, value=%s", key, val)
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}
	return floatVal
}

// GetBool returns the boolean value of the environment variable, or a default
// value if the environment variable is not defined
func GetBool(key string, defaultValue ...bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		panicNotFallback(key, len(defaultValue))
		return defaultValue[0]
	}

	valBool, err := strconv.ParseBool(val)
	if err != nil {
		log.Printf("Warning, Environment Variable <%s> must be Bool, value=%s", key, val)
		return defaultValue[0]
	}
	return valBool
}

// GetArray returns the []string of the environment variable, or a default
// value if the environment variable is not defined
func GetArray(key string, defaultValue []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		if defaultValue != nil {
			return defaultValue
		} else {
			// defaultValue==nil, is Required Variable
			panicAndMassage(key)
		}
	}

	if len(val) == 0 {
		if defaultValue != nil {
			return defaultValue
		} else {
			// defaultValue==nil, is Required Variable
			panicAndMassage(key)
		}
	}

	return strings.Split(val, ",")
}
