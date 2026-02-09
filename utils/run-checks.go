package utils

import (
	"fmt"
	"os"
	conn "todo-api/internal/infrastructure/database/connection"
)

func CheckEnvironmentVariables() error {

	if getEnvironmentVariable("DB_USERNAME") == "" {
		return fmt.Errorf("DB_USERNAME environment variable is not set")
	}

	if getEnvironmentVariable("DB_PASSWORD") == "" {
		return fmt.Errorf("DB_PASSWORD environment variable is not set")
	}

	if getEnvironmentVariable("DB_HOST") == "" {
		return fmt.Errorf("DB_HOST environment variable is not set")
	}

	if getEnvironmentVariable("DB_PORT") == "" {
		return fmt.Errorf("DB_PORT environment variable is not set")
	}

	if getEnvironmentVariable("DB_NAME") == "" {
		return fmt.Errorf("DB_NAME environment variable is not set")
	}

	return nil
}

func CheckDatabaseConnection(username, password, host, port, dbName string) error {

	db, err := conn.NewMySQLDB(username, password, host, port, dbName)
	if err != nil {
		return err
	}

	err = conn.CheckConnectivity(db)
	if err != nil {
		return err
	}

	return nil

}

func getEnvironmentVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}
