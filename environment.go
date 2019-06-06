package tools

import (
	"syscall"
	"strconv"
)

func GetEnvString(key, defaultValue string) string {
	if value, found := syscall.Getenv(key); !found {
		return defaultValue
	} else {
		return value
	}
}

func GetEnvBool(key string, defaultValue bool) bool {
	if value, found := syscall.Getenv(key); !found {
		return defaultValue
	} else {
		return value == "true"
	}
}

func GetEnvNumber(key string, defaultValue int) (int, error) {
	if value, found := syscall.Getenv(key); !found {
		return defaultValue, nil
	} else {
		return strconv.Atoi(value)
	}
}
