package tasks

type TaskType int

type TaskStatus string

const (
	NewTaskID int = -1
)

type Epic struct {
	// Todo
}

type Story struct {
	// Todo
}

type Tasks struct {
	LastTaskID int           `json:"last_task_id"`
	Epics      map[int]Epic  `json:"epics"`
	Stories    map[int]Story `json:"stories"`
}
