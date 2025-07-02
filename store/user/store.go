package user

import (
	Models "awesomeProject/models"
	"database/sql"
	"gofr.dev/pkg/gofr"
	"log"
	"net/http"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddUser(ctx *gofr.Context, name string) error {
	_, err := ctx.SQL.ExecContext(ctx, "insert into USERS (name) values (?)", name)
	if err != nil {
		log.Println(err)
		return Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Adding the Data to the Database"}
	}

	log.Printf("Added user %s", name)

	return nil
}

func (s *Store) GetUserByID(ctx *gofr.Context, id int) (Models.User, error) {
	var uid int

	var name string

	err := ctx.SQL.QueryRowContext(ctx, "select * from USERS where uid=?", id).Scan(&uid, &name)
	if err != nil {
		log.Println(err)
		return Models.User{},
			Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"}
	}

	return Models.User{uid, name}, nil
}

func (s *Store) ViewUser(ctx *gofr.Context) ([]Models.User, error) {
	var users []Models.User
	row, err := ctx.SQL.QueryContext(ctx, "Select * from USERS")

	if err != nil {
		return []Models.User{}, Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"}
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

func (s *Store) CheckUserID(ctx *gofr.Context, id int) bool {
	var uid int
	err := ctx.SQL.QueryRowContext(ctx, "select uid from USERS where uid=?", id).Scan(&uid)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		log.Printf("DB Error: %v", err)

		return false
	}

	return true
}

func (s *Store) CheckIfRowsExists(ctx *gofr.Context) bool {
	var num int
	err := ctx.SQL.QueryRowContext(ctx, "Select COUNT(*) from USERS").Scan(&num)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		return false
	}

	return num > 0
}
