package user

import (
	Model "awesomeProject/models"
	"awesomeProject/store/user"
	"errors"
	"log"
)

type ServiceInterface interface {
	AddUser(name string) error
	ViewTask() ([]Model.User, error)
	GetUserId(id int) (Model.User, error)
	CheckUserID(id int) bool
}

type Service struct {
	store *user.Store
}

func New(store *user.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) AddUser(name string) error {
	if name == "" {
		return errors.New("Name is Empty")
	}
	return s.store.AddUser(name)
}

func (s *Service) ViewTask() ([]Model.User, error) {
	if s.store.CheckIfRowsExists() {
		return s.store.ViewUser()
	}

	return nil, errors.New("No USERS to display.")

}

func (s Service) GetUserId(id int) (Model.User, error) {
	if s.CheckUserID(id) {
		return s.store.GetUserByID(id)
	}
	return Model.User{}, errors.New("User not found")
}

func (s Service) CheckUserID(id int) bool {

	if s.store == nil {
		log.Printf("ERROR, *Store is a nil\n\n")
		return false
	}
	return s.store.CheckUserID(id)
}
