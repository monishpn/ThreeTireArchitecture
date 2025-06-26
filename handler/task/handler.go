package task

import (
	"awesomeProject/models"
	"net/http"
)

type Store interface {
	Create(task models.Task) (models.Task, error)
	GetByID(id int) (models.Task, error)
}

type handler struct {
	str Store
}

// New creates a new task handler
func New(str Store) *handler {
	return &handler{str: str}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) (models.Task, error) {
	// get the body, pathParam, queryParam, etc.

	// intialise the task struct

	t := models.Task{}

	// empty value checks
	// nil checks
	//
	t.Validate()

	return h.str.Create(t)
}
