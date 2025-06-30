package user

import Model "awesomeProject/models"

type UserStore interface {
	AddUser(name string) error
	GetUserByID(id int) (Model.User, error)
	ViewUser() ([]Model.User, error)
	CheckUserID(id int) bool
	CheckIfRowsExists() bool
}
