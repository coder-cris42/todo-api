package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"
)

type WorkflowRepository struct {
	db *sql.DB
}

func NewWorkflowRepository(db *sql.DB) domain.WorkflowRepository {
	return &WorkflowRepository{db: db}
}

func (r *WorkflowRepository) Create(workflow entities.Workflow) (entities.Workflow, error) {
	statusesJSON, err := json.Marshal(workflow.Statuses)
	if err != nil {
		return entities.Workflow{}, fmt.Errorf("failed to marshal statuses: %w", err)
	}

	query := "INSERT INTO workflows (name, statuses, author_id, created_at) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, workflow.Name, statusesJSON, workflow.Author.ID, workflow.CreatedAt)
	if err != nil {
		return entities.Workflow{}, fmt.Errorf("failed to create workflow: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.Workflow{}, fmt.Errorf("failed to get last insert id: %w", err)
	}

	workflow.ID = id
	return workflow, nil
}

func (r *WorkflowRepository) GetByID(id int64) (entities.Workflow, error) {
	query := `SELECT w.id, w.name, w.statuses, w.author_id, w.created_at, u.id, u.name, u.username, u.email 
              FROM workflows w 
              JOIN users u ON w.author_id = u.id 
              WHERE w.id = ?`

	var workflow entities.Workflow
	var user entities.User
	var statusesJSON []byte

	err := r.db.QueryRow(query, id).Scan(
		&workflow.ID, &workflow.Name, &statusesJSON, &user.ID, &workflow.CreatedAt,
		&user.ID, &user.Name, &user.Username, &user.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Workflow{}, fmt.Errorf("workflow not found")
		}
		return entities.Workflow{}, fmt.Errorf("failed to get workflow: %w", err)
	}

	var statuses map[uint8]entities.TaskStatus
	if err := json.Unmarshal(statusesJSON, &statuses); err != nil {
		return entities.Workflow{}, fmt.Errorf("failed to unmarshal statuses: %w", err)
	}

	workflow.Author = user
	workflow.Statuses = statuses

	return workflow, nil
}

func (r *WorkflowRepository) Update(workflow entities.Workflow) (entities.Workflow, error) {
	statusesJSON, err := json.Marshal(workflow.Statuses)
	if err != nil {
		return entities.Workflow{}, fmt.Errorf("failed to marshal statuses: %w", err)
	}

	query := "UPDATE workflows SET name = ?, statuses = ?, author_id = ? WHERE id = ?"
	_, err = r.db.Exec(query, workflow.Name, statusesJSON, workflow.Author.ID, workflow.ID)
	if err != nil {
		return entities.Workflow{}, fmt.Errorf("failed to update workflow: %w", err)
	}

	return workflow, nil
}

func (r *WorkflowRepository) Remove(id int64) error {
	query := "DELETE FROM workflows WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove workflow: %w", err)
	}

	return nil
}

func (r *WorkflowRepository) GetAll() ([]entities.Workflow, error) {
	query := `SELECT w.id, w.name, w.statuses, w.author_id, w.created_at, u.id, u.name, u.username, u.email 
              FROM workflows w 
              JOIN users u ON w.author_id = u.id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all workflows: %w", err)
	}
	defer rows.Close()

	var workflows []entities.Workflow
	for rows.Next() {
		var workflow entities.Workflow
		var user entities.User
		var statusesJSON []byte

		err := rows.Scan(
			&workflow.ID, &workflow.Name, &statusesJSON, &user.ID, &workflow.CreatedAt,
			&user.ID, &user.Name, &user.Username, &user.Email,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan workflow: %w", err)
		}

		var statuses map[uint8]entities.TaskStatus
		if err := json.Unmarshal(statusesJSON, &statuses); err != nil {
			return nil, fmt.Errorf("failed to unmarshal statuses: %w", err)
		}

		workflow.Author = user
		workflow.Statuses = statuses
		workflows = append(workflows, workflow)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating workflows: %w", err)
	}

	return workflows, nil
}
