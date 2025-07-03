package user

import (
	models "awesomeProject/models"
	"gofr.dev/pkg/gofr"
)

type UserStore interface {
	AddUser(ctx *gofr.Context, name string) error
	GetUserByID(ctx *gofr.Context, id int) (models.User, error)
	ViewUser(ctx *gofr.Context) ([]models.User, error)
	CheckUserID(ctx *gofr.Context, id int) bool
	CheckIfRowsExists(ctx *gofr.Context) bool
}
