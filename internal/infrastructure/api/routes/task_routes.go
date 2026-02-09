package routes

import (
	"todo-api/internal/domain"
	"todo-api/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetTaskRoutes(router *gin.RouterGroup, repository domain.TaskRepository) {
	handler := handlers.NewTaskHandler(repository)

	tasks := router.Group("/todo")
	{
		// CRUD operations
		tasks.POST("", handler.CreateTask)
		tasks.GET("", handler.GetAllTasks)
		tasks.GET("/:id", handler.GetTask)
		tasks.PUT("/:id", handler.UpdateTask)
		tasks.DELETE("/:id", handler.DeleteTask)

		// Special queries
		tasks.GET("/responsible/:userID", handler.GetTasksByResponsible)
		tasks.GET("/author/:userID", handler.GetTasksByAuthor)
		tasks.GET("/overdue", handler.GetOverdueTasks)
	}
}
