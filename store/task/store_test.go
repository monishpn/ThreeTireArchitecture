package task

import (
	Models "awesomeProject/models"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func TestAddTask(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("Insert into TASKS").
		WithArgs("Do something", false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := New(db)
	err := store.AddTask("Do something", 1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestViewTask(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "task", "completed", "uid"}).
		AddRow(1, "Task 1", false, 101).
		AddRow(2, "Task 2", true, 102)

	mock.ExpectQuery("select .* from TASKS").
		WillReturnRows(rows)

	store := New(db)
	tasks, err := store.ViewTask()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expected := []Models.Tasks{
		{1, "Task 1", false, 101},
		{2, "Task 2", true, 102},
	}
	if !reflect.DeepEqual(tasks, expected) {
		t.Errorf("Expected: %v, got: %v", expected, tasks)
	}
}

func TestGetByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "task", "completed", "uid"}).
		AddRow(1, "Task 1", false, 101)

	mock.ExpectQuery("select .* from TASKS where id").
		WithArgs(1).
		WillReturnRows(row)

	store := New(db)
	task, err := store.GetByID(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	expected := Models.Tasks{1, "Task 1", false, 101}
	if !reflect.DeepEqual(task, expected) {
		t.Errorf("Expected: %v, got: %v", expected, task)
	}
}

func TestUpdateTask(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("UPDATE TASKS SET completed= true").
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
}

func TestDeleteTask(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("delete from TASKS where id").
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
}

func TestCheckIfExists(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	row := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("select id from TASKS where id").
		WithArgs(1).
		WillReturnRows(row)

	store := New(db)
	exists := store.CheckIfExists(1)
	if !exists {
		t.Errorf("Expected true, got false")
	}
}
