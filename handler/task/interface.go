package task

import (
	"awesomeProject/models"
	"gofr.dev/pkg/gofr"
)

type TaskService interface {
	AddTask(ctx *gofr.Context, task string, uid int) error
	ViewTask(ctx *gofr.Context) ([]models.Tasks, error)
	GetByID(ctx *gofr.Context, id int) (models.Tasks, error)
	UpdateTask(ctx *gofr.Context, id int) (bool, error)
	DeleteTask(ctx *gofr.Context, id int) (bool, error)
}
