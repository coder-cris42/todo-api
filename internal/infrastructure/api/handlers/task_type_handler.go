package handlers

import (
	"net/http"
	"strconv"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type TaskTypeHandler struct {
	repository domain.TaskTypeRepository
}

func NewTaskTypeHandler(repository domain.TaskTypeRepository) *TaskTypeHandler {
	return &TaskTypeHandler{
		repository: repository,
	}
}

// CreateTaskType creates a new task type
// @POST /task-types
func (h *TaskTypeHandler) CreateTaskType(c *gin.Context) {
	var taskType entities.TaskType

	if err := c.ShouldBindJSON(&taskType); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTaskType, err := h.repository.Create(taskType)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusCreated, createdTaskType)
}

// GetTaskType retrieves a task type by ID
// @GET /task-types/:id
func (h *TaskTypeHandler) GetTaskType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type ID"})
		return
	}

	taskType, err := h.repository.GetByID(id)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, taskType)
}

// GetAllTaskTypes retrieves all task types
// @GET /task-types
func (h *TaskTypeHandler) GetAllTaskTypes(c *gin.Context) {
	taskTypes, err := h.repository.GetAll()
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, taskTypes)
}

// UpdateTaskType updates a task type
// @PUT /task-types/:id
func (h *TaskTypeHandler) UpdateTaskType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type ID"})
		return
	}

	var taskType entities.TaskType
	if err := c.ShouldBindJSON(&taskType); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskType.ID = id
	updatedTaskType, err := h.repository.Update(taskType)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, updatedTaskType)
}

// DeleteTaskType deletes a task type
// @DELETE /task-types/:id
func (h *TaskTypeHandler) DeleteTaskType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type ID"})
		return
	}

	err = h.repository.Remove(id)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusNoContent, nil)
}
