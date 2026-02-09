package main

import (
	"fmt"
	"log"
	error_codes "todo-api/internal/infrastructure/api"
	"todo-api/internal/infrastructure/api/routes"
	"todo-api/internal/infrastructure/database/connection"
	"todo-api/internal/infrastructure/database/repositories"
	"todo-api/internal/middleware"
	"todo-api/utils"

	_ "todo-api/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	printBanner()

	err := utils.CheckEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

	DB_USERNAME := utils.GetEnvironmentVariable("DB_USERNAME")
	DB_PASSWORD := utils.GetEnvironmentVariable("DB_PASSWORD")
	DB_HOST := utils.GetEnvironmentVariable("DB_HOST")
	DB_PORT := utils.GetEnvironmentVariable("DB_PORT")
	DB_NAME := utils.GetEnvironmentVariable("DB_NAME")
	SERVER_PORT := utils.GetEnvironmentVariable("SERVER_PORT")

	err = utils.CheckDatabaseConnection(DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	db, err := connection.NewMySQLDB(DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.RedirectTrailingSlash = false

	// Apply global security headers middleware
	router.Use(middleware.SecurityHeaders())
	router.Use(middleware.CORSHeaders())
	router.Use(middleware.RequestID())
	router.Use(middleware.InputValidation())
	router.Use(middleware.ResponseValidation())

	router.GET("/health", func(c *gin.Context) {
		// Security Practice: Return 200 to avoid exposing internal details to potential attackers when the application breaks for some reason
		// Internal error code is returned to not expose internal details about the failure

		err := utils.CheckEnvironmentVariables()
		if err != nil {
			log.Println("Environment variable error:", err)
			c.Header("X-Error-Response", "true")
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.JSON(200, gin.H{
				"error": error_codes.ERROR_ENVIRONMENT_VARIABLES,
			})
			return
		}

		err = utils.CheckDatabaseConnection(DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
		if err != nil {
			log.Println("Database connection error:", err)
			c.Header("X-Error-Response", "true")
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.JSON(200, gin.H{
				"error": error_codes.ERROR_DATABASE_CONNECTION,
			})
			return
		}

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	apiV1 := router.Group("/api/v1")

	routes.SetTaskRoutes(apiV1, repositories.NewTaskRepository(db))
	routes.SetTaskTypeRoutes(apiV1, repositories.NewTaskTypeRepository(db))
	routes.SetTaskStatusRoutes(apiV1, repositories.NewTaskStatusRepository(db))
	routes.SetWorkflowRoutes(apiV1, repositories.NewWorkflowRepository(db))

	// Swagger
	router.Group("")
	error_codes.SetupSwagger(router)

	log.Printf("Starting TO-DO API application on port %s", SERVER_PORT)

	if err := router.Run(":" + SERVER_PORT); err != nil {
		log.Fatal(err)
	}

}

func printBanner() {
	fmt.Println(`
░██████████  ░██████   ░███████     ░██████              ░███    ░█████████  ░██████
    ░██     ░██   ░██  ░██   ░██   ░██   ░██            ░██░██   ░██     ░██   ░██  
    ░██    ░██     ░██ ░██    ░██ ░██     ░██          ░██  ░██  ░██     ░██   ░██  
    ░██    ░██     ░██ ░██    ░██ ░██     ░██ ░██████ ░█████████ ░█████████    ░██  
    ░██    ░██     ░██ ░██    ░██ ░██     ░██         ░██    ░██ ░██           ░██  
    ░██     ░██   ░██  ░██   ░██   ░██   ░██          ░██    ░██ ░██           ░██  
    ░██      ░██████   ░███████     ░██████           ░██    ░██ ░██         ░██████`)
}
