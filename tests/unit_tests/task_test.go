package unittests

import (
	"testing"
	"time"
	todo "todo-api/internal/domain/entities"
)

func TestTask_IsOverdue(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	task := todo.NewTask("t", "d", 1, past, todo.TaskType{ID: 1})
	task.Completed = false

	if !task.IsOverdue() {
		t.Fatal("expected task to be overdue")
	}
}

func TestTask_AssignToAndUpdatedAt(t *testing.T) {
	task := todo.NewTask("t", "d", 1, time.Now().Add(24*time.Hour), todo.TaskType{ID: 1})
	before := task.UpdatedAt
	task.AssignTo(42)
	if task.ResponsibleID != 42 {
		t.Fatalf("expected ResponsibleID 42 got %d", task.ResponsibleID)
	}
	if !task.UpdatedAt.After(before) {
		t.Fatalf("expected UpdatedAt to be updated")
	}
}

func TestTask_ChangeStatusMarksCompletedWhenLast(t *testing.T) {
	// create statuses
	s1 := todo.NewTaskStatus("First", true)
	s1.ID = 10
	s2 := todo.NewTaskStatus("Last", true)
	s2.ID = 20

	wf := todo.NewWorkflow("w", map[uint8]todo.TaskStatus{0: s1, 1: s2}, todo.User{ID: 1})

	task := todo.NewTask("t", "d", 1, time.Now().Add(24*time.Hour), todo.TaskType{ID: 1})
	task.Workflow = wf

	task.ChangeStatus(s2)
	if !task.Completed {
		t.Fatalf("expected task to be marked completed when status is last")
	}
}
