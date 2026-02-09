package routes

import (
	"todo-api/internal/domain"
	"todo-api/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetWorkflowRoutes(router *gin.RouterGroup, repository domain.WorkflowRepository) {
	handler := handlers.NewWorkflowHandler(repository)

	workflows := router.Group("/workflows")
	{
		workflows.POST("", handler.CreateWorkflow)
		workflows.GET("", handler.GetAllWorkflows)
		workflows.GET("/:id", handler.GetWorkflow)
		workflows.PUT("/:id", handler.UpdateWorkflow)
		workflows.DELETE("/:id", handler.DeleteWorkflow)
	}
}
