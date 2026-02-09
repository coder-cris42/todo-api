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

type mockWorkflowRepo struct {
	CreateFn func(w entities.Workflow) (entities.Workflow, error)
	GetAllFn func() ([]entities.Workflow, error)
}

func (m *mockWorkflowRepo) Create(w entities.Workflow) (entities.Workflow, error) {
	return m.CreateFn(w)
}
func (m *mockWorkflowRepo) GetByID(id int64) (entities.Workflow, error) {
	return entities.Workflow{}, nil
}
func (m *mockWorkflowRepo) Update(w entities.Workflow) (entities.Workflow, error) {
	return entities.Workflow{}, nil
}
func (m *mockWorkflowRepo) Remove(id int64) error                { return nil }
func (m *mockWorkflowRepo) GetAll() ([]entities.Workflow, error) { return m.GetAllFn() }

func TestCreateWorkflow_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockWorkflowRepo{}
	repo.CreateFn = func(w entities.Workflow) (entities.Workflow, error) {
		w.ID = 1
		return w, nil
	}

	handler := handlers.NewWorkflowHandler(repo)

	body := entities.Workflow{Name: "Default"}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodPost, "/workflows", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.CreateWorkflow(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d got %d", http.StatusCreated, w.Code)
	}
}

func TestGetAllWorkflows_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockWorkflowRepo{}
	repo.GetAllFn = func() ([]entities.Workflow, error) { return []entities.Workflow{{ID: 1, Name: "W"}}, nil }

	handler := handlers.NewWorkflowHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/workflows", nil)
	c.Request = req

	handler.GetAllWorkflows(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d got %d", http.StatusOK, w.Code)
	}
}
