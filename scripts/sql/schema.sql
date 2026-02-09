-- ============================================================================
-- TODO API DATABASE SCHEMA
-- ============================================================================
-- Created: 2026-02-08
-- Description: Complete database schema for Todo API application
-- Database: todo-database
-- Engine: MySQL InnoDB
-- Charset: utf8mb4
-- ============================================================================

-- ============================================================================
-- CREATE DATABASE
-- ============================================================================
DROP DATABASE IF EXISTS `todo-database`;
CREATE DATABASE `todo-database` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `todo-database`;

-- ============================================================================
-- USERS TABLE
-- ============================================================================
-- Stores user information
-- Fields:
--   id: Unique identifier (auto-increment)
--   name: Full name of the user
--   username: Unique username for login
--   email: Email address (should be unique)
--   created_at: Timestamp when user was created
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_username (username),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- TASK_STATUSES TABLE
-- ============================================================================
-- Predefined task status definitions (e.g., Todo, In Progress, Done)
-- Fields:
--   id: Unique identifier (auto-increment)
--   label: Human-readable status label
--   active: Boolean indicating if status is active/available for use
DROP TABLE IF EXISTS `task_statuses`;
CREATE TABLE `task_statuses` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    label VARCHAR(100) NOT NULL UNIQUE,
    active BOOLEAN NOT NULL DEFAULT true,
    INDEX idx_active (active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- TASK_TYPES TABLE
-- ============================================================================
-- Task type definitions (e.g., Bug, Feature, Enhancement)
-- Fields:
--   id: Unique identifier (auto-increment)
--   name: Name of the task type
DROP TABLE IF EXISTS `task_types`;
CREATE TABLE `task_types` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- WORKFLOWS TABLE
-- ============================================================================
-- Workflow definitions that define the sequence of statuses for tasks
-- Fields:
--   id: Unique identifier (auto-increment)
--   name: Name of the workflow
--   statuses: JSON field storing the mapping of status order to status IDs
--   author_id: Foreign key to the user who created the workflow
--   created_at: Timestamp when workflow was created
DROP TABLE IF EXISTS `workflows`;
CREATE TABLE `workflows` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    statuses JSON NOT NULL,
    author_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    INDEX idx_author_id (author_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- TASKS TABLE
-- ============================================================================
-- Main tasks table storing task information and relationships
-- Fields:
--   id: Unique identifier (auto-increment)
--   title: Task title
--   description: Detailed description of the task
--   status_id: Current status of the task (foreign key to task_statuses)
--   parent_id: Optional parent task for subtasks (self-referencing foreign key)
--   author_id: User who created the task (foreign key to users)
--   deadline: Task deadline/due date
--   created_at: Timestamp when task was created
--   updated_at: Timestamp of last update
--   responsible_id: User assigned to complete the task (foreign key to users)
--   workflow_id: Workflow template for this task (foreign key to workflows)
--   type_id: Type of task (foreign key to task_types)
--   completed: Boolean indicating if task is completed
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description LONGTEXT,
    status_id BIGINT NOT NULL,
    parent_id BIGINT NULL,
    author_id BIGINT NOT NULL,
    deadline DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    responsible_id BIGINT NULL,
    workflow_id BIGINT NOT NULL,
    type_id BIGINT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false,
    
    FOREIGN KEY (status_id) REFERENCES task_statuses(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES tasks(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (responsible_id) REFERENCES users(id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (workflow_id) REFERENCES workflows(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (type_id) REFERENCES task_types(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    
    INDEX idx_status_id (status_id),
    INDEX idx_parent_id (parent_id),
    INDEX idx_author_id (author_id),
    INDEX idx_responsible_id (responsible_id),
    INDEX idx_workflow_id (workflow_id),
    INDEX idx_type_id (type_id),
    INDEX idx_deadline (deadline),
    INDEX idx_completed (completed),
    INDEX idx_created_at (created_at),
    INDEX idx_author_status (author_id, status_id),
    INDEX idx_responsible_status (responsible_id, status_id),
    INDEX idx_overdue (completed, deadline)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- SAMPLE DATA FOR TESTING
-- ============================================================================

-- ============================================================================
-- INSERT USERS
-- ============================================================================
-- Users with different roles and responsibilities
INSERT INTO users (name, username, email) VALUES
('John Doe', 'johndoe', 'john@example.com'),
('Jane Smith', 'janesmith', 'jane@example.com'),
('Bob Johnson', 'bjohnson', 'bob@example.com'),
('Alice Williams', 'awilliams', 'alice@example.com'),
('Carlos Rodriguez', 'crodriguez', 'carlos@example.com'),
('Emma Thompson', 'ethompson', 'emma@example.com');

-- ============================================================================
-- INSERT TASK STATUSES
-- ============================================================================
-- Standard workflow statuses
INSERT INTO task_statuses (label, active) VALUES
('Backlog', true),
('Todo', true),
('In Progress', true),
('In Review', true),
('Done', true),
('Cancelled', true),
('Blocked', true),
('On Hold', true);

-- ============================================================================
-- INSERT TASK TYPES
-- ============================================================================
-- Different types of tasks
INSERT INTO task_types (name) VALUES
('Bug'),
('Feature'),
('Enhancement'),
('Documentation'),
('Research'),
('Testing'),
('Refactoring');

-- ============================================================================
-- INSERT WORKFLOWS
-- ============================================================================
-- Workflow 1: Default Workflow (Kanban style)
-- Mapping: 0->Backlog, 1->Todo, 2->In Progress, 3->In Review, 4->Done
INSERT INTO workflows (name, statuses, author_id) VALUES
('Default Workflow', '{
  "0": {"id": 1, "label": "Backlog", "active": true},
  "1": {"id": 2, "label": "Todo", "active": true},
  "2": {"id": 3, "label": "In Progress", "active": true},
  "3": {"id": 4, "label": "In Review", "active": true},
  "4": {"id": 5, "label": "Done", "active": true}
}', 1);

-- Workflow 2: Agile Development
-- Mapping: 0->Backlog, 1->Todo, 2->In Progress, 3->In Review, 4->Testing, 5->Done
INSERT INTO workflows (name, statuses, author_id) VALUES
('Agile Development', '{
  "0": {"id": 1, "label": "Backlog", "active": true},
  "1": {"id": 2, "label": "Todo", "active": true},
  "2": {"id": 3, "label": "In Progress", "active": true},
  "3": {"id": 4, "label": "In Review", "active": true},
  "4": {"id": 5, "label": "Done", "active": true}
}', 2);

-- Workflow 3: Support Ticket Workflow
-- Mapping: 0->Backlog, 1->Todo, 2->In Progress, 3->Blocked, 4->Done
INSERT INTO workflows (name, statuses, author_id) VALUES
('Support Ticket Workflow', '{
  "0": {"id": 1, "label": "Backlog", "active": true},
  "1": {"id": 2, "label": "Todo", "active": true},
  "2": {"id": 3, "label": "In Progress", "active": true},
  "3": {"id": 7, "label": "Blocked", "active": true},
  "4": {"id": 5, "label": "Done", "active": true}
}', 3);

-- ============================================================================
-- INSERT TASKS
-- ============================================================================
-- Task 1: Implement User Authentication
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Implement User Authentication',
  'Create a complete user authentication system with JWT tokens, password hashing, and session management. Must include login, logout, and refresh token functionality.',
  3,
  1,
  '2026-02-28',
  2,
  1,
  2,
  false
);

-- Task 2: Fix Login Bug
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Fix Login Bug on Mobile Devices',
  'Users are reporting that login is failing on mobile browsers. The issue seems to be related to cookie handling on iOS Safari.',
  3,
  1,
  '2026-02-15',
  3,
  1,
  1,
  false
);

-- Task 3: Database Optimization
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Optimize Database Queries',
  'Review and optimize slow database queries. Add appropriate indexes and consider query refactoring for better performance.',
  2,
  2,
  '2026-03-10',
  4,
  1,
  4,
  false
);

-- Task 4: API Documentation
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Write API Documentation',
  'Create comprehensive API documentation including endpoint descriptions, request/response examples, and authentication requirements.',
  2,
  3,
  '2026-02-20',
  5,
  2,
  4,
  false
);

-- Task 5: Complete Task (Done)
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Setup Development Environment',
  'Configure and document the complete development environment setup for new team members.',
  5,
  1,
  '2026-01-30',
  2,
  1,
  4,
  true
);

-- Task 6: Implement Email Notifications
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Implement Email Notifications',
  'Add email notification system for task assignments, deadline reminders, and status updates. Should support multiple email templates.',
  2,
  2,
  '2026-03-05',
  6,
  1,
  2,
  false
);

-- Task 7: Subtask of Task 1
INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
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
  false
);

