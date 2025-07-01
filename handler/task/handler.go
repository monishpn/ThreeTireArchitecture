package task

import (
	Models "awesomeProject/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
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
func (h *Handler) Addtask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	msg, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	var reqBody Models.AddTaskRequest

	err = json.Unmarshal(msg, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	err = h.service.AddTask(reqBody.Task, reqBody.UserID)

	if err != nil {
		cErr := err.(Models.CustomError)

		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task added"))
}

// Viewtask godoc
// @Summary View all tasks
// @Description Returns a list of all tasks
// @Tags task
// @Produce json
// @Success 200 {array} Models.Tasks
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [get]
func (h *Handler) Viewtask(w http.ResponseWriter, _ *http.Request) {
	ans, err := h.service.ViewTask()
	if err != nil {
		cErr := err.(Models.CustomError)

		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))

		return
	}

	b, _ := json.Marshal(ans)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
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
func (h *Handler) Gettask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	index, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Printf("%s", err.Error())

		return
	}

	var ans Models.Tasks

	ans, err = h.service.GetByID(index)
	if err != nil {
		cErr := err.(Models.CustomError)

		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ans.String()))
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
func (h *Handler) Updatetask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	index, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Printf("%s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	var ans bool
	ans, err = h.service.UpdateTask(index)

	if err != nil {
		cErr := err.(Models.CustomError)

		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))

		return
	}

	if ans {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task updated"))
	}
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
func (h *Handler) Deletetask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	index, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%s", err.Error())

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	var ans bool
	ans, err = h.service.DeleteTask(index)

	if err != nil {
		cErr := err.(Models.CustomError)

		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))

		return
	}

	if ans {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task deleted"))
	}
}
