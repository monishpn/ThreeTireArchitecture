package user

import (
	Model "awesomeProject/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockService struct {
	returnErr bool
}

func (m mockService) AddUser(name string) error {

	if !m.returnErr {
		return nil
	}
	return Model.CustomError{http.StatusOK, "Checking Handler of AddUser"}
}

func (m mockService) ViewTask() (Model.UserSlice, error) {

	if m.returnErr {
		return nil, Model.CustomError{http.StatusOK, "Checking Handler of ViewUser"}

	}

	return Model.UserSlice{{1, "Ram"}}, nil
}

func (m mockService) GetUserId(id int) (Model.User, error) {

	if m.returnErr {
		return Model.User{}, Model.CustomError{http.StatusOK, "Checking Handler of GetUserId"}
	}
	return Model.User{1, "Ram"}, nil

}

func TestAdd(t *testing.T) {
	hand := New(&mockService{})
	user := `{
		"user":"Ram"
	}`
	req := httptest.NewRequest("GET", "/user", strings.NewReader(user))
	w := httptest.NewRecorder()

	hand.AddUser(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("AddUser Test Failed, expected: %v, got : %v", http.StatusCreated, w.Code)
	}

	//JSON error check
	user = `{
		"Ram"
	}`
	req = httptest.NewRequest("GET", "/user", strings.NewReader(user))
	w = httptest.NewRecorder()

	hand.AddUser(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("AddUser Test Failed, expected: %v, got : %v", http.StatusInternalServerError, w.Code)
	}

	//Error Check
	hand = New(&mockService{true})
	user = `{
		"user":"Ram"
	}`
	exp := Model.CustomError{http.StatusOK, "Checking Handler of AddUser"}

	req = httptest.NewRequest("GET", "/user", strings.NewReader(user))
	w = httptest.NewRecorder()

	hand.AddUser(w, req)

	if w.Code != exp.Code {
		t.Errorf("AddUser Test Failed, expected: %v, got : %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != exp.Message {
		t.Errorf("AddUser Test Failed, expected: %v, got : %v", exp.Message, w.Body.String())
	}

}

func TestGetByID(t *testing.T) {
	hand := New(&mockService{})

	req := httptest.NewRequest("GET", "/user/{id}", nil)
	w := httptest.NewRecorder()

	req.SetPathValue("id", "1")

	hand.GetUserByID(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != (Model.User{1, "Ram"}).String() {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", (Model.User{1, "Ram"}).String(), w.Body.String())
	}

	//Error Check
	hand = New(&mockService{true})
	exp := Model.CustomError{http.StatusOK, "Checking Handler of GetUserId"}
	req = httptest.NewRequest("GET", "/user/{id}", nil)
	w = httptest.NewRecorder()
	req.SetPathValue("id", "1")
	hand.GetUserByID(w, req)
	if w.Code != exp.Code {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != exp.Message {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", exp.Message, w.Body.String())
	}

	//strconv Error Check
	hand = New(&mockService{})

	req = httptest.NewRequest("GET", "/user/{id}", nil)
	w = httptest.NewRecorder()

	req.SetPathValue("id", "r")

	hand.GetUserByID(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", http.StatusOK, w.Code)
	}
}

func TestView(t *testing.T) {
	hand := New(&mockService{})

	req := httptest.NewRequest("GET", "/user/{id}", nil)
	w := httptest.NewRecorder()

	hand.Viewuser(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", http.StatusOK, w.Code)
	}

	//Error Check
	hand = New(&mockService{true})
	exp := Model.CustomError{http.StatusOK, "Checking Handler of ViewUser"}
	req = httptest.NewRequest("GET", "/user/{id}", nil)
	w = httptest.NewRecorder()
	req.SetPathValue("id", "1")
	hand.Viewuser(w, req)
	if w.Code != exp.Code {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != exp.Message {
		t.Errorf("GetUserByID Test Failed, expected: %v, got : %v", exp.Message, w.Body.String())
	}
}
