# TODO API - OpenAPI/Swagger Documentation

This document describes the Swagger/OpenAPI endpoint integration for the TODO API.

## Overview

The API now includes comprehensive OpenAPI (Swagger) documentation that is automatically generated and served at runtime. This provides an interactive API explorer where you can test all endpoints directly in your browser.

## Accessing Swagger UI

Once the application is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

The API endpoints are available at:

```
Base URL: http://localhost:8080/api/v1
```

## API Documentation Structure

### Tags

The API is organized into the following resource categories:

1. **Health** - Application health check endpoints
2. **Tasks** - Task management operations
3. **Task Status** - Task status management operations
4. **Task Type** - Task type management operations
5. **Workflows** - Workflow management operations

## Endpoints

### Health Check

- **GET** `/health` - Check the health status of the API and database connection

### Tasks (Base Path: `/todo`)

- **POST** `/todo` - Create a new task
- **GET** `/todo` - Get all tasks
- **GET** `/todo/{id}` - Get a specific task
- **PUT** `/todo/{id}` - Update a task
- **DELETE** `/todo/{id}` - Delete a task
- **GET** `/todo/responsible/{userID}` - Get tasks by responsible user
- **GET** `/todo/author/{userID}` - Get tasks by author
- **GET** `/todo/overdue` - Get overdue tasks

### Task Status (Base Path: `/statuses`)

- **POST** `/statuses` - Create a new task status
- **GET** `/statuses` - Get all task statuses
- **GET** `/statuses/{id}` - Get a specific task status
- **PUT** `/statuses/{id}` - Update a task status
- **DELETE** `/statuses/{id}` - Delete a task status

### Task Type (Base Path: `/task-type`)

- **POST** `/task-type` - Create a new task type
- **GET** `/task-type` - Get all task types
- **GET** `/task-type/{id}` - Get a specific task type
- **PUT** `/task-type/{id}` - Update a task type
- **DELETE** `/task-type/{id}` - Delete a task type

### Workflows (Base Path: `/workflows`)

- **POST** `/workflows` - Create a new workflow
- **GET** `/workflows` - Get all workflows
- **GET** `/workflows/{id}` - Get a specific workflow
- **PUT** `/workflows/{id}` - Update a workflow
- **DELETE** `/workflows/{id}` - Delete a workflow

## Running the Application

### Prerequisites

- Go 1.24.9 or later
- MySQL database running
- Environment variables configured

### Environment Variables

```bash
DB_USERNAME=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=todo_database
SERVER_PORT=8080
```

### Start the Application

```bash
# Install dependencies
make install-deps

# Build and run
make run

# Or run directly
go run ./cmd/main.go
```

### Building for Production

```bash
make build
./bin/todo-api
```

## OpenAPI Documentation Files

The Swagger documentation is composed of:

1. **[docs/swagger.yaml](docs/swagger.yaml)** - OpenAPI 3.0 specification in YAML format
2. **[docs/docs.go](docs/docs.go)** - Generated Go file containing embedded JSON specification
3. **[internal/infrastructure/api/swagger.go](internal/infrastructure/api/swagger.go)** - Gin middleware to serve Swagger UI

## Testing API Endpoints with Swagger UI

1. Open your browser and navigate to `http://localhost:8080/swagger/index.html`
2. Click on any endpoint to expand it
3. Click the "Try it out" button
4. Fill in the required parameters and request body
5. Click "Execute" to send the request
6. View the response below

## Example Requests

### Create a Task

```bash
curl -X POST http://localhost:8080/api/v1/todo \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Fix login bug",
    "description": "There is an issue with the login form validation",
    "authorID": 1,
    "deadline": "2026-02-28T15:30:00Z"
  }'
```

### Get All Tasks

```bash
curl http://localhost:8080/api/v1/todo
```

### Get Task by ID

```bash
curl http://localhost:8080/api/v1/todo/1
```

### Update a Task

