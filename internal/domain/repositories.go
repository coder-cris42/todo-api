package domain

import "todo-api/internal/domain/entities"

type TaskStatusRepository interface {
	Create(status entities.TaskStatus) (entities.TaskStatus, error)
	GetByID(id int64) (entities.TaskStatus, error)
	Update(status entities.TaskStatus) (entities.TaskStatus, error)
	Remove(id int64) error
	GetAll() ([]entities.TaskStatus, error)
}

type TaskTypeRepository interface {
	Create(taskType entities.TaskType) (entities.TaskType, error)
	GetByID(id int64) (entities.TaskType, error)
	Update(taskType entities.TaskType) (entities.TaskType, error)
	Remove(id int64) error
	GetAll() ([]entities.TaskType, error)
}

type WorkflowRepository interface {
	Create(workflow entities.Workflow) (entities.Workflow, error)
	GetByID(id int64) (entities.Workflow, error)
	Update(workflow entities.Workflow) (entities.Workflow, error)
	Remove(id int64) error
	GetAll() ([]entities.Workflow, error)
}

type TaskRepository interface {
	Create(task entities.Task) (entities.Task, error)
	GetByID(id int64) (entities.Task, error)
	Update(task entities.Task) (entities.Task, error)
	GetAll() ([]entities.Task, error)
	GetAllByResponsible(userID int64) ([]entities.Task, error)
	GetAllByAuthor(userID int64) ([]entities.Task, error)
	GetAllByStatus(status entities.TaskStatus) ([]entities.Task, error)
	GetAllOverdue() ([]entities.Task, error)
	Remove(id int64) error
}

type StatusRepository interface {
	Create(status entities.TaskStatus) (entities.TaskStatus, error)
	GetByID(id int64) (entities.TaskStatus, error)
	Update(status entities.TaskStatus) (entities.TaskStatus, error)
	Remove(id int64) error
	GetAll() ([]entities.TaskStatus, error)
}
