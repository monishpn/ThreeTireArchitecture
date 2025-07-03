package user

import (
	"awesomeProject/models"
	"gofr.dev/pkg/gofr"
)

type UserService interface {
	AddUser(ctx *gofr.Context, name string) error
	ViewTask(ctx *gofr.Context) (models.UserSlice, error)
	GetUserId(ctx *gofr.Context, id int) (models.User, error)
}
