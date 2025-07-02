package task

import (
	Models "awesomeProject/models"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"strconv"
)

type Handler struct {
	service TaskService
}

// New creates a new task handler.
func New(service TaskService) *Handler {
	return &Handler{service: service}
}

// Addtask godoc
// @Summary Add a new task
// @Description Adds a task to the database for a given user
// @Tags task
// @Accept json
// @Produce plain
// @Param task body Models.AddTaskRequest true "Task input"
// @Success 201 {string} string "Task added"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [post]
func (h *Handler) Addtask(ctx *gofr.Context) (any, error) {
	var reqBody Models.AddTaskRequest

	err := ctx.Bind(&reqBody)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"Give Correct Input"}}
	}

	err = h.service.AddTask(ctx, reqBody.Task, reqBody.UserID)

	if err != nil {
		return nil, err
	}

	return "Task added", nil
}

// Viewtask godoc
// @Summary View all tasks
// @Description Returns a list of all tasks
// @Tags task
// @Produce json
// @Success 200 {array} Models.Tasks
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [get]
func (h *Handler) Viewtask(ctx *gofr.Context) (any, error) {
	ans, err := h.service.ViewTask(ctx)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

// Gettask godoc
// @Summary Get task by ID
// @Description Retrieves task details by ID
// @Tags task
// @Produce plain
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task details"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /task/{id} [get]
func (h *Handler) Gettask(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Request.PathParam("id"))
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"Invalid Param"}}
	}

	ans, err := h.service.GetByID(ctx, id)
	if err != nil {
		return Models.Tasks{}, err
	}

	return ans, nil
}

// Updatetask godoc
// @Summary Update task status
// @Description Updates task status to complete/incomplete
// @Tags task
// @Produce plain
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task updated"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /task/{id} [put]
func (h *Handler) Updatetask(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Request.PathParam("id"))
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"Invalid Param"}}
	}

	_, err = h.service.UpdateTask(ctx, id)

	if err != nil {
		return nil, err
	}

	return "Task updated", nil
}

// Deletetask godoc
// @Summary Delete task
// @Description Deletes a task by its ID
// @Tags task
// @Produce plain
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /task/{id} [delete]
func (h *Handler) Deletetask(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Request.PathParam("id"))
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"Invalid Param"}}
	}

	_, err = h.service.DeleteTask(ctx, id)

	if err != nil {
		return nil, err
	}

	return "Task deleted", nil
}
