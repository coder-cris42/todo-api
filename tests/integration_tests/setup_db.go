package integrationtests

import (
	"context"
	"database/sql"
)

// The functions below map 1:1 to SQL statements found in scripts/sql/schema.sql.
// Each function accepts a context and *sql.DB and executes the preserved SQL.

func DropTodoDatabase(ctx context.Context, db *sql.DB) error {
	query := "DROP DATABASE IF EXISTS `todo-database`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func CreateTodoDatabase(ctx context.Context, db *sql.DB) error {
	query := "CREATE DATABASE `todo-database` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func UseTodoDatabase(ctx context.Context, db *sql.DB) error {
	query := "USE `todo-database`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func DropUsersTable(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS `users`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func CreateUsersTable(ctx context.Context, db *sql.DB) error {
	query := "CREATE TABLE `users` (\n" +
		"    id BIGINT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    name VARCHAR(255) NOT NULL,\n" +
		"    username VARCHAR(100) NOT NULL UNIQUE,\n" +
		"    email VARCHAR(255) NOT NULL UNIQUE,\n" +
		"    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,\n" +
		"    INDEX idx_username (username),\n" +
		"    INDEX idx_email (email)\n" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func DropTaskStatusesTable(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS `task_statuses`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func CreateTaskStatusesTable(ctx context.Context, db *sql.DB) error {
	query := "CREATE TABLE `task_statuses` (\n" +
		"    id BIGINT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    label VARCHAR(100) NOT NULL UNIQUE,\n" +
		"    active BOOLEAN NOT NULL DEFAULT true,\n" +
		"    INDEX idx_active (active)\n" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func DropTaskTypesTable(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS `task_types`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func CreateTaskTypesTable(ctx context.Context, db *sql.DB) error {
	query := "CREATE TABLE `task_types` (\n" +
		"    id BIGINT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    name VARCHAR(100) NOT NULL UNIQUE\n" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func DropWorkflowsTable(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS `workflows`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func CreateWorkflowsTable(ctx context.Context, db *sql.DB) error {
	query := "CREATE TABLE `workflows` (\n" +
		"    id BIGINT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    name VARCHAR(255) NOT NULL,\n" +
		"    statuses JSON NOT NULL,\n" +
		"    author_id BIGINT NOT NULL,\n" +
		"    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,\n" +
		"    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE,\n" +
		"    INDEX idx_author_id (author_id)\n" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func DropTasksTable(ctx context.Context, db *sql.DB) error {
	query := "DROP TABLE IF EXISTS `tasks`;"
	_, err := db.ExecContext(ctx, query)
	return err
}

func CreateTasksTable(ctx context.Context, db *sql.DB) error {
	query := "CREATE TABLE `tasks` (\n" +
		"    id BIGINT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    title VARCHAR(255) NOT NULL,\n" +
		"    description LONGTEXT,\n" +
		"    status_id BIGINT NOT NULL,\n" +
		"    parent_id BIGINT NULL,\n" +
		"    author_id BIGINT NOT NULL,\n" +
		"    deadline DATETIME NOT NULL,\n" +
		"    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,\n" +
		"    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n" +
		"    responsible_id BIGINT NULL,\n" +
		"    workflow_id BIGINT NOT NULL,\n" +
		"    type_id BIGINT NOT NULL,\n" +
		"    completed BOOLEAN NOT NULL DEFAULT false,\n" +
		"    \n" +
		"    FOREIGN KEY (status_id) REFERENCES task_statuses(id) ON DELETE RESTRICT ON UPDATE CASCADE,\n" +
		"    FOREIGN KEY (parent_id) REFERENCES tasks(id) ON DELETE CASCADE ON UPDATE CASCADE,\n" +
		"    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE,\n" +
		"    FOREIGN KEY (responsible_id) REFERENCES users(id) ON DELETE SET NULL ON UPDATE CASCADE,\n" +
		"    FOREIGN KEY (workflow_id) REFERENCES workflows(id) ON DELETE RESTRICT ON UPDATE CASCADE,\n" +
		"    FOREIGN KEY (type_id) REFERENCES task_types(id) ON DELETE RESTRICT ON UPDATE CASCADE,\n" +
		"    \n" +
		"    INDEX idx_status_id (status_id),\n" +
		"    INDEX idx_parent_id (parent_id),\n" +
		"    INDEX idx_author_id (author_id),\n" +
		"    INDEX idx_responsible_id (responsible_id),\n" +
		"    INDEX idx_workflow_id (workflow_id),\n" +
		"    INDEX idx_type_id (type_id),\n" +
		"    INDEX idx_deadline (deadline),\n" +
		"    INDEX idx_completed (completed),\n" +
		"    INDEX idx_created_at (created_at),\n" +
		"    INDEX idx_author_status (author_id, status_id),\n" +
		"    INDEX idx_responsible_status (responsible_id, status_id),\n" +
		"    INDEX idx_overdue (completed, deadline)\n" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, query)
	return err
}

// Sample data inserts
func InsertSampleUsers(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO users (name, username, email) VALUES\n" +
		"('John Doe', 'johndoe', 'john@example.com'),\n" +
		"('Jane Smith', 'janesmith', 'jane@example.com'),\n" +
		"('Bob Johnson', 'bjohnson', 'bob@example.com'),\n" +
		"('Alice Williams', 'awilliams', 'alice@example.com'),\n" +
		"('Carlos Rodriguez', 'crodriguez', 'carlos@example.com'),\n" +
		"('Emma Thompson', 'ethompson', 'emma@example.com');"
	return db.ExecContext(ctx, query)
}

func InsertSampleTaskStatuses(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO task_statuses (label, active) VALUES\n" +
		"('Backlog', true),\n" +
		"('Todo', true),\n" +
		"('In Progress', true),\n" +
		"('In Review', true),\n" +
		"('Done', true),\n" +
		"('Cancelled', true),\n" +
		"('Blocked', true),\n" +
		"('On Hold', true);"
	return db.ExecContext(ctx, query)
}

func InsertSampleTaskTypes(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO task_types (name) VALUES\n" +
		"('Bug'),\n" +
		"('Feature'),\n" +
		"('Enhancement'),\n" +
		"('Documentation'),\n" +
		"('Research'),\n" +
		"('Testing'),\n" +
		"('Refactoring');"
	return db.ExecContext(ctx, query)
}

func InsertWorkflowDefault(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO workflows (name, statuses, author_id) VALUES\n" +
		"('Default Workflow', '{\n  \"0\": {\"id\": 1, \"label\": \"Backlog\", \"active\": true},\n  \"1\": {\"id\": 2, \"label\": \"Todo\", \"active\": true},\n  \"2\": {\"id\": 3, \"label\": \"In Progress\", \"active\": true},\n  \"3\": {\"id\": 4, \"label\": \"In Review\", \"active\": true},\n  \"4\": {\"id\": 5, \"label\": \"Done\", \"active\": true}\n}', 1);"
	return db.ExecContext(ctx, query)
}

func InsertWorkflowAgile(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO workflows (name, statuses, author_id) VALUES\n" +
		"('Agile Development', '{\n  \"0\": {\"id\": 1, \"label\": \"Backlog\", \"active\": true},\n  \"1\": {\"id\": 2, \"label\": \"Todo\", \"active\": true},\n  \"2\": {\"id\": 3, \"label\": \"In Progress\", \"active\": true},\n  \"3\": {\"id\": 4, \"label\": \"In Review\", \"active\": true},\n  \"4\": {\"id\": 5, \"label\": \"Done\", \"active\": true}\n}', 2);"
	return db.ExecContext(ctx, query)
}

func InsertWorkflowSupport(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO workflows (name, statuses, author_id) VALUES\n" +
		"('Support Ticket Workflow', '{\n  \"0\": {\"id\": 1, \"label\": \"Backlog\", \"active\": true},\n  \"1\": {\"id\": 2, \"label\": \"Todo\", \"active\": true},\n  \"2\": {\"id\": 3, \"label\": \"In Progress\", \"active\": true},\n  \"3\": {\"id\": 7, \"label\": \"Blocked\", \"active\": true},\n  \"4\": {\"id\": 5, \"label\": \"Done\", \"active\": true}\n}', 3);"
	return db.ExecContext(ctx, query)
}

// Individual task inserts (one function per INSERT statement)
func InsertTask1(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Implement User Authentication',\n" +
		"  'Create a complete user authentication system with JWT tokens, password hashing, and session management. Must include login, logout, and refresh token functionality.',\n" +
		"  3,\n" +
		"  1,\n" +
		"  '2026-02-28',\n" +
		"  2,\n" +
		"  1,\n" +
		"  2,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask2(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Fix Login Bug on Mobile Devices',\n" +
		"  'Users are reporting that login is failing on mobile browsers. The issue seems to be related to cookie handling on iOS Safari.',\n" +
		"  3,\n" +
		"  1,\n" +
		"  '2026-02-15',\n" +
		"  3,\n" +
		"  1,\n" +
		"  1,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask3(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Optimize Database Queries',\n" +
		"  'Review and optimize slow database queries. Add appropriate indexes and consider query refactoring for better performance.',\n" +
		"  2,\n" +
		"  2,\n" +
		"  '2026-03-10',\n" +
		"  4,\n" +
		"  1,\n" +
		"  4,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask4(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Write API Documentation',\n" +
		"  'Create comprehensive API documentation including endpoint descriptions, request/response examples, and authentication requirements.',\n" +
		"  2,\n" +
		"  3,\n" +
		"  '2026-02-20',\n" +
		"  5,\n" +
		"  2,\n" +
		"  4,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask5(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Setup Development Environment',\n" +
		"  'Configure and document the complete development environment setup for new team members.',\n" +
		"  5,\n" +
		"  1,\n" +
		"  '2026-01-30',\n" +
		"  2,\n" +
		"  1,\n" +
		"  4,\n" +
		"  true\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask6(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Implement Email Notifications',\n" +
		"  'Add email notification system for task assignments, deadline reminders, and status updates. Should support multiple email templates.',\n" +
		"  2,\n" +
		"  2,\n" +
		"  '2026-03-05',\n" +
		"  6,\n" +
		"  1,\n" +
		"  2,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask7(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Implement JWT Token Generation',\n" +
		"  'Create JWT token generation and validation logic. Include expiration handling and refresh token mechanism.',\n" +
		"  3,\n" +
		"  1,\n" +
		"  1,\n" +
		"  '2026-02-20',\n" +
		"  2,\n" +
		"  1,\n" +
		"  2,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask8(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Implement Password Hashing',\n" +
		"  'Set up secure password hashing using bcrypt. Create password validation and reset functionality.',\n" +
		"  4,\n" +
		"  1,\n" +
		"  1,\n" +
		"  '2026-02-25',\n" +
		"  2,\n" +
		"  1,\n" +
		"  2,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask9(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Research API Caching Strategies',\n" +
		"  'Research and evaluate different caching strategies (Redis, Memcached, etc.) for API responses to improve performance.',\n" +
		"  1,\n" +
		"  3,\n" +
		"  '2026-02-28',\n" +
		"  4,\n" +
		"  1,\n" +
		"  5,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask10(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Implement Payment Processing',\n" +
		"  'Integrate payment processing gateway (Stripe/PayPal). Blocked waiting for business requirements clarification.',\n" +
		"  7,\n" +
		"  2,\n" +
		"  '2026-03-15',\n" +
		"  5,\n" +
		"  3,\n" +
		"  2,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask11(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Unit Tests for Authentication Module',\n" +
		"  'Write comprehensive unit tests for the authentication module covering all edge cases and error scenarios.',\n" +
		"  2,\n" +
		"  1,\n" +
		"  '2026-02-22',\n" +
		"  3,\n" +
		"  1,\n" +
		"  6,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}

func InsertTask12(ctx context.Context, db *sql.DB) (sql.Result, error) {
	query := "INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)\n" +
		"VALUES (\n" +
		"  'Refactor Task Repository Layer',\n" +
		"  'Refactor the task repository to improve code organization and reduce duplication. Consider implementing repository pattern.',\n" +
		"  2,\n" +
		"  3,\n" +
		"  '2026-03-01',\n" +
		"  2,\n" +
		"  1,\n" +
		"  7,\n" +
		"  false\n" +
		");"
	return db.ExecContext(ctx, query)
}
