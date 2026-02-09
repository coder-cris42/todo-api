package handlers

import (
	"net/http"
	"strconv"
	"todo-api/internal/domain"
	"todo-api/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type WorkflowHandler struct {
	repository domain.WorkflowRepository
}

func NewWorkflowHandler(repository domain.WorkflowRepository) *WorkflowHandler {
	return &WorkflowHandler{
		repository: repository,
	}
}

// CreateWorkflow creates a new workflow
// @POST /workflows
func (h *WorkflowHandler) CreateWorkflow(c *gin.Context) {
	var workflow entities.Workflow

	if err := c.ShouldBindJSON(&workflow); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdWorkflow, err := h.repository.Create(workflow)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusCreated, createdWorkflow)
}

// GetWorkflow retrieves a workflow by ID
// @GET /workflows/:id
func (h *WorkflowHandler) GetWorkflow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workflow ID"})
		return
	}

	workflow, err := h.repository.GetByID(id)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, workflow)
}

// GetAllWorkflows retrieves all workflows
// @GET /workflows
func (h *WorkflowHandler) GetAllWorkflows(c *gin.Context) {
	workflows, err := h.repository.GetAll()
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, workflows)
}

// UpdateWorkflow updates a workflow
// @PUT /workflows/:id
func (h *WorkflowHandler) UpdateWorkflow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workflow ID"})
		return
	}

	var workflow entities.Workflow
	if err := c.ShouldBindJSON(&workflow); err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workflow.ID = id
	updatedWorkflow, err := h.repository.Update(workflow)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addSuccessHeaders(c)
	addValidationHeaders(c)
	c.JSON(http.StatusOK, updatedWorkflow)
}

// DeleteWorkflow deletes a workflow
// @DELETE /workflows/:id
func (h *WorkflowHandler) DeleteWorkflow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		addErrorHeaders(c)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workflow ID"})
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
