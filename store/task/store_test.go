package task

import (
	"awesomeProject/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func mockAllocation(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	return sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
}

func TestAddTask(t *testing.T) {
	db, mock, _ := mockAllocation(t)
	defer db.Close()

	mock.ExpectExec("Insert into TASKS (task,completed,uid) values (?,?,?)").
		WithArgs("Testing", false, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	svc := New(db)

	err := svc.AddTask("Testing", 1)
	if err != nil {
		t.Errorf("Error while adding Task : %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Error While checking ecpectation in AddTask : %v", err)
	}

	// Error Checks
	mock.ExpectExec("Insert into TASKS (task,completed,uid) values (?,?)").WithArgs("testing").WillReturnResult(sqlmock.NewResult(0, 1))

	err = svc.AddTask("Testing", 1)

	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestViewTask(t *testing.T) {
	db, mock, _ := mockAllocation(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "task", "completed", "uid"}).
		AddRow(1, "Task 1", false, 101).
		AddRow(2, "Task 2", true, 102)

	mock.ExpectQuery("select * from TASKS").
		WillReturnRows(rows)

	exptedTasks := []models.Tasks{
		{1, "Task 1", false, 101},
		{2, "Task 2", true, 102},
	}

	store := New(db)
	tasks, err := store.ViewTask()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if err != nil {
		t.Errorf("Error while Checking Expectations in ViewTask : %v", err)
	}

	if !reflect.DeepEqual(exptedTasks, tasks) {
		t.Errorf("Some error in the output")
	}

	// Error Check
	mock.ExpectQuery("Select * from Tasks").WillReturnRows(rows)

	_, err = store.ViewTask()

	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestGetByID(t *testing.T) {
	db, mock, _ := mockAllocation(t)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "task", "completed", "uid"}).
		AddRow(1, "Task 1", false, 101)

	mock.ExpectQuery("select * from TASKS where id=?").
		WithArgs(1).
		WillReturnRows(row)

	expected := models.Tasks{1, "Task 1", false, 101}
	store := New(db)
	tasks, err := store.GetByID(1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(expected, tasks) {
		t.Errorf("Some error in the output")
	}

	// error check
	mock.ExpectQuery("select * from Task where uid").WillReturnRows(row)

	_, err = store.GetByID(1)

	if err == nil {
		t.Errorf("Expected Error but got none")
	}
}

func TestUpdateTask(t *testing.T) {
	db, mock, _ := mockAllocation(t)
	defer db.Close()

	mock.ExpectExec("UPDATE TASKS SET completed= true WHERE id=?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	store := New(db)
	ok, err := store.UpdateTask(1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if !ok {
		t.Errorf("Expected true, got false")
	}

	// error check
	mock.ExpectExec("UPDATE task SET completed= true WHERE id").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = store.UpdateTask(1)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestDeleteTask(t *testing.T) {
	db, mock, _ := mockAllocation(t)
	defer db.Close()

	mock.ExpectExec("delete from TASKS where id=?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	store := New(db)
	ok, err := store.DeleteTask(1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if !ok {
		t.Errorf("Expected true, got false")
	}

	// error check
	mock.ExpectExec("delete from TASKS where id").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = store.DeleteTask(1)

	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestCheckIfExists(t *testing.T) {
	db, mock, _ := mockAllocation(t)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("select id from TASKS where id=?").
		WithArgs(1).
		WillReturnRows(row)

	store := New(db)
	exists := store.CheckIfExists(1)

	if !exists {
		t.Errorf("Expected true, got false")
	}

	// Error check empty row
	row = sqlmock.NewRows([]string{"id"})

	mock.ExpectQuery("select id from TASKS where id=?").
		WithArgs(1).
		WillReturnRows(row)

	exists = store.CheckIfExists(1)
	if exists {
		t.Errorf("Expected false, got true")
	}
}
