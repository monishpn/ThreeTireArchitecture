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
		return models.CustomError{Code: http.StatusBadRequest, Message: "Task is Empty"}
	}

	check := s.userService.CheckUserID(ctx, uid)

	if check {
		return s.store.AddTask(ctx, task, uid)
	}
	return models.CustomError{Code: http.StatusBadRequest, Message: "No user found"}
}

func (s *Service) ViewTask(ctx *gofr.Context) ([]models.Tasks, error) {
	return s.store.ViewTask(ctx)
}

func (s *Service) GetByID(ctx *gofr.Context, i int) (any, error) {
	if s.store.CheckIfExists(ctx, i) {
		return s.store.GetByID(ctx, i)
	}
	return nil, models.CustomError{Code: http.StatusBadRequest, Message: "No task found"}
}

func (s *Service) UpdateTask(ctx *gofr.Context, i int) error {
	if s.store.CheckIfExists(ctx, i) {
		return s.store.UpdateTask(ctx, i)
	}
	return models.CustomError{Code: http.StatusBadRequest, Message: "No task found"}
}

func (s *Service) DeleteTask(ctx *gofr.Context, i int) error {
	if s.store.CheckIfExists(ctx, i) {
		return s.store.DeleteTask(ctx, i)
	}
	return models.CustomError{Code: http.StatusBadRequest, Message: "No task found"}
}
