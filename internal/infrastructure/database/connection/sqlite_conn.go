package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// NewSQLiteInMemoryDB creates a new SQLite in-memory database connection.
// This is useful for integration tests where you want fast, isolated test databases.
func NewSQLiteInMemoryDB() (*sql.DB, error) {
	// Use simple :memory: for in-memory SQLite database
	// This will keep the database in memory for as long as the connection exists
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		db.Close()
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}

	// Pool settings
	db.SetMaxOpenConns(1) // SQLite in-memory must use single connection
	db.SetMaxIdleConns(1) // Keep connection alive
	db.SetConnMaxLifetime(time.Hour)

	// Verify connectivity
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return db, nil
}

// NewSQLiteFileDB creates a new SQLite file-based database connection.
// Useful for persistent integration tests or debugging.
func NewSQLiteFileDB(filePath string) (*sql.DB, error) {
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", filePath)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		db.Close()
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}

	// Pool settings
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	// Verify connectivity
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return db, nil
}
