package user

import (
	Model "awesomeProject/models"
	"gofr.dev/pkg/gofr"
	"strconv"
)

type handler struct {
	service UserService
}

func New(service UserService) *handler {
	return &handler{service}
}

func (h *handler) AddUser(ctx *gofr.Context) (any, error) {

	var reqBody Model.Input

	err := ctx.Bind(&reqBody)
	if err != nil {
		return nil, err
	}

	err = h.service.AddUser(reqBody.T)
	if err != nil {
		return nil, err
	}

	return "User Created", nil
}

func (h *handler) GetUserByID(ctx *gofr.Context) (any, error) {

	id, err := strconv.Atoi(ctx.Request.PathParam("id"))
	if err != nil {
		return nil, err
	}

	ans, err := h.service.GetUserId(id)

	if err != nil {
		return nil, err
	}

	return ans, nil
}

func (h *handler) Viewuser(ctx *gofr.Context) (any, error) {
	ans, err := h.service.ViewTask()
	if err != nil {
		return nil, err
	}

	return ans, nil
}
