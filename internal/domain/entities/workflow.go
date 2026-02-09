package entities

import (
	"fmt"
	"time"
)

type Workflow struct {
	ID        int64
	Name      string
	Statuses  map[uint8]TaskStatus
	Author    User
	CreatedAt time.Time
}

func NewWorkflow(name string, statuses map[uint8]TaskStatus, author User) Workflow {
	return Workflow{
		Name:      name,
		Statuses:  statuses,
		Author:    author,
		CreatedAt: time.Now(),
	}
}

func (self *Workflow) AddStatus(key uint8, status TaskStatus) error {

	err := self.ValidateStatus(status)
	if err != nil {
		return err
	}
	self.Statuses[key] = status
	return nil
}

func (self *Workflow) RemoveStatus(key uint8) error {
	delete(self.Statuses, key)
	return nil
}

func (self *Workflow) ValidateStatus(status TaskStatus) error {

	if !status.Active {
		return fmt.Errorf("Staus %s is not active", status.Label)
	}

	for _, s := range self.Statuses {
		if s.ID == status.ID {
			return fmt.Errorf("Status %s already exists in the workflow", status.Label)
		}

	}
	return nil
}
