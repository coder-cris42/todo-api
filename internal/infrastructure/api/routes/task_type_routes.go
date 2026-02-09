package routes

import (
	"todo-api/internal/domain"
	"todo-api/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetTaskTypeRoutes(router *gin.RouterGroup, repository domain.TaskTypeRepository) {
	handler := handlers.NewTaskTypeHandler(repository)

	types := router.Group("/task-type")
	{
		types.POST("", handler.CreateTaskType)
		types.GET("", handler.GetAllTaskTypes)
		types.GET("/:id", handler.GetTaskType)
		types.PUT("/:id", handler.UpdateTaskType)
		types.DELETE("/:id", handler.DeleteTaskType)
	}
}
