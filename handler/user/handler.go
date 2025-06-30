package user

import (
	Model "awesomeProject/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type handler struct {
	service UserService
}

func New(service UserService) *handler {
	return &handler{service}
}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	msg, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Error Reading Body: %s\n", err)
		return
	}

	var input struct {
		T string `json:"name"`
	}

	err = json.Unmarshal(msg, &input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Error Parsing Body: %s\n", err)
		return
	}

	err = h.service.AddUser(input.T)
	if err != nil {
		cErr, _ := err.(Model.CustomError)
		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created"))
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	index, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ans, err := h.service.GetUserId(index)

	if err != nil {
		cErr, _ := err.(Model.CustomError)
		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ans.String()))

}

func (h *handler) Viewuser(w http.ResponseWriter, r *http.Request) {

	ans, err := h.service.ViewTask()
	if err != nil {
		cErr, _ := err.(Model.CustomError)

		w.WriteHeader(cErr.Code)
		w.Write([]byte(cErr.Message))
		return

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Something went wrong!`))
		return

	}

	b, _ := json.Marshal(ans)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
