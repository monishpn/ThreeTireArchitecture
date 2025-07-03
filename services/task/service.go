package task

import (
	models "awesomeProject/models"
	"gofr.dev/pkg/gofr"
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

func (s *Service) AddTask(ctx *gofr.Context, task string, uid int) error {
	if task == "" {
		return models.CustomError{http.StatusBadRequest, "Task is Empty"}
	}

	check := s.userService.CheckUserID(ctx, uid)

	if check {
		return s.store.AddTask(ctx, task, uid)
	}
	return models.CustomError{http.StatusBadRequest, "No user found"}
}

func (s *Service) ViewTask(ctx *gofr.Context) ([]models.Tasks, error) {
	return s.store.ViewTask(ctx)
}

func (s *Service) GetByID(ctx *gofr.Context, i int) (models.Tasks, error) {
	if s.store.CheckIfExists(ctx, i) {
		return s.store.GetByID(ctx, i)
	}
	return models.Tasks{}, models.CustomError{http.StatusBadRequest, "No task found"}
}

func (s *Service) UpdateTask(ctx *gofr.Context, i int) (bool, error) {
	if s.store.CheckIfExists(ctx, i) {
		return s.store.UpdateTask(ctx, i)
	}
	return false, models.CustomError{http.StatusBadRequest, "No task found"}
}

func (s *Service) DeleteTask(ctx *gofr.Context, i int) (bool, error) {
	if s.store.CheckIfExists(ctx, i) {
		return s.store.DeleteTask(ctx, i)
	}
	return false, models.CustomError{http.StatusBadRequest, "No task found"}
}
