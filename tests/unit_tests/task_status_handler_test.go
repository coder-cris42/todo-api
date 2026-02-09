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

type mockStatusRepo struct {
	CreateFn func(status entities.TaskStatus) (entities.TaskStatus, error)
	GetAllFn func() ([]entities.TaskStatus, error)
}

func (m *mockStatusRepo) Create(status entities.TaskStatus) (entities.TaskStatus, error) {
	return m.CreateFn(status)
}
func (m *mockStatusRepo) GetByID(id int64) (entities.TaskStatus, error) {
	return entities.TaskStatus{}, nil
}
func (m *mockStatusRepo) Update(status entities.TaskStatus) (entities.TaskStatus, error) {
	return entities.TaskStatus{}, nil
}
func (m *mockStatusRepo) Remove(id int64) error                  { return nil }
func (m *mockStatusRepo) GetAll() ([]entities.TaskStatus, error) { return m.GetAllFn() }

func TestCreateTaskStatus_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockStatusRepo{}
	repo.CreateFn = func(status entities.TaskStatus) (entities.TaskStatus, error) {
		status.ID = 1
		return status, nil
	}

	handler := handlers.NewTaskStatusHandler(repo)

	body := entities.TaskStatus{Label: "Todo", Active: true}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodPost, "/task-statuses", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.CreateTaskStatus(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d got %d", http.StatusCreated, w.Code)
	}
}

func TestGetAllTaskStatuses_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockStatusRepo{}
	repo.GetAllFn = func() ([]entities.TaskStatus, error) { return []entities.TaskStatus{{ID: 1, Label: "T"}}, nil }

	handler := handlers.NewTaskStatusHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/task-statuses", nil)
	c.Request = req

	handler.GetAllTaskStatuses(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d got %d", http.StatusOK, w.Code)
	}
}
