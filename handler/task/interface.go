package task

import Models "awesomeProject/models"

type TaskService interface {
	AddTask(task string, uid int) error
	ViewTask() ([]Models.Tasks, error)
	GetByID(id int) (Models.Tasks, error)
	UpdateTask(id int) (bool, error)
	DeleteTask(id int) (bool, error)
}