```bash
curl -X PUT http://localhost:8080/api/v1/todo/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated title",
    "completed": true
  }'
```

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/api/v1/todo/1
```

### Get Overdue Tasks

```bash
curl http://localhost:8080/api/v1/todo/overdue
```

### Create a Task Status

```bash
curl -X POST http://localhost:8080/api/v1/statuses \
  -H "Content-Type: application/json" \
  -d '{
    "label": "In Progress",
    "active": true
  }'
```

### Create a Workflow

```bash
curl -X POST http://localhost:8080/api/v1/workflows \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Standard Workflow",
    "statuses": {
      "0": {"label": "To Do", "active": true},
      "1": {"label": "In Progress", "active": true},
      "2": {"label": "Done", "active": true}
    }
  }'
```

## Swagger UI Features

The Swagger UI provides:

- **Interactive API Documentation** - See all endpoints and their parameters
- **Try It Out** - Test endpoints directly from the browser
- **Request/Response Examples** - View sample requests and responses
- **Parameter Validation** - Automatic validation of parameters
- **Response Status Codes** - Understand different response codes
- **Download Specification** - Export API specification as JSON or YAML

## API Schemas

The documentation includes comprehensive schema definitions for:

- **Task** - Task with all properties
- **CreateTaskRequest** - Request body for creating tasks
- **UpdateTaskRequest** - Request body for updating tasks
- **TaskStatus** - Task status entity
- **TaskType** - Task type entity
- **Workflow** - Workflow entity
- **User** - User entity
- **Error** - Error response format

## Response Formats

### Success Response Example

```json
{
  "id": 1,
  "title": "Fix login bug",
  "description": "There is an issue with the login form validation",
  "status": {
    "id": 1,
    "label": "To Do",
    "active": true
  },
  "authorID": 1,
  "responsibleID": 2,
  "deadline": "2026-02-28T15:30:00Z",
  "createdAt": "2026-02-08T10:30:00Z",
  "updatedAt": "2026-02-08T10:30:00Z",
  "completed": false,
  "type": {
    "id": 1,
    "name": "Bug"
  }
}
```

### Error Response Example

```json
{
  "error": "ERROR_DATABASE_CONNECTION",
  "message": "Failed to connect to database"
}
```

## Integration Details

### Swagger Setup File

The [internal/infrastructure/api/swagger.go](internal/infrastructure/api/swagger.go) file registers the Swagger endpoint:

```go
func SetupSwagger(router *gin.Engine) {
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

### Main Application Integration

In [cmd/main.go](cmd/main.go), the Swagger setup is initialized:

```go
import _ "todo-api/docs"

// In main function:
error_codes.SetupSwagger(router)
```

## Updating API Documentation

If you modify your API endpoints:

1. Update the OpenAPI specification in [docs/swagger.yaml](docs/swagger.yaml)
2. Update the corresponding [docs/docs.go](docs/docs.go) file
3. Rebuild and restart the application

## Useful Links

- [OpenAPI 3.0 Specification](https://spec.openapis.org/oas/v3.0.3)
- [Swagger UI Documentation](https://swagger.io/tools/swagger-ui/)
- [Gin Swagger Documentation](https://github.com/swaggo/gin-swagger)
- [Swagger Editor](https://editor.swagger.io/) - Online editor for testing YAML

## Troubleshooting

### Swagger UI not loading

1. Check that the application is running on the correct port
2. Ensure the URL is correct: `http://localhost:8080/swagger/index.html`
3. Check browser console for errors
4. Verify that `docs` package is imported in `main.go`

### Endpoints not showing in Swagger UI

1. Ensure all routes are registered before calling `SetupSwagger()`
2. Check that the OpenAPI specification includes all endpoints
3. Verify the server URL matches your actual API URL

## Support

For questions or issues with the API or Swagger documentation, please refer to:
- [Gin Framework Documentation](https://gin-gonic.com/)
- [Swagger/OpenAPI Documentation](https://swagger.io/specification/)
- Application repository and issue tracker
