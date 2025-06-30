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
		cErr := err.(Models.CustomError)
		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task added"))
	return

}

func (h *Handler) Viewtask(w http.ResponseWriter, r *http.Request) {
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
