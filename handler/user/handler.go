package user

import (
	"awesomeProject/models"
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
	var reqBody models.User

	err := ctx.Bind(&reqBody)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"body"}}
	}

	err = h.service.AddUser(ctx, reqBody.Name)
	if err != nil {
		return nil, err
	}

	return "User Created", nil
}

func (h *handler) GetUserByID(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Request.PathParam("id"))
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"id"}}
	}

	ans, err := h.service.GetUserId(ctx, id)

	if err != nil {
		return nil, err
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
