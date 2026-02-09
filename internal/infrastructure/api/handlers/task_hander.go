package handlers

import (
	"net/http"
	"strconv"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	repository domain.TaskRepository
}

func NewTaskHandler(repository domain.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repository: repository,
	}
}

// CreateTask creates a new task
// @POST /todo
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := h.repository.Create(task)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusCreated, createdTask)
}

// GetTask retrieves a task by ID
// @GET /todo/:id
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.repository.GetByID(id)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, task)
}

// GetAllTasks retrieves all tasks
// @GET /todo
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.repository.GetAll()
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, tasks)
}

// UpdateTask updates a task
// @PUT /todo/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = id
	updatedTask, err := h.repository.Update(task)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask deletes a task
// @DELETE /todo/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
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

// GetTasksByResponsible retrieves all tasks assigned to a user
// @GET /todo/responsible/:userID
func (h *TaskHandler) GetTasksByResponsible(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := h.repository.GetAllByResponsible(userID)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, tasks)
}

// GetTasksByAuthor retrieves all tasks created by a user
// @GET /todo/author/:userID
func (h *TaskHandler) GetTasksByAuthor(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := h.repository.GetAllByAuthor(userID)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, tasks)
}

// GetOverdueTasks retrieves all overdue tasks
// @GET /todo/overdue
func (h *TaskHandler) GetOverdueTasks(c *gin.Context) {
	tasks, err := h.repository.GetAllOverdue()
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, tasks)
}
