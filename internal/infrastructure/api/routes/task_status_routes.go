package routes

import (
	"todo-api/internal/domain"
	"todo-api/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetTaskStatusRoutes(router *gin.RouterGroup, repository domain.TaskStatusRepository) {
	handler := handlers.NewTaskStatusHandler(repository)

	statuses := router.Group("/statuses")
	{
		statuses.POST("", handler.CreateTaskStatus)
		statuses.GET("", handler.GetAllTaskStatuses)
		statuses.GET("/:id", handler.GetTaskStatus)
		statuses.PUT("/:id", handler.UpdateTaskStatus)
		statuses.DELETE("/:id", handler.DeleteTaskStatus)
	}
}
