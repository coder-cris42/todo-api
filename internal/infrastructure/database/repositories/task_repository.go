package repositories

import (
	"database/sql"
	"fmt"
	"time"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	return &TaskRepository{db: db}
}

func (self *TaskRepository) Create(task entities.Task) (entities.Task, error) {
	query := `INSERT INTO tasks (title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	var parentID *int64
	if task.Parent != nil {
		parentID = &task.Parent.ID
	}

	result, err := self.db.Exec(query,
		task.Title, task.Description, task.Status.ID, parentID, task.AuthorID,
		task.Deadline, task.CreatedAt, task.UpdatedAt, task.ResponsibleID,
		task.Workflow.ID, task.Type.ID, task.Completed,
	)
	if err != nil {
		return entities.Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.Task{}, fmt.Errorf("failed to get last insert id: %w", err)
	}

	task.ID = id
	return task, nil
}

func (self *TaskRepository) GetByID(id int64) (entities.Task, error) {
	query := `SELECT id, title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed 
              FROM tasks WHERE id = ?`

	var task entities.Task
	var parentID sql.NullInt64
	var statusID, workflowID, typeID int64

	err := self.db.QueryRow(query, id).Scan(
		&task.ID, &task.Title, &task.Description, &statusID, &parentID, &task.AuthorID,
		&task.Deadline, &task.CreatedAt, &task.UpdatedAt, &task.ResponsibleID,
		&workflowID, &typeID, &task.Completed,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Task{}, fmt.Errorf("task not found")
		}
		return entities.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	if parentID.Valid {
		task.Parent = &entities.Task{ID: parentID.Int64}
	}

	task.Status = entities.TaskStatus{ID: statusID}
	task.Workflow = entities.Workflow{ID: workflowID}
	task.Type = entities.TaskType{ID: typeID}

	return task, nil
}

func (self *TaskRepository) Update(task entities.Task) (entities.Task, error) {
	query := `UPDATE tasks SET title = ?, description = ?, status_id = ?, parent_id = ?, 
              deadline = ?, updated_at = ?, responsible_id = ?, workflow_id = ?, type_id = ?, completed = ? 
              WHERE id = ?`

	var parentID *int64
	if task.Parent != nil {
		parentID = &task.Parent.ID
	}

	_, err := self.db.Exec(query,
		task.Title, task.Description, task.Status.ID, parentID,
		task.Deadline, time.Now(), task.ResponsibleID, task.Workflow.ID, task.Type.ID,
		task.Completed, task.ID,
	)
	if err != nil {
		return entities.Task{}, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}

func (self *TaskRepository) GetAll() ([]entities.Task, error) {
	query := `SELECT id, title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed 
              FROM tasks`

	return self.scanTasks(self.db.Query(query))
}

func (self *TaskRepository) GetAllByResponsible(userID int64) ([]entities.Task, error) {
	query := `SELECT id, title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed 
              FROM tasks WHERE responsible_id = ?`

	return self.scanTasks(self.db.Query(query, userID))
}

func (self *TaskRepository) GetAllByAuthor(userID int64) ([]entities.Task, error) {
	query := `SELECT id, title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed 
              FROM tasks WHERE author_id = ?`

	return self.scanTasks(self.db.Query(query, userID))
}

func (self *TaskRepository) GetAllByStatus(status entities.TaskStatus) ([]entities.Task, error) {
	query := `SELECT id, title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed 
              FROM tasks WHERE status_id = ?`

	return self.scanTasks(self.db.Query(query, status.ID))
}

func (self *TaskRepository) GetAllOverdue() ([]entities.Task, error) {
	query := `SELECT id, title, description, status_id, parent_id, author_id, deadline, 
              created_at, updated_at, responsible_id, workflow_id, type_id, completed 
              FROM tasks WHERE completed = false AND deadline < NOW()`

	return self.scanTasks(self.db.Query(query))
}

func (self *TaskRepository) Remove(id int64) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := self.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove task: %w", err)
	}

	return nil
}

func (self *TaskRepository) scanTasks(rows *sql.Rows, err error) ([]entities.Task, error) {
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []entities.Task
	for rows.Next() {
		var task entities.Task
		var parentID sql.NullInt64
		var statusID, workflowID, typeID int64

		err := rows.Scan(
			&task.ID, &task.Title, &task.Description, &statusID, &parentID, &task.AuthorID,
			&task.Deadline, &task.CreatedAt, &task.UpdatedAt, &task.ResponsibleID,
			&workflowID, &typeID, &task.Completed,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		if parentID.Valid {
			task.Parent = &entities.Task{ID: parentID.Int64}
		}

		task.Status = entities.TaskStatus{ID: statusID}
		task.Workflow = entities.Workflow{ID: workflowID}
		task.Type = entities.TaskType{ID: typeID}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tasks: %w", err)
	}

	return tasks, nil
}
