package task

import (
	Model "awesomeProject/models"
	"errors"
	"net/http"
	"reflect"
	"testing"
)

type mockTaskStore struct {
	returnErr     bool
	existsChecker bool
}

func (m mockTaskStore) AddTask(task string, uid int) error {
	if m.returnErr {
		return Model.CustomError{http.StatusBadRequest, "DB AddTask Error"}
	}
	return nil
}

func (m mockTaskStore) ViewTask() ([]Model.Tasks, error) {
	if m.returnErr {
		return nil, Model.CustomError{http.StatusInternalServerError, "DB ViewTask Error"}
	}
	return []Model.Tasks{{Tid: 1, Task: "Read", Completed: false, UserID: 1}}, nil
}

func (m mockTaskStore) GetByID(id int) (Model.Tasks, error) {
	if m.returnErr {
		return Model.Tasks{}, Model.CustomError{http.StatusNotFound, "DB GetByID Error"}
	}
	return Model.Tasks{Tid: id, Task: "Read", Completed: false, UserID: 1}, nil
}

func (m mockTaskStore) UpdateTask(id int) (bool, error) {
	if m.returnErr {
		return false, Model.CustomError{http.StatusInternalServerError, "DB UpdateTask Error"}
	}
	return true, nil
}

func (m mockTaskStore) DeleteTask(id int) (bool, error) {
	if m.returnErr {
		return false, Model.CustomError{http.StatusInternalServerError, "DB DeleteTask Error"}
	}
	return true, nil
}

func (m mockTaskStore) CheckIfExists(id int) bool {
	return m.existsChecker
}

type mockUserService struct {
	returnValid bool
}

func (m mockUserService) CheckUserID(id int) bool {
	return m.returnValid
}

func TestAddTask(t *testing.T) {
	svc := New(&mockTaskStore{}, &mockUserService{returnValid: true})

	err := svc.AddTask("Read", 1)
	if err != nil {
		t.Errorf("AddTask failed unexpectedly: %v", err)
	}

	// Empty Task
	err = svc.AddTask("", 1)
	exp := Model.CustomError{http.StatusBadRequest, "Task is Empty"}
	if !errors.Is(err, exp) {
		t.Errorf("Expected error for empty task, got: %v", err)
	}

	// Invalid User
	svc = New(&mockTaskStore{}, &mockUserService{returnValid: false})
	err = svc.AddTask("Read", 1)
	exp = Model.CustomError{http.StatusBadRequest, "No user found"}
	if !errors.Is(err, exp) {
		t.Errorf("Expected user error, got: %v", err)
	}
}

func TestViewTask(t *testing.T) {
	exp := []Model.Tasks{{Tid: 1, Task: "Read", Completed: false, UserID: 1}}
	svc := New(&mockTaskStore{}, &mockUserService{})

	got, err := svc.ViewTask()
	if err != nil {
		t.Errorf("ViewTask unexpected error: %v", err)
	}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("Expected: %v, got: %v", exp, got)
	}
}

func TestGetByID(t *testing.T) {
	exp := Model.Tasks{Tid: 1, Task: "Read", Completed: false, UserID: 1}
	svc := New(&mockTaskStore{existsChecker: true}, &mockUserService{})

	got, err := svc.GetByID(1)
	if err != nil {
		t.Errorf("GetByID unexpected error: %v", err)
	}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("Expected: %v, got: %v", exp, got)
	}

	// Task does not exist
	svc = New(&mockTaskStore{existsChecker: false}, &mockUserService{})
	got, err = svc.GetByID(1)
	expErr := Model.CustomError{http.StatusBadRequest, "No task found"}
	if !errors.Is(err, expErr) {
		t.Errorf("Expected: %v, got: %v", expErr, err)
	}
	if got != (Model.Tasks{}) {
		t.Errorf("Expected empty result, got: %v", got)
	}
}

func TestUpdateTask(t *testing.T) {
	svc := New(&mockTaskStore{existsChecker: true}, &mockUserService{})
	ok, err := svc.UpdateTask(1)
	if err != nil || !ok {
		t.Errorf("UpdateTask failed unexpectedly: ok=%v err=%v", ok, err)
	}

	svc = New(&mockTaskStore{existsChecker: false}, &mockUserService{})
	ok, err = svc.UpdateTask(1)
	exp := Model.CustomError{http.StatusBadRequest, "No task found"}
	if !errors.Is(err, exp) {
		t.Errorf("Expected: %v, got: %v", exp, err)
	}
	if ok {
		t.Errorf("Expected false, got: %v", ok)
	}
}

func TestDeleteTask(t *testing.T) {
	svc := New(&mockTaskStore{existsChecker: true}, &mockUserService{})
	ok, err := svc.DeleteTask(1)
	if err != nil || !ok {
		t.Errorf("DeleteTask failed unexpectedly: ok=%v err=%v", ok, err)
	}

	svc = New(&mockTaskStore{existsChecker: false}, &mockUserService{})
	ok, err = svc.DeleteTask(1)
	exp := Model.CustomError{http.StatusBadRequest, "No task found"}
	if !errors.Is(err, exp) {
		t.Errorf("Expected: %v, got: %v", exp, err)
	}
	if ok {
		t.Errorf("Expected false, got: %v", ok)
	}
}
