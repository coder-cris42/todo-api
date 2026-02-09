package unittests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"todo-api/internal/domain/entities"
	"todo-api/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

type mockTaskRepo struct {
	CreateFn              func(task entities.Task) (entities.Task, error)
	GetByIDFn             func(id int64) (entities.Task, error)
	UpdateFn              func(task entities.Task) (entities.Task, error)
	GetAllFn              func() ([]entities.Task, error)
	GetAllByResponsibleFn func(userID int64) ([]entities.Task, error)
	RemoveFn              func(id int64) error
}

func (m *mockTaskRepo) Create(task entities.Task) (entities.Task, error) { return m.CreateFn(task) }
func (m *mockTaskRepo) GetByID(id int64) (entities.Task, error)          { return m.GetByIDFn(id) }
func (m *mockTaskRepo) Update(task entities.Task) (entities.Task, error) { return m.UpdateFn(task) }
func (m *mockTaskRepo) GetAll() ([]entities.Task, error)                 { return m.GetAllFn() }
func (m *mockTaskRepo) GetAllByResponsible(userID int64) ([]entities.Task, error) {
	return m.GetAllByResponsibleFn(userID)
}
func (m *mockTaskRepo) GetAllByAuthor(userID int64) ([]entities.Task, error) { return nil, nil }
func (m *mockTaskRepo) GetAllByStatus(status entities.TaskStatus) ([]entities.Task, error) {
	return nil, nil
}
func (m *mockTaskRepo) GetAllOverdue() ([]entities.Task, error) { return nil, nil }
func (m *mockTaskRepo) Remove(id int64) error                   { return m.RemoveFn(id) }

func TestCreateTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTaskRepo{}
	repo.CreateFn = func(task entities.Task) (entities.Task, error) {
		task.ID = 123
		return task, nil
	}

	handler := handlers.NewTaskHandler(repo)

	body := entities.Task{Title: "T1", Description: "D1", AuthorID: 1, Deadline: time.Now()}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.CreateTask(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d got %d, body: %s", http.StatusCreated, w.Code, w.Body.String())
	}
}

func TestGetTask_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTaskRepo{}
	repo.GetByIDFn = func(id int64) (entities.Task, error) { return entities.Task{}, errors.New("not found") }

	handler := handlers.NewTaskHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/todo/999", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "999"}}

	handler.GetTask(c)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status %d got %d", http.StatusNotFound, w.Code)
	}
}

func TestGetAllTasks_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTaskRepo{}
	repo.GetAllFn = func() ([]entities.Task, error) {
		return []entities.Task{{ID: 1, Title: "A"}}, nil
	}

	handler := handlers.NewTaskHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/todo", nil)
	c.Request = req

	handler.GetAllTasks(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d got %d", http.StatusOK, w.Code)
	}
}

func TestUpdateTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTaskRepo{}
	repo.UpdateFn = func(task entities.Task) (entities.Task, error) { return task, nil }

	handler := handlers.NewTaskHandler(repo)

	body := entities.Task{Title: "Updated"}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodPut, "/todo/1", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	handler.UpdateTask(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d got %d", http.StatusOK, w.Code)
	}
}

func TestDeleteTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTaskRepo{}
	repo.RemoveFn = func(id int64) error { return nil }

	handler := handlers.NewTaskHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodDelete, "/todo/1", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	handler.DeleteTask(c)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status %d got %d", http.StatusNoContent, w.Code)
	}
}

func TestGetTasksByResponsible_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &mockTaskRepo{}
	repo.GetAllByResponsibleFn = func(userID int64) ([]entities.Task, error) {
		return []entities.Task{{ID: 5, Title: "R"}}, nil
	}

	handler := handlers.NewTaskHandler(repo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/todo/responsible/2", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "userID", Value: "2"}}

	handler.GetTasksByResponsible(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d got %d", http.StatusOK, w.Code)
	}
}
