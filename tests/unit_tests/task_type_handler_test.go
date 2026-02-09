package unittests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-api/internal/domain/entities"
	"todo-api/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

type mockTypeRepo struct {
	CreateFn func(taskType entities.TaskType) (entities.TaskType, error)
	GetAllFn func() ([]entities.TaskType, error)
}

func (m *mockTypeRepo) Create(taskType entities.TaskType) (entities.TaskType, error) {
	return m.CreateFn(taskType)
}
func (m *mockTypeRepo) GetByID(id int64) (entities.TaskType, error) { return entities.TaskType{}, nil }
func (m *mockTypeRepo) Update(taskType entities.TaskType) (entities.TaskType, error) {
	return entities.TaskType{}, nil
}
func (m *mockTypeRepo) Remove(id int64) error                { return nil }
func (m *mockTypeRepo) GetAll() ([]entities.TaskType, error) { return m.GetAllFn() }

func TestCreateTaskType_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTypeRepo{}
	repo.CreateFn = func(taskType entities.TaskType) (entities.TaskType, error) {
		taskType.ID = 2
		return taskType, nil
	}

	handler := handlers.NewTaskTypeHandler(repo)

	body := entities.TaskType{Name: "Bug"}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodPost, "/task-types", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.CreateTaskType(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d got %d", http.StatusCreated, w.Code)
	}
}

func TestGetAllTaskTypes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTypeRepo{}
	repo.GetAllFn = func() ([]entities.TaskType, error) { return []entities.TaskType{{ID: 1, Name: "Bug"}}, nil }

	handler := handlers.NewTaskTypeHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/task-types", nil)
	c.Request = req

	handler.GetAllTaskTypes(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d got %d", http.StatusOK, w.Code)
	}
}
