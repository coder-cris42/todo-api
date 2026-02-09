# SQLite Integration Tests Configuration

This document explains how to use SQLite in-memory database for running integration tests in the Todo API project.

## Overview

SQLite is a lightweight, file-based SQL database that can run entirely in-memory. This makes it ideal for integration tests because it:

- **No external dependencies**: No need to run a separate MySQL server
- **Fast**: In-memory databases are extremely fast
- **Isolated**: Each test can have its own independent database
- **Easy cleanup**: Database is automatically cleaned up when connection closes

## Files Created

### 1. **sqlite_conn.go**
Location: `internal/infrastructure/database/connection/sqlite_conn.go`

Provides two functions for creating SQLite database connections:

```go
// In-memory SQLite database (recommended for tests)
NewSQLiteInMemoryDB() (*sql.DB, error)

// File-based SQLite database (for persistent testing or debugging)
NewSQLiteFileDB(filePath string) (*sql.DB, error)
```

### 2. **setup_db_sqlite.go**
Location: `tests/integration_tests/setup_db_sqlite.go`

Contains SQLite-compatible versions of all database setup functions:

- `CreateUsersTableSQLite()` - Creates users table
- `CreateTaskStatusesTableSQLite()` - Creates task_statuses table
- `CreateTaskTypesTableSQLite()` - Creates task_types table
- `CreateWorkflowsTableSQLite()` - Creates workflows table
- `CreateTasksTableSQLite()` - Creates tasks table
- `InsertSampleUsersSQLite()` - Inserts 6 sample users
- `InsertSampleTaskStatusesSQLite()` - Inserts 8 task statuses
- `InsertSampleTaskTypesSQLite()` - Inserts 7 task types
- `InsertWorkflowDefaultSQLite()`, `InsertWorkflowAgileSQLite()`, `InsertWorkflowSupportSQLite()` - Insert sample workflows
- `InsertTask1SQLite()` through `InsertTask12SQLite()` - Insert 12 sample tasks
- `CleanupTablesSQLite()` - Drops all tables for cleanup

### 3. **setup_tests.go** (Updated)
Location: `tests/integration_tests/setup_tests.go`

Now includes two setup functions:

- `SetupTests(db *sql.DB)` - Original MySQL setup (still available)
- `SetupTestsSQLite(db *sql.DB) error` - New SQLite setup

### 4. **test_helpers.go**
Location: `tests/integration_tests/test_helpers.go`

Provides convenient helper functions for test database management:

```go
// Initialize test database with configuration
InitializeTestDatabase(config TestDatabaseConfig) (*sql.DB, error)

// Clean up and close test database
CleanupTestDatabase(db *sql.DB, dbType string) error

// Reset database to initial state (drop and recreate tables)
ResetTestDatabase(db *sql.DB, dbType string) error

// Get connection string for test database
GetTestDatabaseConnectionString(config TestDatabaseConfig) string
```

### 5. **example_test.go**
Location: `tests/integration_tests/example_test.go`

Provides example test cases demonstrating how to use the SQLite configuration.

## Usage Examples

### Basic Setup (In-Memory SQLite)

```go
package integrationtests

import (
    "context"
    "testing"
)

func TestUserCreation(t *testing.T) {
    // Initialize SQLite in-memory database with sample data
    db, err := InitializeTestDatabase(TestDatabaseConfig{
        Type: "sqlite",
        Path: "", // empty path means in-memory
    })
    if err != nil {
        t.Fatalf("Failed to initialize test database: %v", err)
    }
    defer CleanupTestDatabase(db, "sqlite")

    // Your test code here
    ctx := context.Background()
    var count int
    db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
    
    if count != 6 {
        t.Errorf("Expected 6 users, got %d", count)
    }
}
```

### File-Based SQLite (for debugging)

```go
func TestWithFileBased(t *testing.T) {
    db, err := InitializeTestDatabase(TestDatabaseConfig{
        Type: "sqlite",
        Path: "/tmp/test_todo.db",
    })
    if err != nil {
        t.Fatalf("Failed to initialize test database: %v", err)
    }
    defer CleanupTestDatabase(db, "sqlite")
    
    // Your test code here
}
```

### Database Reset Between Tests

```go
func TestWithReset(t *testing.T) {
    db, err := InitializeTestDatabase(TestDatabaseConfig{
        Type: "sqlite",
        Path: "",
    })
    if err != nil {
        t.Fatalf("Failed to initialize test database: %v", err)
    }
    defer CleanupTestDatabase(db, "sqlite")

    // First test
    // ... test code ...

    // Reset database to clean state
    if err := ResetTestDatabase(db, "sqlite"); err != nil {
        t.Fatalf("Failed to reset database: %v", err)
    }

    // Second test with fresh database
    // ... test code ...
}
```

### Direct Function Usage

If you prefer more control, you can use the setup functions directly:

