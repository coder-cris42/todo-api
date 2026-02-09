package integrationtests

import (
	"testing"
	"todo-api/internal/infrastructure/database/repositories"
)

// ExampleTestWithSQLiteInMemory demonstrates how to use SQLite in-memory for integration tests
func TestGetTaskById(t *testing.T) {
	// Initialize SQLite in-memory database with sample data
	db, err := InitializeTestDatabase(TestDatabaseConfig{
		Type: "sqlite",
		Path: "", // empty path means in-memory
	})
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer func() {
		if err := CleanupTestDatabase(db, "sqlite"); err != nil {
			t.Logf("Cleanup error: %v", err)
		}
	}()

	taskRepository := repositories.NewTaskRepository(db)

	firstTask, err := taskRepository.GetByID(1)
	if err != nil {
		t.Fatalf("Failed to get first task: %v", err)
	}

	if firstTask.ID != 1 {
		t.Errorf("Expected task ID 1, got %d", firstTask.ID)
	}
	if firstTask.Title != "Implement User Authentication" {
		t.Errorf("Expected task title to be 'Implement User Authentication', got '%s'", firstTask.Title)
	}

	t.Logf("Successfully queried first task with ID %d from SQLite in-memory database", firstTask.ID)
}

func TestGetAllTasks(t *testing.T) {
	// Initialize SQLite in-memory database with sample data
	db, err := InitializeTestDatabase(TestDatabaseConfig{
		Type: "sqlite",
		Path: "", // empty path means in-memory
	})
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer func() {
		if err := CleanupTestDatabase(db, "sqlite"); err != nil {
			t.Logf("Cleanup error: %v", err)
		}
	}()

	taskRepository := repositories.NewTaskRepository(db)

	allTasks, err := taskRepository.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all tasks: %v", err)
	}

	if len(allTasks) == 0 {
		t.Error("Expected at least one task, got none")
	}

	if len(allTasks) != 12 {
		t.Errorf("Expected 12 tasks, got %d", len(allTasks))
	}

	t.Logf("Successfully queried all tasks from SQLite in-memory database, got %d tasks", len(allTasks))
}
