package task

import (
	models "awesomeProject/models"
	"database/sql"
	"gofr.dev/pkg/gofr"
	"log"
	"net/http"
)

type Store struct {
	db *sql.DB
}

// New creates a new task store
func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddTask(ctx *gofr.Context, task string, uid int) error {
	_, err := ctx.SQL.ExecContext(ctx, "Insert into TASKS (task,completed,uid) values (?,?,?)", task, false, uid)
	if err != nil {
		log.Printf("Error in Task/STORE.AddTask: %v", err)
		return models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Adding the Data to the Database"}
	}

	return nil
}

func (_ *Store) ViewTask(ctx *gofr.Context) ([]models.Tasks, error) {
	var tID int

	var task string

	var completed bool

	var uid int

	var answers []models.Tasks

	row, err := ctx.SQL.QueryContext(ctx, "select * from TASKS")
	if err != nil {
		log.Printf("Error in Task/STORE.View: %v", err)

		return []models.Tasks{}, models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Error While retrieving the Data from the Database",
		}
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&tID, &task, &completed, &uid)
		if err != nil {
			log.Printf("Error in Task/STORE.View: %v", err)
			return []models.Tasks{}, models.CustomError{Code: http.StatusInternalServerError, Message: "Error While reading the data in row "}
		}
		answers = append(answers, models.Tasks{tID, task, completed, uid})
	}

	return answers, nil
}

func (_ *Store) GetByID(ctx *gofr.Context, id int) (models.Tasks, error) {
	var tID int

	var task string

	var completed bool

	var uid int

	err := ctx.SQL.QueryRowContext(ctx, "select * from TASKS where id=?", id).Scan(&tID, &task, &completed, &uid)
	if err != nil {
		log.Printf("Error in Task/STORE.GetByID: %v", err)
		return models.Tasks{},
			models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"}
	}
	return models.Tasks{tID, task, completed, uid}, nil
}

func (s *Store) UpdateTask(ctx *gofr.Context, id int) error {
	_, err := ctx.SQL.ExecContext(ctx, "UPDATE TASKS SET completed= true WHERE id=?", id)
	if err != nil {
		log.Printf("Error in Task/STORE.UpdateTask: %v", err)
		return models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Updating the database "}
	}

	return nil
}

func (s *Store) DeleteTask(ctx *gofr.Context, id int) error {
	_, err := ctx.SQL.ExecContext(ctx, "delete from TASKS where id=?", id)
	if err != nil {
		log.Printf("Error in Task/STORE.DeleteTask: %v", err)
		return models.CustomError{Code: http.StatusInternalServerError, Message: "Error While deleting data in Database"}
	}

	return nil
}

func (s *Store) CheckIfExists(ctx *gofr.Context, i int) bool {
	var id int
	err := ctx.SQL.QueryRowContext(ctx, "select id from TASKS where id=?", i).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		log.Printf("Error in Task/STORE.CheckIfExists: %v", err)

		return false
	}

	return true
}
