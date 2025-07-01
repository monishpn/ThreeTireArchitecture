package task

import (
	Model "awesomeProject/models"
	"gofr.dev/pkg/gofr"
)

type TaskStore interface {
	AddTask(ctx *gofr.Context, task string, uid int) error
	ViewTask(ctx *gofr.Context) ([]Model.Tasks, error)
	GetByID(ctx *gofr.Context, id int) (Model.Tasks, error)
	UpdateTask(ctx *gofr.Context, id int) (bool, error)
	DeleteTask(ctx *gofr.Context, id int) (bool, error)
	CheckIfExists(ctx *gofr.Context, i int) bool
}

type UserService interface {
	CheckUserID(ctx *gofr.Context, id int) bool
}
