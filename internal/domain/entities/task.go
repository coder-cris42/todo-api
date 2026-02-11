package entities

import "time"

type Task struct {
	ID            int64
	Title         string
	Description   string
	Status        TaskStatus
	Parent        *Task
	AuthorID      int64
	Deadline      DateTime
	CreatedAt     DateTime
	UpdatedAt     DateTime
	ResponsibleID int64
	Workflow      Workflow
	Type          TaskType
	Completed     bool
}

func NewTask(title string, description string, authorID int64, deadline time.Time, taskType TaskType) *Task {
	return &Task{
		Title:       title,
		Description: description,
		AuthorID:    authorID,
		Deadline:    NewDateTime(deadline),
		Type:        taskType,
		CreatedAt:   Now(),
	}
}

func (self *Task) IsOverdue() bool {
	return !self.Completed && time.Now().After(self.Deadline.Time)
}

func (self *Task) AssignTo(userID int64) {
	self.ResponsibleID = userID
	self.UpdatedAt = Now()
}

func (self *Task) ChangeStatus(newStatus TaskStatus) {
	self.Status = newStatus
	if self.isTheLastStatus() {
		self.Completed = true
	}
	self.UpdatedAt = Now()
}

func (self *Task) IsCompleted() bool {
	return self.Completed
}

func (self *Task) isTheLastStatus() bool {

	var lastStatus uint8 = 0
	for key := range self.Workflow.Statuses {
		if key > lastStatus {
			lastStatus = key
		}
	}

	return self.Status.ID == self.Workflow.Statuses[lastStatus].ID

}

func (self *Task) SetWorkflow(newWorkflow Workflow) {
	self.Workflow = newWorkflow
	self.UpdatedAt = Now()
}
