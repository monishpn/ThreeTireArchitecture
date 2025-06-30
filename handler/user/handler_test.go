package user

import (
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {

	test_cases := []struct {
		desc   string
		input  string
		ipErr  error
		opCode int
		opMsg  []byte

		ifMock bool
	}{
		{"Normal Adding", `{   "name":"Ram"  }`, nil, http.StatusCreated, []byte("User Created"), true},
		{"JSON error", "'{\"Ram\"}'", nil, http.StatusInternalServerError, []byte("Error Unmarshalling"), false},
		{"Error Check", `{   "name":"Ram"  }`, models.CustomError{http.StatusInternalServerError, "Checking error"}, http.StatusInternalServerError, []byte("Checking error"), true},
	}

	ctrl := gomock.NewController(t)
	mockService := NewMockUserService(ctrl)
	svc := New(mockService)

	for _, tc := range test_cases {
		if tc.ifMock {
			mockService.EXPECT().AddUser("Ram").Return(tc.ipErr)
		}
		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(tc.input))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		svc.AddUser(w, req)

		if w.Code != tc.opCode {
			t.Errorf("%s: Expected response code %d, got %d", tc.desc, tc.opCode, w.Code)
		}
		if w.Body.String() != string(tc.opMsg) {
			t.Errorf("%s: Expected message %s, got %s", tc.desc, tc.opMsg, w.Body.String())
		}

	}
}

func TestGetUserByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockService := NewMockUserService(ctrl)
	svc := New(mockService)

	inp := models.User{1, "Ram"}
	mockService.EXPECT().GetUserId(1).Return(inp, nil)
	req := httptest.NewRequest(http.MethodGet, "/user/{id}", nil)
	w := httptest.NewRecorder()
	req.SetPathValue("id", "1")
	svc.GetUserByID(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != (inp.String()) {
		t.Errorf("Expected message %s, got %s", inp.String(), w.Body.String())
	}

	//for error output
	mockService.EXPECT().GetUserId(1).Return(models.User{}, models.CustomError{http.StatusInternalServerError, "Checking error"})
	req = httptest.NewRequest(http.MethodGet, "/user/{id}", nil)
	w = httptest.NewRecorder()
	req.SetPathValue("id", "1")
	svc.GetUserByID(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "Checking error" {
		t.Errorf("Expected message Checking error, got %s", inp.String())
	}

	//Path value error
	req = httptest.NewRequest(http.MethodGet, "/user/{id}", nil)
	w = httptest.NewRecorder()
	req.SetPathValue("id", "r")
	svc.GetUserByID(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "Invalid ID" {
		t.Errorf("Expected message Invalid ID, got %s", w.Body.String())
	}

}

func TestViewUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockService := NewMockUserService(ctrl)
	svc := New(mockService)

	inp := models.UserSlice{{1, "Ram"}, {2, "Shyam"}}
	mockService.EXPECT().ViewTask().Return(inp, nil)
	req := httptest.NewRequest(http.MethodGet, "/user/{id}", nil)
	w := httptest.NewRecorder()
	svc.Viewuser(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, w.Code)
	}
	b, _ := json.Marshal(inp)
	if !bytes.Equal(b, w.Body.Bytes()) {
		t.Errorf("Expected message %v, got %v", b, w.Body.Bytes())
	}

	//for error output
	mockService.EXPECT().ViewTask().Return(inp, models.CustomError{http.StatusInternalServerError, "Checking error"})
	req = httptest.NewRequest(http.MethodGet, "/user/{id}", nil)
	w = httptest.NewRecorder()
	svc.Viewuser(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "Checking error" {
		t.Errorf("Expected message Checking error, got %s", w.Body.String())
	}

}
