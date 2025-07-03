package task

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

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskService(ctrl)

	svc := New(mockService)

	testCases := []struct {
		desc   string
		input  string
		ipErr  error
		opCode int
		opMsg  []byte
		ifMock bool
	}{
		{
			"Normal Adding",
			`{"task":"Do Homework", "userID":1}`,
			nil,
			http.StatusCreated,
			[]byte("Task added"),
			true,
		},
		{
			"JSON Error",
			`{"task":"Do Homework", "userID":"abc"}`,
			nil,
			http.StatusInternalServerError,
			[]byte("Internal Server Error"),
			false,
		},
		{
			"Service Error",
			`{"task":"Do Homework", "userID":1}`,
			models.CustomError{Code: http.StatusBadRequest, Message: "Invalid User"},
			http.StatusBadRequest,
			[]byte("Invalid User"),
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if tc.ifMock {
				mockService.EXPECT().AddTask("Do Homework", 1).Return(tc.ipErr)
			}

			req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(tc.input))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			svc.Addtask(w, req)

			if w.Code != tc.opCode {
				t.Errorf("Expected status %d, got %d", tc.opCode, w.Code)
			}

			if !bytes.Contains(w.Body.Bytes(), tc.opMsg) {
				t.Errorf("Expected response to contain %q, got %q", tc.opMsg, w.Body.String())
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskService(ctrl)
	svc := New(mockService)

	task := models.Tasks{Tid: 1, Task: "Read book", Completed: false, UserID: 2}

	// Success case
	mockService.EXPECT().GetByID(1).Return(task, nil)
	req := httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req.SetPathValue("id", "1")

	w := httptest.NewRecorder()
	svc.Gettask(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}

	if w.Body.String() != task.String() {
		t.Errorf("Expected %s, got %s", task.String(), w.Body.String())
	}

	// Custom error
	mockService.EXPECT().GetByID(1).Return(models.Tasks{}, models.CustomError{Code: http.StatusNotFound, Message: "Task not found"})

	req = httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req.SetPathValue("id", "1")

	w = httptest.NewRecorder()
	svc.Gettask(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}

	if w.Body.String() != "Task not found" {
		t.Errorf("Expected 'Task not found', got %s", w.Body.String())
	}

	// Invalid path param
	req = httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req.SetPathValue("id", "abc")

	w = httptest.NewRecorder()
	svc.Gettask(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestViewTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskService(ctrl)

	svc := New(mockService)

	tasks := []models.Tasks{
		{Tid: 1, Task: "A", Completed: false, UserID: 1},
		{Tid: 2, Task: "B", Completed: false, UserID: 2},
	}

	// Success case
	mockService.EXPECT().ViewTask().Return(tasks, nil)
	req := httptest.NewRequest(http.MethodGet, "/task", nil)

	w := httptest.NewRecorder()
	svc.Viewtask(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}

	expected, _ := json.Marshal(tasks)

	if !bytes.Equal(w.Body.Bytes(), expected) {
		t.Errorf("Expected %s, got %s", expected, w.Body.Bytes())
	}

	// Error case
	mockService.EXPECT().ViewTask().Return(nil, models.CustomError{Code: http.StatusInternalServerError, Message: "DB Error"})
	req = httptest.NewRequest(http.MethodGet, "/task", nil)

	w = httptest.NewRecorder()
	svc.Viewtask(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500, got %d", w.Code)
	}

	if w.Body.String() != "DB Error" {
		t.Errorf("Expected DB Error, got %s", w.Body.String())
	}
}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskService(ctrl)
	svc := New(mockService)

	mockService.EXPECT().UpdateTask(1).Return(true, nil)
	req := httptest.NewRequest(http.MethodPut, "/task/{id}", nil)

	req.SetPathValue("id", "1")

	w := httptest.NewRecorder()

	svc.Updatetask(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if w.Body.String() != "Task updated" {
		t.Errorf("Expected 'Task updated', got %s", w.Body.String())
	}

	// Error case
	mockService.EXPECT().UpdateTask(2).Return(false, models.CustomError{Code: 404, Message: "Task not found"})
	req = httptest.NewRequest(http.MethodPut, "/task/{id}", nil)

	req.SetPathValue("id", "2")

	w = httptest.NewRecorder()

	svc.Updatetask(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}

	if w.Body.String() != "Task not found" {
		t.Errorf("Expected 'Task not found', got %s", w.Body.String())
	}

	// Invalid ID
	req = httptest.NewRequest(http.MethodPut, "/task/{id}", nil)
	req.SetPathValue("id", "xyz")

	w = httptest.NewRecorder()
	svc.Updatetask(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskService(ctrl)
	svc := New(mockService)

	mockService.EXPECT().DeleteTask(1).Return(true, nil)
	req := httptest.NewRequest(http.MethodDelete, "/task/{id}", nil)
	req.SetPathValue("id", "1")

	w := httptest.NewRecorder()
	svc.Deletetask(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if w.Body.String() != "Task deleted" {
		t.Errorf("Expected 'Task deleted', got %s", w.Body.String())
	}

	mockService.EXPECT().DeleteTask(2).Return(false, models.CustomError{Code: 500, Message: "Internal error"})
	req = httptest.NewRequest(http.MethodDelete, "/task/{id}", nil)
	req.SetPathValue("id", "2")

	w = httptest.NewRecorder()
	svc.Deletetask(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500, got %d", w.Code)
	}

	if w.Body.String() != "Internal error" {
		t.Errorf("Expected 'Internal error', got %s", w.Body.String())
	}

	// Invalid ID
	req = httptest.NewRequest(http.MethodDelete, "/task/{id}", nil)
	req.SetPathValue("id", "bad")

	w = httptest.NewRecorder()
	svc.Deletetask(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}
