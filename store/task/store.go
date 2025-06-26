package task

import (
	"awesomeProject/models"
	"database/sql"
)

type store struct {
	db *sql.DB
}

// New creates a new task store
func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s store) Create(task models.Task) (models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetByID(id int) (models.Task, error) {
	//TODO implement me
	panic("implement me")

	task.Create
}
