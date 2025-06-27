package user

import (
	Model "awesomeProject/models"
	"awesomeProject/store/user"
	"errors"
)

type Service struct {
	store *user.Store
}

func New(store *user.Store) *Service {
	return &Service{store: store}
}

func (s *Service) Add_User(name string) error {
	if name == "" {
		return errors.New("Name is Empty")
	}
	return s.store.AddUser(name)
}

func (s *Service) View_Task() ([]Model.User, error) {
	if s.store.CheckIfRowsExists() {
		return s.store.ViewUser()
	}

	return nil, errors.New("No USERS to display.")

}

func (s *Service) Get_User_ID(id int) (Model.User, error) {
	if s.store.CheckUserID(id) {
		return s.store.GetUserByID(id)
	}
	return Model.User{}, errors.New("User not found")
}
