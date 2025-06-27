package user

import (
	Model "awesomeProject/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type UserService interface {
	AddUser(name string) error
	ViewTask() ([]Model.User, error)
	GetUserId(id int) (Model.User, error)
	CheckUserID(id int) bool
}

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
		fmt.Fprintf(w, "Error Reading Body: %s\n", err)
		return
	}

	var input struct {
		T string `json:"name"`
	}

	err = json.Unmarshal(msg, &input)
	if err != nil {
		fmt.Fprintf(w, "Error Parsing Body: %s\n", err)
		return
	}

	err = h.service.AddUser(input.T)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error Adding User: %s\n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s added.\n", input.T)
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	index, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error in USER_HANDLER: GetUserByID : %v", err)
		return
	}

	ans, err := h.service.GetUserId(index)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error in USER_HANDLER: GetUserByID : %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID : %v,  Name: %v", ans.UserID, ans.Name)

}

func (h *handler) Viewuser(w http.ResponseWriter, r *http.Request) {

	ans, err := h.service.ViewTask()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("USER_HANDLER:VIEW : %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, v := range ans {
		fmt.Fprintf(w, "ID : %v,  Name: %v\n", v.UserID, v.Name)
	}

}
