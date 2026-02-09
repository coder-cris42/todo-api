package utils

import "os"

func GetEnvironmentVariable(key string) string {

	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value

}
