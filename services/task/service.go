package task

import (
	Model "awesomeProject/models"
	UserService "awesomeProject/services/user"
	Store "awesomeProject/store/task"
	"errors"
	"log"
)

type Service struct {
	store       *Store.Store
	userService *UserService.Service
}

func New(store *Store.Store, userService *UserService.Service) *Service {
	return &Service{
		store:       store,
		userService: userService,
	}
}

func (s *Service) AddTask(task string, uid int) error {
	if task == "" {
		return errors.New("task is Empty")
	}

	check := s.userService.CheckUserID(uid)

	if check == true {
		return s.store.AddTask(task, uid)
	}
	return errors.New("no User found")
}

func (s *Service) ViewTask() ([]Model.Tasks, error) {

	return s.store.ViewTask()
}

func (s *Service) GetByID(i int) (Model.Tasks, error) {

	if s.store.CheckIfExists(i) {
		ans, err := s.store.GetByID(i)
		if err != nil {
			log.Printf("Error in SERVICES.GetByID: %v", err)
			return Model.Tasks{}, err
		}
		return ans, nil
	}
	return Model.Tasks{}, errors.New("ID not found")
}

func (s *Service) UpdateTask(i int) (bool, error) {
	if s.store.CheckIfExists(i) {
		ans, err := s.store.UpdateTask(i)
		if err != nil {
			log.Printf("Error in SERVICES.UpdateTask: %v", err)
			return false, err
		}
		return ans, nil
	}
	return false, errors.New("ID not found")
}

func (s *Service) DeleteTask(i int) (bool, error) {
	if s.store.CheckIfExists(i) {
		ans, err := s.store.DeleteTask(i)
		if err != nil {
			log.Printf("Error in SERVICES.DeleteTask: %v", err)
			return false, err
		}
		return ans, nil
	}
	return false, errors.New("ID not found")
}
