package repositories

import (
	"database/sql"
	"fmt"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"
)

type TaskStatusRepository struct {
	db *sql.DB
}

func NewTaskStatusRepository(db *sql.DB) domain.TaskStatusRepository {
	return &TaskStatusRepository{db: db}
}

func (r *TaskStatusRepository) Create(status entities.TaskStatus) (entities.TaskStatus, error) {
	query := "INSERT INTO task_statuses (label, active) VALUES (?, ?)"
	result, err := r.db.Exec(query, status.Label, status.Active)
	if err != nil {
		return entities.TaskStatus{}, fmt.Errorf("failed to create task status: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.TaskStatus{}, fmt.Errorf("failed to get last insert id: %w", err)
	}

	status.ID = id
	return status, nil
}

func (r *TaskStatusRepository) GetByID(id int64) (entities.TaskStatus, error) {
	query := "SELECT id, label, active FROM task_statuses WHERE id = ?"
	var status entities.TaskStatus

	err := r.db.QueryRow(query, id).Scan(&status.ID, &status.Label, &status.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.TaskStatus{}, fmt.Errorf("task status not found")
		}
		return entities.TaskStatus{}, fmt.Errorf("failed to get task status: %w", err)
	}

	return status, nil
}

func (r *TaskStatusRepository) Update(status entities.TaskStatus) (entities.TaskStatus, error) {
	query := "UPDATE task_statuses SET label = ?, active = ? WHERE id = ?"
	_, err := r.db.Exec(query, status.Label, status.Active, status.ID)
	if err != nil {
		return entities.TaskStatus{}, fmt.Errorf("failed to update task status: %w", err)
	}

	return status, nil
}

func (r *TaskStatusRepository) Remove(id int64) error {
	query := "DELETE FROM task_statuses WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove task status: %w", err)
	}

	return nil
}

func (r *TaskStatusRepository) GetAll() ([]entities.TaskStatus, error) {
	query := "SELECT id, label, active FROM task_statuses"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all task statuses: %w", err)
	}
	defer rows.Close()

	var statuses []entities.TaskStatus
	for rows.Next() {
		var status entities.TaskStatus
		err := rows.Scan(&status.ID, &status.Label, &status.Active)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task status: %w", err)
		}
		statuses = append(statuses, status)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating task statuses: %w", err)
	}

	return statuses, nil
}
