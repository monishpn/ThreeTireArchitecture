package task

import (
	Model "awesomeProject/models"
	"net/http"
)

type Service struct {
	store       TaskStore
	userService UserService
}

func New(store TaskStore, userService UserService) *Service {
	return &Service{
		store:       store,
		userService: userService,
	}
}

func (s *Service) AddTask(task string, uid int) error {
	if task == "" {
		return Model.CustomError{http.StatusBadRequest, "Task is Empty"}
	}

	check := s.userService.CheckUserID(uid)

	if check {
		return s.store.AddTask(task, uid)
	}
	return Model.CustomError{http.StatusBadRequest, "No user found"}
}

func (s *Service) ViewTask() ([]Model.Tasks, error) {
	return s.store.ViewTask()
}

func (s *Service) GetByID(i int) (Model.Tasks, error) {
	if s.store.CheckIfExists(i) {
		return s.store.GetByID(i)
	}
	return Model.Tasks{}, Model.CustomError{http.StatusBadRequest, "No task found"}
}

func (s *Service) UpdateTask(i int) (bool, error) {
	if s.store.CheckIfExists(i) {
		return s.store.UpdateTask(i)
	}
	return false, Model.CustomError{http.StatusBadRequest, "No task found"}
}

func (s *Service) DeleteTask(i int) (bool, error) {
	if s.store.CheckIfExists(i) {
		return s.store.DeleteTask(i)
	}
	return false, Model.CustomError{http.StatusBadRequest, "No task found"}
}
