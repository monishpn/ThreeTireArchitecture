package user

import (
	models "awesomeProject/models"
	"gofr.dev/pkg/gofr"
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

func (s *Service) AddUser(ctx *gofr.Context, name string) error {
	if name == "" {
		return models.CustomError{http.StatusBadRequest, "Empty String given as input"}
	}

	return s.store.AddUser(ctx, name)
}

func (s *Service) ViewTask(ctx *gofr.Context) (models.UserSlice, error) {
	if s.store.CheckIfRowsExists(ctx) {
		return s.store.ViewUser(ctx)
	}

	return models.UserSlice{}, models.CustomError{http.StatusNoContent, "No user Found"}
}

func (s Service) GetUserId(ctx *gofr.Context, id int) (any, error) {
	if s.CheckUserID(ctx, id) {
		return s.store.GetUserByID(ctx, id)
	}

	return nil, models.CustomError{Code: http.StatusNotFound, Message: "user does not exists"}
}

func (s Service) CheckUserID(ctx *gofr.Context, id int) bool {
	return s.store.CheckUserID(ctx, id)
}