-- Task 8: Subtask of Task 1
INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
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
  false
);

-- Task 9: Research Task
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Research API Caching Strategies',
  'Research and evaluate different caching strategies (Redis, Memcached, etc.) for API responses to improve performance.',
  1,
  3,
  '2026-02-28',
  4,
  1,
  5,
  false
);

-- Task 10: Blocked Task
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Implement Payment Processing',
  'Integrate payment processing gateway (Stripe/PayPal). Blocked waiting for business requirements clarification.',
  7,
  2,
  '2026-03-15',
  5,
  3,
  2,
  false
);

-- Task 11: Testing Task
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Unit Tests for Authentication Module',
  'Write comprehensive unit tests for the authentication module covering all edge cases and error scenarios.',
  2,
  1,
  '2026-02-22',
  3,
  1,
  6,
  false
);

-- Task 12: Refactoring Task
INSERT INTO tasks (title, description, status_id, author_id, deadline, responsible_id, workflow_id, type_id, completed)
VALUES (
  'Refactor Task Repository Layer',
  'Refactor the task repository to improve code organization and reduce duplication. Consider implementing repository pattern.',
  2,
  3,
  '2026-03-01',
  2,
  1,
  7,
  false
);

-- ============================================================================
-- INDEXES FOR COMMON QUERIES
-- ============================================================================
-- These indexes are already created above but listed here for reference:
-- - idx_active on task_statuses(active) - for filtering active statuses
-- - idx_author_id on workflows(author_id) - for finding workflows by author
-- - idx_status_id on tasks(status_id) - for filtering tasks by status
-- - idx_author_id on tasks(author_id) - for finding tasks by author
-- - idx_responsible_id on tasks(responsible_id) - for finding assigned tasks
-- - idx_workflow_id on tasks(workflow_id) - for finding tasks by workflow
-- - idx_type_id on tasks(type_id) - for finding tasks by type
-- - idx_deadline on tasks(deadline) - for sorting by deadline
-- - idx_completed on tasks(completed) - for filtering incomplete tasks
-- - idx_created_at on tasks(created_at) - for sorting by creation time
-- - idx_author_status on tasks(author_id, status_id) - composite for author+status queries
-- - idx_responsible_status on tasks(responsible_id, status_id) - composite for responsible+status
-- - idx_overdue on tasks(completed, deadline) - for finding overdue tasks
