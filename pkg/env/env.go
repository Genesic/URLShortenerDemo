package env

import (
	"fmt"
	"os"
	"strconv"
)

func Get(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Errorf("key(%s) not found", key))
	}

	return val
}

func GetDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetDefaultInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Errorf("[env] parse int failed key: %s, val:%s, err: %s", key, val, err))
	}

	return v
}