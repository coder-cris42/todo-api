package integrationtests

import (
	"context"
	"database/sql"
)

// SQLite versions of database setup functions
// SQLite doesn't support CREATE DATABASE, so we skip those functions and use the tables directly

// CreateUsersTableSQLite creates the users table compatible with SQLite
func CreateUsersTableSQLite(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := db.ExecContext(ctx, query); err != nil {
		return err
	}

	indexQuery := `CREATE INDEX IF NOT EXISTS idx_username ON users(username);`
	if _, err := db.ExecContext(ctx, indexQuery); err != nil {
		return err
	}

	indexQuery2 := `CREATE INDEX IF NOT EXISTS idx_email ON users(email);`
	if _, err := db.ExecContext(ctx, indexQuery2); err != nil {
		return err
	}

	return nil
}

// DropUsersTableSQLite drops the users table
func DropUsersTableSQLite(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS users;"
	_, err := db.ExecContext(ctx, query)
	return err
}

// CreateTaskStatusesTableSQLite creates the task_statuses table compatible with SQLite
func CreateTaskStatusesTableSQLite(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS task_statuses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		label TEXT NOT NULL UNIQUE,
		active BOOLEAN NOT NULL DEFAULT 1
	);`
	if _, err := db.ExecContext(ctx, query); err != nil {
		return err
	}

	indexQuery := `CREATE INDEX IF NOT EXISTS idx_active ON task_statuses(active);`
	if _, err := db.ExecContext(ctx, indexQuery); err != nil {
		return err
	}

	return nil
}

// DropTaskStatusesTableSQLite drops the task_statuses table
func DropTaskStatusesTableSQLite(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS task_statuses;"
	_, err := db.ExecContext(ctx, query)
	return err
}

// CreateTaskTypesTableSQLite creates the task_types table compatible with SQLite
func CreateTaskTypesTableSQLite(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS task_types (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	);`
	_, err := db.ExecContext(ctx, query)
	return err
}

// DropTaskTypesTableSQLite drops the task_types table
func DropTaskTypesTableSQLite(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS task_types;"
	_, err := db.ExecContext(ctx, query)
	return err
}

// CreateWorkflowsTableSQLite creates the workflows table compatible with SQLite
func CreateWorkflowsTableSQLite(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS workflows (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		statuses TEXT NOT NULL,
		author_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE
	);`
	if _, err := db.ExecContext(ctx, query); err != nil {
		return err
	}

	indexQuery := `CREATE INDEX IF NOT EXISTS idx_author_id ON workflows(author_id);`
	if _, err := db.ExecContext(ctx, indexQuery); err != nil {
		return err
	}

	return nil
}

// DropWorkflowsTableSQLite drops the workflows table
func DropWorkflowsTableSQLite(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS workflows;"
	_, err := db.ExecContext(ctx, query)
	return err
}

// CreateTasksTableSQLite creates the tasks table compatible with SQLite
func CreateTasksTableSQLite(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status_id INTEGER NOT NULL,
		parent_id INTEGER,
		author_id INTEGER NOT NULL,
		deadline DATETIME NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		responsible_id INTEGER,
		workflow_id INTEGER NOT NULL,
		type_id INTEGER NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		FOREIGN KEY (status_id) REFERENCES task_statuses(id) ON DELETE RESTRICT ON UPDATE CASCADE,
		FOREIGN KEY (parent_id) REFERENCES tasks(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE,
		FOREIGN KEY (responsible_id) REFERENCES users(id) ON DELETE SET NULL ON UPDATE CASCADE,
		FOREIGN KEY (workflow_id) REFERENCES workflows(id) ON DELETE RESTRICT ON UPDATE CASCADE,
		FOREIGN KEY (type_id) REFERENCES task_types(id) ON DELETE RESTRICT ON UPDATE CASCADE
	);`
	if _, err := db.ExecContext(ctx, query); err != nil {
		return err
	}

	// Create all indexes separately
	indexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_status_id ON tasks(status_id);`,
		`CREATE INDEX IF NOT EXISTS idx_parent_id ON tasks(parent_id);`,
		`CREATE INDEX IF NOT EXISTS idx_author_id ON tasks(author_id);`,
		`CREATE INDEX IF NOT EXISTS idx_responsible_id ON tasks(responsible_id);`,
		`CREATE INDEX IF NOT EXISTS idx_workflow_id ON tasks(workflow_id);`,
		`CREATE INDEX IF NOT EXISTS idx_type_id ON tasks(type_id);`,
		`CREATE INDEX IF NOT EXISTS idx_deadline ON tasks(deadline);`,
		`CREATE INDEX IF NOT EXISTS idx_completed ON tasks(completed);`,
		`CREATE INDEX IF NOT EXISTS idx_created_at ON tasks(created_at);`,
		`CREATE INDEX IF NOT EXISTS idx_author_status ON tasks(author_id, status_id);`,
		`CREATE INDEX IF NOT EXISTS idx_responsible_status ON tasks(responsible_id, status_id);`,
		`CREATE INDEX IF NOT EXISTS idx_overdue ON tasks(completed, deadline);`,
	}

	for _, indexQuery := range indexes {
		if _, err := db.ExecContext(ctx, indexQuery); err != nil {
			return err
		}
	}

	return nil
}

// DropTasksTableSQLite drops the tasks table
func DropTasksTableSQLite(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS tasks;"
	_, err := db.ExecContext(ctx, query)
	return err
}

// Sample data inserts (compatible with SQLite)

func InsertSampleUsersSQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO users (name, username, email) VALUES
		('John Doe', 'johndoe', 'john@example.com'),
		('Jane Smith', 'janesmith', 'jane@example.com'),
		('Bob Johnson', 'bjohnson', 'bob@example.com'),
		('Alice Williams', 'awilliams', 'alice@example.com'),
		('Carlos Rodriguez', 'crodriguez', 'carlos@example.com'),
		('Emma Thompson', 'ethompson', 'emma@example.com');`
	return db.ExecContext(ctx, query)
}

func InsertSampleTaskStatusesSQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO task_statuses (label, active) VALUES
		('Backlog', 1),
		('Todo', 1),
		('In Progress', 1),
		('In Review', 1),
		('Done', 1),
		('Cancelled', 1),
		('Blocked', 1),
		('On Hold', 1);`
	return db.ExecContext(ctx, query)
}

func InsertSampleTaskTypesSQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO task_types (name) VALUES
		('Bug'),
		('Feature'),
		('Enhancement'),
		('Documentation'),
		('Research'),
		('Testing'),
		('Refactoring');`
	return db.ExecContext(ctx, query)
}

func InsertWorkflowDefaultSQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO workflows (name, statuses, author_id) VALUES
		('Default Workflow', '{"0": {"id": 1, "label": "Backlog", "active": true}, "1": {"id": 2, "label": "Todo", "active": true}, "2": {"id": 3, "label": "In Progress", "active": true}, "3": {"id": 4, "label": "In Review", "active": true}, "4": {"id": 5, "label": "Done", "active": true}}', 1);`
	return db.ExecContext(ctx, query)
}

func InsertWorkflowAgileSQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO workflows (name, statuses, author_id) VALUES
		('Agile Development', '{"0": {"id": 1, "label": "Backlog", "active": true}, "1": {"id": 2, "label": "Todo", "active": true}, "2": {"id": 3, "label": "In Progress", "active": true}, "3": {"id": 4, "label": "In Review", "active": true}, "4": {"id": 5, "label": "Done", "active": true}}', 2);`
	return db.ExecContext(ctx, query)
}

func InsertWorkflowSupportSQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO workflows (name, statuses, author_id) VALUES
		('Support Ticket Workflow', '{"0": {"id": 1, "label": "Backlog", "active": true}, "1": {"id": 2, "label": "Todo", "active": true}, "2": {"id": 3, "label": "In Progress", "active": true}, "3": {"id": 7, "label": "Blocked", "active": true}, "4": {"id": 5, "label": "Done", "active": true}}', 3);`
	return db.ExecContext(ctx, query)
}

// Task inserts (compatible with SQLite)

func InsertTask1SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Implement User Authentication',
			'Create a complete user authentication system with JWT tokens, password hashing, and session management. Must include login, logout, and refresh token functionality.',
			3,
			1,
			'2026-02-28',
			2,
			1,
			2,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask2SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Fix Login Bug on Mobile Devices',
			'Users are reporting that login is failing on mobile browsers. The issue seems to be related to cookie handling on iOS Safari.',
			3,
			1,
			'2026-02-15',
			3,
			1,
			1,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask3SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Optimize Database Queries',
			'Review and optimize slow database queries. Add appropriate indexes and consider query refactoring for better performance.',
			2,
			2,
			'2026-03-10',
			4,
			1,
			4,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask4SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Write API Documentation',
			'Create comprehensive API documentation including endpoint descriptions, request/response examples, and authentication requirements.',
			2,
			3,
			'2026-02-20',
			5,
			2,
			4,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask5SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Setup Development Environment',
			'Configure and document the complete development environment setup for new team members.',
			5,
			1,
			'2026-01-30',
			2,
			1,
			4,
			1
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask6SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Implement Email Notifications',
			'Add email notification system for task assignments, deadline reminders, and status updates. Should support multiple email templates.',
			2,
			2,
			'2026-03-05',
			6,
			1,
			2,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask7SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Implement JWT Token Generation',
			'Create JWT token generation and validation logic. Include expiration handling and refresh token mechanism.',
			3,
			1,
			1,
			'2026-02-20',
			2,
			1,
			2,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask8SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Implement Password Hashing',
			'Set up secure password hashing using bcrypt. Create password validation and reset functionality.',
			4,
			1,
			1,
			'2026-02-25',
			2,
			1,
			2,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask9SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Research API Caching Strategies',
			'Research and evaluate different caching strategies (Redis, Memcached, etc.) for API responses to improve performance.',
			1,
			3,
			'2026-02-28',
			4,
			1,
			5,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask10SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Implement Payment Processing',
			'Integrate payment processing gateway (Stripe/PayPal). Blocked waiting for business requirements clarification.',
			7,
			2,
			'2026-03-15',
			5,
			3,
			2,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask11SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Unit Tests for Authentication Module',
			'Write comprehensive unit tests for the authentication module covering all edge cases and error scenarios.',
			2,
			1,
			'2026-02-22',
			3,
			1,
			6,
			0
		);`
	return db.ExecContext(ctx, query)
}

func InsertTask12SQLite(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
		VALUES (
			'Refactor Task Repository Layer',
			'Refactor the task repository to improve code organization and reduce duplication. Consider implementing repository pattern.',
			2,
			3,
			'2026-03-01',
			2,
			1,
			7,
			0
		);`
	return db.ExecContext(ctx, query)
}

// CleanupTables drops all test tables
func CleanupTablesSQLite(ctx context.Context, db *sql.DB) error {
	tables := []string{"tasks", "workflows", "task_types", "task_statuses", "users"}
	for _, table := range tables {
		query := "DROP TABLE IF EXISTS " + table + ";"
		if _, err := db.ExecContext(ctx, query); err != nil {
			return err
		}
	}
	return nil
}
