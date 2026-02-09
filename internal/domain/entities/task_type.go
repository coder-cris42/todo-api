package entities

type TaskType struct {
	ID   int64
	Name string
}

func NewTaskType(name string) TaskType {
	return TaskType{
		Name: name,
	}
}