```go
func TestDirect(t *testing.T) {
    db, err := connection.NewSQLiteInMemoryDB()
    if err != nil {
        t.Fatalf("Failed to create database: %v", err)
    }
    defer db.Close()

    // Setup database schema and sample data
    if err := SetupTestsSQLite(db); err != nil {
        t.Fatalf("Failed to setup database: %v", err)
    }

    // Your test code here

    // Cleanup (optional, in-memory will be cleaned up on close)
    ctx := context.Background()
    CleanupTablesSQLite(ctx, db)
}
```

## Database Schema and Sample Data

### Tables Created:
1. **users** - 6 sample users
2. **task_statuses** - 8 statuses (Backlog, Todo, In Progress, In Review, Done, Cancelled, Blocked, On Hold)
3. **task_types** - 7 types (Bug, Feature, Enhancement, Documentation, Research, Testing, Refactoring)
4. **workflows** - 3 workflows (Default, Agile Development, Support Ticket)
5. **tasks** - 12 sample tasks with various statuses and relationships

### Sample Data:
- **6 Users**: John Doe, Jane Smith, Bob Johnson, Alice Williams, Carlos Rodriguez, Emma Thompson
- **12 Tasks**: Including authentication tasks, bug fixes, database optimization, documentation, etc.
- **Foreign Key Relationships**: All relationships are preserved from the original MySQL schema

## Running Integration Tests

### Run all integration tests:
```bash
go test ./tests/integration_tests -v
```

### Run specific test:
```bash
go test ./tests/integration_tests -v -run TestExampleWithSQLiteInMemory
```

### Run with coverage:
```bash
go test ./tests/integration_tests -v -cover
```

## Environment Setup

### 1. Add SQLite driver to go.mod (already done):
The `github.com/mattn/go-sqlite3` package has been added to your `go.mod` file.

### 2. Download dependencies:
```bash
go mod download
go mod tidy
```

### 3. Install SQLite3 development headers (if needed):

**Ubuntu/Debian:**
```bash
sudo apt-get install sqlite3 libsqlite3-dev
```

**macOS:**
```bash
brew install sqlite3
```

**Windows:**
SQLite3 is typically bundled with the Go sqlite3 driver.

## Key Differences Between MySQL and SQLite

### 1. **Database Operations**
- MySQL: `CREATE DATABASE`, `USE DATABASE`
- SQLite: No database creation needed (uses tables directly)

### 2. **Data Types**
- MySQL: `BOOLEAN` → SQLite: `BOOLEAN` (stored as 0/1 or true/false)
- MySQL: `BIGINT AUTO_INCREMENT` → SQLite: `INTEGER PRIMARY KEY AUTOINCREMENT`
- MySQL: `DATETIME` → SQLite: `DATETIME` (text format)
- MySQL: `LONGTEXT` → SQLite: `TEXT` (no length limit)

### 3. **JSON Handling**
- MySQL: Native JSON type
- SQLite: Stores JSON as TEXT

### 4. **Index Creation**
- MySQL: Inline with CREATE TABLE
- SQLite: Can be created inline or with separate CREATE INDEX

## Troubleshooting

### Issue: "database is locked"
**Solution**: SQLite in-memory is single-threaded. If running parallel tests, use separate databases or file-based SQLite.

### Issue: "no such table"
**Solution**: Ensure `SetupTestsSQLite()` was called before querying. Check error handling on initialization.

### Issue: "FOREIGN KEY constraint failed"
**Solution**: SQLite requires explicit enablement of foreign keys. The code handles this, but ensure you're using the provided setup functions.

### Issue: Import cycle error
**Solution**: Make sure the package structure is correct and you're importing `connection` package properly.

## Best Practices

1. **Always use defer for cleanup**:
   ```go
   defer CleanupTestDatabase(db, "sqlite")
   ```

2. **Use InitializeTestDatabase for simplicity**:
   ```go
   db, err := InitializeTestDatabase(TestDatabaseConfig{
       Type: "sqlite",
       Path: "",
   })
   ```

3. **Reset database for test isolation**:
   ```go
   if err := ResetTestDatabase(db, "sqlite"); err != nil {
       t.Fatal(err)
   }
   ```

4. **Test with both in-memory and file-based** when appropriate:
   - In-memory: For fast unit-like tests
   - File-based: For debugging database state

5. **Use context for all database operations**:
   ```go
   ctx := context.Background()
   db.QueryContext(ctx, query)
   ```

## Migrating Existing Tests

If you have existing MySQL integration tests, you can migrate them by:

1. Replacing MySQL initialization with SQLite:
   ```go
   // Old
   db, _ := connection.NewMySQLDB(...)
   
   // New
   db, _ := InitializeTestDatabase(TestDatabaseConfig{
       Type: "sqlite",
       Path: "",
   })
   ```

2. Update any MySQL-specific SQL if needed (mostly handled by the provided functions)

3. Add cleanup defer statement:
   ```go
   defer CleanupTestDatabase(db, "sqlite")
   ```

## References

- [SQLite Documentation](https://www.sqlite.org/docs.html)
- [Go SQLite3 Driver](https://github.com/mattn/go-sqlite3)
- [database/sql Package](https://golang.org/pkg/database/sql/)
