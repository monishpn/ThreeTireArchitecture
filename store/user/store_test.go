package user

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func mockAllocation(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	return sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
}
func TestAddUser(t *testing.T) {
	db, mock, err := mockAllocation(t)
	if err != nil {
		t.Errorf("Error while establising SQLmock : %v", err)
	}

	defer db.Close()
	svc := New(db)

	mock.ExpectExec("insert into USERS (name) values (?) ").WithArgs("Ram").WillReturnResult(sqlmock.NewResult(1, 1))

	err = svc.AddUser("Ram")
	if err != nil {
		t.Errorf("Error while adding data : %v", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("Failed AddUser : %v", err)
	}

}

func TestGetUserByID(t *testing.T) {
	db, mock, _ := mockAllocation(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"uid", "name"}).AddRow(1, "Ram")

	mock.ExpectQuery("select * from USERS where uid=?").WillReturnRows(rows)

	svc := New(db)

	_, err := svc.GetUserByID(1)
	if err != nil {
		t.Errorf("Error with database addition in GetUserByID : %v", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("Failed GetUserByID : %v", err)
	}
}

func TestViewUser(t *testing.T) {
	db, mock, _ := mockAllocation(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"uid", "name"}).AddRow(1, "Ram").AddRow(2, "Shyam")

	mock.ExpectQuery("Select * from USERS").WillReturnRows(rows)

	svc := New(db)

	_, err := svc.ViewUser()
	if err != nil {
		t.Errorf("Error with database addition in GetUserByID : %v", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("Failed ViewUser : %v", err)
	}
}

func TestCheckUserID(t *testing.T) {
	db, mock, _ := mockAllocation(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"uid"}).AddRow("Ram")

	mock.ExpectQuery("select uid from USERS where uid=?").WithArgs(1).WillReturnRows(rows)

	svc := New(db)

	_ = svc.CheckUserID(1)

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed CheckUserID : %v", err)
	}
}

func TestCheckIfRowsExists(t *testing.T) {
	db, mock, _ := mockAllocation(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"bum"}).AddRow(1)

	mock.ExpectQuery("Select COUNT(*) from USERS").WillReturnRows(rows)

	svc := New(db)

	_ = svc.CheckIfRowsExists()

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed CheckIfRowsExists : %v", err)
	}
}
