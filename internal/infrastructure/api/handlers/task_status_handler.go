package handlers

import (
	"net/http"
	"strconv"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type TaskStatusHandler struct {
	repository domain.TaskStatusRepository
}

func NewTaskStatusHandler(repository domain.TaskStatusRepository) *TaskStatusHandler {
	return &TaskStatusHandler{
		repository: repository,
	}
}

// CreateTaskStatus creates a new task status
// @POST /task-statuses
func (h *TaskStatusHandler) CreateTaskStatus(c *gin.Context) {
	var status entities.TaskStatus

	if err := c.ShouldBindJSON(&status); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdStatus, err := h.repository.Create(status)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusCreated, createdStatus)
}

// GetTaskStatus retrieves a task status by ID
// @GET /task-statuses/:id
func (h *TaskStatusHandler) GetTaskStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	status, err := h.repository.GetByID(id)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, status)
}

// GetAllTaskStatuses retrieves all task statuses
// @GET /task-statuses
func (h *TaskStatusHandler) GetAllTaskStatuses(c *gin.Context) {
	statuses, err := h.repository.GetAll()
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, statuses)
}

// UpdateTaskStatus updates a task status
// @PUT /task-statuses/:id
func (h *TaskStatusHandler) UpdateTaskStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	var status entities.TaskStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status.ID = id
	updatedStatus, err := h.repository.Update(status)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, updatedStatus)
}

// DeleteTaskStatus deletes a task status
// @DELETE /task-statuses/:id
func (h *TaskStatusHandler) DeleteTaskStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
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
