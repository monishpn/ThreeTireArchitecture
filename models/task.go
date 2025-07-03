package models

import "fmt"

type Tasks struct {
	Tid       int
	Task      string
	Completed bool
	UserID    int
}

type AddTaskRequest struct {
	Task   string `json:"task"`
	UserID int    `json:"userID"`
}

func (t Tasks) String() string {
	return fmt.Sprintf("ID: %d, Task: %s, Status: %v, User ID: %d", t.Tid, t.Task, t.Completed, t.UserID)
}
