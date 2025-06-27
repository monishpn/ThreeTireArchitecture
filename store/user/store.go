package user

import (
	"awesomeProject/models"
	"database/sql"
	"log"
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
		return err
	}
	log.Printf("Added user %s", name)
	return nil
}

func (s *Store) GetUserByID(id int) (models.User, error) {

	var uid int
	var name string

	err := s.db.QueryRow("select * from USERS where uid=?", id).Scan(&uid, &name)
	if err != nil {
		return models.User{}, err
	}
	return models.User{uid, name}, nil

}

func (s *Store) ViewUser() ([]models.User, error) {
	var users []models.User
	row, err := s.db.Query("Select * from USERS")
	if err != nil {
		return users, err
	}

	defer row.Close()
	var uid int
	var name string

	for row.Next() {
		err = row.Scan(&uid, &name)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, models.User{uid, name})
	}
	return users, nil
}

func (s *Store) CheckUserID(id int) bool {
	var uid int
	err := s.db.QueryRow("select * form USERS where uid=?", id).Scan(&uid)
	if err != nil {
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
