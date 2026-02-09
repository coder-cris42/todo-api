package integrationtests

import (
	"context"
	"database/sql"
)

// SetupTests sets up the test database using MySQL
// This is the original setup function for MySQL databases
func SetupTests(db *sql.DB) {
	ctx := context.Background()

	DropTodoDatabase(ctx, db)
	CreateTodoDatabase(ctx, db)
	UseTodoDatabase(ctx, db)
	CreateUsersTable(ctx, db)
	CreateTaskStatusesTable(ctx, db)
	CreateTaskTypesTable(ctx, db)
	CreateWorkflowsTable(ctx, db)
	CreateTasksTable(ctx, db)
	InsertSampleUsers(ctx, db)
	InsertSampleTaskStatuses(ctx, db)
	InsertSampleTaskTypes(ctx, db)
	InsertWorkflowDefault(ctx, db)
	InsertWorkflowAgile(ctx, db)
	InsertWorkflowSupport(ctx, db)
	InsertTask1(ctx, db)
	InsertTask2(ctx, db)
	InsertTask3(ctx, db)
	InsertTask4(ctx, db)
	InsertTask5(ctx, db)
	InsertTask6(ctx, db)
	InsertTask7(ctx, db)
	InsertTask8(ctx, db)
	InsertTask9(ctx, db)
	InsertTask10(ctx, db)
	InsertTask11(ctx, db)
	InsertTask12(ctx, db)
}

// SetupTestsSQLite sets up the test database using SQLite in-memory
// This is the new setup function for SQLite in-memory databases, ideal for integration tests
func SetupTestsSQLite(db *sql.DB) error {
	ctx := context.Background()

	// Drop existing tables (SQLite doesn't have DROP DATABASE)
	if err := CleanupTablesSQLite(ctx, db); err != nil {
		return err
	}

	// Create all tables
	if err := CreateUsersTableSQLite(ctx, db); err != nil {
		return err
	}
	if err := CreateTaskStatusesTableSQLite(ctx, db); err != nil {
		return err
	}
	if err := CreateTaskTypesTableSQLite(ctx, db); err != nil {
		return err
	}
	if err := CreateWorkflowsTableSQLite(ctx, db); err != nil {
		return err
	}
	if err := CreateTasksTableSQLite(ctx, db); err != nil {
		return err
	}

	// Insert sample data
	if _, err := InsertSampleUsersSQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertSampleTaskStatusesSQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertSampleTaskTypesSQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertWorkflowDefaultSQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertWorkflowAgileSQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertWorkflowSupportSQLite(ctx, db); err != nil {
		return err
	}

	// Insert sample tasks
	if _, err := InsertTask1SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask2SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask3SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask4SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask5SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask6SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask7SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask8SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask9SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask10SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask11SQLite(ctx, db); err != nil {
		return err
	}
	if _, err := InsertTask12SQLite(ctx, db); err != nil {
		return err
	}

	return nil
}
