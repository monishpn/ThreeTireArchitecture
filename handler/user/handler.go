package user

import (
	Model "awesomeProject/models"
	"gofr.dev/pkg/gofr"
	gofrHttp "gofr.dev/pkg/gofr/http"
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
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"Give Correct Input"}}
	}

	err = h.service.AddUser(ctx, reqBody.T)
	if err != nil {
		return nil, err
	}

	return "User Created", nil
}

func (h *handler) GetUserByID(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Request.PathParam("id"))
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"Invalid Param"}}
	}

	ans, err := h.service.GetUserId(ctx, id)

	if err != nil {
		return Model.User{}, err
	}

	return ans, nil
}

func (h *handler) Viewuser(ctx *gofr.Context) (any, error) {
	ans, err := h.service.ViewTask(ctx)
	if err != nil {
		return nil, err
	}

	return ans, nil
}
