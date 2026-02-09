package repositories

import (
	"database/sql"
	"fmt"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"
)

type TaskTypeRepository struct {
	db *sql.DB
}

func NewTaskTypeRepository(db *sql.DB) domain.TaskTypeRepository {
	return &TaskTypeRepository{db: db}
}

func (r *TaskTypeRepository) Create(taskType entities.TaskType) (entities.TaskType, error) {
	query := "INSERT INTO task_types (name) VALUES (?)"
	result, err := r.db.Exec(query, taskType.Name)
	if err != nil {
		return entities.TaskType{}, fmt.Errorf("failed to create task type: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.TaskType{}, fmt.Errorf("failed to get last insert id: %w", err)
	}

	taskType.ID = id
	return taskType, nil
}

func (r *TaskTypeRepository) GetByID(id int64) (entities.TaskType, error) {
	query := "SELECT id, name FROM task_types WHERE id = ?"
	var taskType entities.TaskType

	err := r.db.QueryRow(query, id).Scan(&taskType.ID, &taskType.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.TaskType{}, fmt.Errorf("task type not found")
		}
		return entities.TaskType{}, fmt.Errorf("failed to get task type: %w", err)
	}

	return taskType, nil
}

func (r *TaskTypeRepository) Update(taskType entities.TaskType) (entities.TaskType, error) {
	query := "UPDATE task_types SET name = ? WHERE id = ?"
	_, err := r.db.Exec(query, taskType.Name, taskType.ID)
	if err != nil {
		return entities.TaskType{}, fmt.Errorf("failed to update task type: %w", err)
	}

	return taskType, nil
}

func (r *TaskTypeRepository) Remove(id int64) error {
	query := "DELETE FROM task_types WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove task type: %w", err)
	}

	return nil
}

func (r *TaskTypeRepository) GetAll() ([]entities.TaskType, error) {
	query := "SELECT id, name FROM task_types"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all task types: %w", err)
	}
	defer rows.Close()

	var taskTypes []entities.TaskType
	for rows.Next() {
		var taskType entities.TaskType
		err := rows.Scan(&taskType.ID, &taskType.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task type: %w", err)
		}
		taskTypes = append(taskTypes, taskType)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating task types: %w", err)
	}

	return taskTypes, nil
}
