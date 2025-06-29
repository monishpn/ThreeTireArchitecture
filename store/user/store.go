package user

import (
	Models "awesomeProject/models"
	"database/sql"
	"log"
	"net/http"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddUser(name string) error {
	_, err := s.db.Exec("insert into USERS (name) values (?)", name)
	if err != nil {
		return Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Adding the Data to the Database"}

	}
	log.Printf("Added user %s", name)
	return nil
}

func (s *Store) GetUserByID(id int) (Models.User, error) {

	var uid int
	var name string

	err := s.db.QueryRow("select * from USERS where uid=?", id).Scan(&uid, &name)
	if err != nil {
		return Models.User{}, Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"}

	}
	return Models.User{uid, name}, nil

}

func (s *Store) ViewUser() ([]Models.User, error) {
	var users []Models.User
	row, err := s.db.Query("Select * from USERS")
	if err != nil {
		return users, Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"}

	}

	defer row.Close()
	var uid int
	var name string

	for row.Next() {
		err = row.Scan(&uid, &name)
		if err != nil {
			return []Models.User{}, Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While reading the data in row "}

		}
		users = append(users, Models.User{uid, name})
	}

	return users, nil
}

func (s *Store) CheckUserID(id int) bool {
	var uid int
	err := s.db.QueryRow("select uid from USERS where uid=?", id).Scan(&uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("DB Error: %v", err)
		return false

	}
	return true
}

func (s *Store) CheckIfRowsExists() bool {
	var num int
	s.db.QueryRow("Select COUNT(*) from USERS").Scan(&num)
	if num > 0 {
		return true
	}
	return false
}
