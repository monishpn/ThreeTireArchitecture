package user

import Model "awesomeProject/models"

type UserService interface {
	AddUser(name string) error
	ViewTask() (Model.UserSlice, error)
	GetUserId(id int) (Model.User, error)
}
