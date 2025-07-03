package models

import "fmt"

type Tasks struct {
	Tid       int    `json:"tid"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"uid"`
}

func (t Tasks) String() string {
	return fmt.Sprintf("ID: %d, Task: %s, Status: %v, User ID: %d", t.Tid, t.Task, t.Completed, t.UserID)
}
