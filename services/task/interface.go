package task

import Model "awesomeProject/models"

type TaskStore interface {
	AddTask(task string, uid int) error
	ViewTask() ([]Model.Tasks, error)
	GetByID(id int) (Model.Tasks, error)
	UpdateTask(id int) (bool, error)
	DeleteTask(id int) (bool, error)
	CheckIfExists(i int) bool
}

type UserService interface {
	CheckUserID(id int) bool
}
