package entities

type TaskStatus struct {
	ID     int64
	Label  string
	Active bool
}

func NewTaskStatus(label string, active bool) TaskStatus {
	return TaskStatus{
		Label:  label,
		Active: active,
	}
}

func (self *TaskStatus) Deactivate() {
	self.Active = false
}

func (self *TaskStatus) Activate() {
	self.Active = true
}

func (self *TaskStatus) Rename(newLabel string) {
	self.Label = newLabel
}
