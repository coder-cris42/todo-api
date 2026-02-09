package integrationtests

import (
	"context"
	"database/sql"
	"fmt"

	"todo-api/internal/infrastructure/database/connection"
)

// TestDatabaseConfig holds configuration for test database initialization
type TestDatabaseConfig struct {
	Type string // "sqlite" or "mysql"
	Path string // for sqlite file-based databases
}

// InitializeTestDatabase creates and initializes a test database
// For SQLite in-memory, use Type: "sqlite" and leave Path empty
// For SQLite file-based, use Type: "sqlite" and set Path to the file location
// For MySQL, use Type: "mysql"
func InitializeTestDatabase(config TestDatabaseConfig) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch config.Type {
	case "sqlite":
		if config.Path == "" {
			// In-memory SQLite
			db, err = connection.NewSQLiteInMemoryDB()
		} else {
			// File-based SQLite
			db, err = connection.NewSQLiteFileDB(config.Path)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite database: %w", err)
		}

		// Setup the database schema and sample data
		if err := SetupTestsSQLite(db); err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to setup sqlite database: %w", err)
		}

	case "mysql":
		// For MySQL, you need to provide connection parameters via environment variables
		// DB_USER, DB_PASSWORD, DB_HOST, DB_PORT
		username := "root"
		password := ""
		host := "localhost"
		port := "3306"
		dbName := ""

		db, err = connection.NewMySQLDB(username, password, host, port, dbName)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql database: %w", err)
		}

		// Setup the database schema and sample data
		SetupTests(db)

	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}

	return db, nil
}

// CleanupTestDatabase cleans up the test database
func CleanupTestDatabase(db *sql.DB, dbType string) error {
	if db == nil {
		return nil
	}

	ctx := context.Background()

	if dbType == "sqlite" {
		if err := CleanupTablesSQLite(ctx, db); err != nil {
			return fmt.Errorf("failed to cleanup sqlite database: %w", err)
		}
	}

	return db.Close()
}

// ResetTestDatabase drops and recreates all tables (useful for test isolation)
func ResetTestDatabase(db *sql.DB, dbType string) error {
	ctx := context.Background()

	if dbType == "sqlite" {
		if err := CleanupTablesSQLite(ctx, db); err != nil {
			return err
		}
		return SetupTestsSQLite(db)
	}

	// For MySQL, drop and recreate database
	if dbType == "mysql" {
		if err := DropTodoDatabase(ctx, db); err != nil {
			return err
		}
		if err := CreateTodoDatabase(ctx, db); err != nil {
			return err
		}
		if err := UseTodoDatabase(ctx, db); err != nil {
			return err
		}
		SetupTests(db)
		return nil
	}

	return fmt.Errorf("unsupported database type: %s", dbType)
}

// GetTestDatabaseConnectionString returns a DSN string for the test database
func GetTestDatabaseConnectionString(config TestDatabaseConfig) string {
	if config.Type == "sqlite" {
		if config.Path == "" {
			return ":memory:"
		}
		return fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.Path)
	}
	return ""
}
