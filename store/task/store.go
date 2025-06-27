package task

import (
	Models "awesomeProject/models"
	"database/sql"
	"log"
)

type TaskStore interface {
	AddTask(task string, uid int) error
	ViewTask() ([]Models.Tasks, error)
	GetByID(id int) (Models.Tasks, error)
	UpdateTask(id int) (bool, error)
	DeleteTask(id int) (bool, error)
	CheckIfExists(i int) bool
}

type Store struct {
	db *sql.DB
}

// New creates a new task store
func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddTask(task string, uid int) error {

	_, err := s.db.Exec("Insert into TASKS (task,completed,uid) values (?,?,?)", task, false, uid)
	if err != nil {
		log.Printf("Error in Task/STORE.AddTask: %v", err)
		return err

	}
	return nil
}

func (s *Store) ViewTask() ([]Models.Tasks, error) {

	var tID int
	var task string
	var completed bool
	var uid int

	var answers []Models.Tasks

	row, err := s.db.Query("select * from TASKS")
	if err != nil {
		log.Printf("Error in Task/STORE.View: %v", err)
		return []Models.Tasks{}, err
	}

	defer row.Close()
	for row.Next() {
		err := row.Scan(&tID, &task, &completed, &uid)
		if err != nil {
			log.Printf("Error in Task/STORE.View: %v", err)
			return []Models.Tasks{}, err
		}
		answers = append(answers, Models.Tasks{tID, task, completed, uid})

	}
	return answers, nil
}

func (s *Store) GetByID(id int) (Models.Tasks, error) {

	var tID int
	var task string
	var completed bool
	var uid int

	err := s.db.QueryRow("select * from TASKS where id=?", id).Scan(&tID, &task, &completed, &uid)
	if err != nil {
		log.Printf("Error in Task/STORE.GetByID: %v", err)
		return Models.Tasks{}, err
	}
	return Models.Tasks{tID, task, completed, uid}, nil
}

func (s *Store) UpdateTask(id int) (bool, error) {

	_, err := s.db.Exec("UPDATE TASKS SET completed= true WHERE id=?", id)
	if err != nil {
		log.Printf("Error in Task/STORE.UpdateTask: %v", err)
		return false, err
	}
	return true, nil
}

func (s *Store) DeleteTask(id int) (bool, error) {

	_, err := s.db.Exec("delete from TASKS where id=?", id)
	if err != nil {
		log.Printf("Error in Task/STORE.DeleteTask: %v", err)
		return false, err
	}
	return true, nil
}

func (s *Store) CheckIfExists(i int) bool {
	ans := s.db.QueryRow("select id from TASKS where id=?", i)

	var id int
	err := ans.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("Error in Task/STORE.CheckIfExists: %v", err)
		return false
	}
	return true
}
