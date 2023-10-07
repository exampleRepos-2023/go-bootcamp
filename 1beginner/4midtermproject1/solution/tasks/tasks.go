package tasks

type TaskType int

type TaskStatus string

const (
	NewTaskID int        = -1
	Open      TaskStatus = "OPEN"
	InProgress TaskStatus = "IN_PROGRESS"
	Resolved TaskStatus = "RESOLVED"
	Closed   TaskStatus = "CLOSED"
)

type Epic struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status   TaskStatus `json:"status"`
	StoryIDs []int      `json:"story_ids"`
}

type Story struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status TaskStatus `json:"status"`
}

type Tasks struct {
	LastTaskID int          `json:"last_task_id"`
	Epics      map[int]Epic `json:"epics"`
	Stories map[int]Story `json:"stories"`
}