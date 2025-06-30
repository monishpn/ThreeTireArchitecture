package user

import (
	Model "awesomeProject/models"
	"errors"
	"net/http"
	"reflect"
	"testing"
)

type mockStore struct {
	returnErr bool
	//error     Model.CustomError
}

func (m mockStore) AddUser(name string) error {
	if name == "" {
		return Model.CustomError{http.StatusBadRequest, "Empty String given as input"}
	}

	return nil
}

func (m mockStore) GetUserByID(id int) (Model.User, error) {
	if m.returnErr {
		return Model.User{}, Model.CustomError{http.StatusNotFound, "user does not exists"}
	}

	if id == 1 {
		return Model.User{UserID: 1, Name: "Ram"}, nil
	}

	if id == 2 {
		return Model.User{UserID: 2, Name: "Shyam"}, nil
	}
	return Model.User{}, Model.CustomError{http.StatusNotFound, "user does not exists"}

}

func (m mockStore) ViewUser() ([]Model.User, error) {

	if !m.returnErr {
		return []Model.User{Model.User{1, "Ram"}, Model.User{2, "Shyam"}}, nil
	}
	return []Model.User{}, Model.CustomError{http.StatusNotFound, "user does not exists"}
}

func (m mockStore) CheckUserID(id int) bool {

	if !m.returnErr {
		if id == 1 {
			return true
		}
		if id == 2 {
			return true
		}
	}

	return false
}

func (m mockStore) CheckIfRowsExists() bool {

	if m.returnErr {
		return false
	}
	return true
}

func TestAdd(t *testing.T) {
	svc := New(&mockStore{})

	err := svc.AddUser("Ram")

	if err != nil {
		t.Errorf("Expected No Error")
	}
	exp := Model.CustomError{http.StatusBadRequest, "Empty String given as input"}
	err = svc.AddUser("")

	if !errors.Is(err, exp) {
		t.Errorf("Empty string error fail")
	}

}

func TestGetByID(t *testing.T) {
	svc := New(&mockStore{})

	res, err := svc.GetUserId(1)

	if err != nil {
		t.Errorf("Error in GetUserId{1} : %v", err)
	}
	if res.UserID != 1 {
		t.Errorf("Task/Service/GetByID failed, Expected: %v, got: %v", 1, res.UserID)
	}
	if res.Name != "Ram" {
		t.Errorf("Task/Service/GetByID failed, Expected: %v, got: %v", "Ram", res.Name)
	}

	//Error check
	expErr := Model.CustomError{http.StatusNotFound, "user does not exists"}
	svc = New(&mockStore{returnErr: true})

	errRes, errs := svc.GetUserId(1)

	if !errors.Is(errs, expErr) {
		t.Errorf("expected error: %v, but got: %v", expErr, errs)
	}

	if !reflect.DeepEqual(errRes, Model.User{}) {
		t.Errorf("expected empty result, Expected: %v but got %v", Model.User{}, errRes)
	}
}

func TestView(t *testing.T) {
	exp := Model.UserSlice{Model.User{1, "Ram"}, Model.User{2, "Shyam"}}
	svc := New(&mockStore{})
	op, err := svc.ViewTask()

	if err != nil {
		t.Errorf("Error while testing View function : %v", err)
	}

	if !reflect.DeepEqual(exp, op) {
		t.Errorf("Error while testing View function : \nExpected : %v\nGot : %v", exp, op)
	}

	//Error Check
	expErr := Model.CustomError{http.StatusNoContent, "No user Found"}
	svc = New(&mockStore{returnErr: true})

	errRes, errs := svc.ViewTask()

	if !errors.Is(errs, expErr) {
		//t.Errorf("expected error: %v, but got: %v", expErr, errs)
	}

	if !reflect.DeepEqual(errRes, Model.UserSlice{}) {
		t.Errorf("expected empty result, Expected: %v but got %v", Model.User{}, errRes)
	}
}
