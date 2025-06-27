package task

import (
	Models "awesomeProject/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type TaskService interface {
	AddTask(task string, uid int) error
	ViewTask() ([]Models.Tasks, error)
	GetByID(id int) (Models.Tasks, error)
	UpdateTask(id int) (bool, error)
	DeleteTask(id int) (bool, error)
}

type Handler struct {
	service TaskService
}

// New creates a new task handler
func New(service TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Addtask(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	msg, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	var reqBody struct {
		T string `json:"task"`
		U int    `json:"userID"`
	}

	err = json.Unmarshal(msg, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	err = h.service.AddTask(reqBody.T, reqBody.U)

	if err != nil {
		fmt.Fprintf(w, "Error in HANDLER.AddTask: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task added"))
	return

}

func (h *Handler) Viewtask(w http.ResponseWriter, r *http.Request) {
	ans, err := h.service.ViewTask()
	if err != nil {
		log.Printf("Error in HANDLER.Viewtask: %v", err)
		return
	}
	for _, v := range ans {
		fmt.Fprintf(w, "ID: %d, Task: %s, Completed: %t, UserID: %v\n", v.Tid, v.Task, v.Completed, v.UserID)
	}
}

func (h *Handler) Gettask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	index, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	var ans Models.Tasks

	ans, err = h.service.GetByID(index)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %d, Task: %s, Completed: %t\n", ans.Tid, ans.Task, ans.Completed)
}

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
		log.Printf("%s", err.Error())
		return
	}
	if ans {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task updated"))
	}
}

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
		log.Printf("%s", err.Error())
		return
	}
	if ans {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task deleted"))

	}
}
