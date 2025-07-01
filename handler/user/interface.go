package user

import (
	Model "awesomeProject/models"
	"gofr.dev/pkg/gofr"
)

type UserService interface {
	AddUser(ctx *gofr.Context, name string) error
	ViewTask(ctx *gofr.Context) (Model.UserSlice, error)
	GetUserId(ctx *gofr.Context, id int) (Model.User, error)
}
