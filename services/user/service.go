package user

import (
	Model "awesomeProject/models"
	"net/http"
)

type Service struct {
	store UserStore
}

func New(store UserStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) AddUser(name string) error {
	if name == "" {
		return Model.CustomError{http.StatusBadRequest, "Empty String given as input"}
	}

	return s.store.AddUser(name)
}

func (s *Service) ViewTask() (Model.UserSlice, error) {
	if s.store.CheckIfRowsExists() {
		return s.store.ViewUser()
	}

	return Model.UserSlice{}, Model.CustomError{http.StatusNoContent, "No user Found"}
}

func (s Service) GetUserId(id int) (Model.User, error) {
	if s.CheckUserID(id) {
		return s.store.GetUserByID(id)
	}

	return Model.User{}, Model.CustomError{Code: http.StatusNotFound, Message: "user does not exists"}
}

func (s Service) CheckUserID(id int) bool {
	return s.store.CheckUserID(id)
}
