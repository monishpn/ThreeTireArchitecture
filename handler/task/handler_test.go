package task

import (
	Models "awesomeProject/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockService struct {
	returnErr bool
}

func (m mockService) AddTask(task string, uid int) error {
	if m.returnErr {
		return Models.CustomError{Code: http.StatusBadRequest, Message: "AddTask Error"}
	}
	return nil
}

func (m mockService) ViewTask() ([]Models.Tasks, error) {
	if m.returnErr {
		return nil, Models.CustomError{Code: http.StatusNotFound, Message: "ViewTask Error"}
	}
	return []Models.Tasks{{Tid: 1, Task: "Demo", Completed: false, UserID: 1}}, nil
}

func (m mockService) GetByID(id int) (Models.Tasks, error) {
	if m.returnErr {
		return Models.Tasks{}, Models.CustomError{Code: http.StatusNotFound, Message: "GetByID Error"}
	}
	return Models.Tasks{Tid: 1, Task: "Demo", Completed: false, UserID: 1}, nil
}

func (m mockService) UpdateTask(id int) (bool, error) {
	if m.returnErr {
		return false, Models.CustomError{Code: http.StatusBadRequest, Message: "UpdateTask Error"}
	}
	return true, nil
}

func (m mockService) DeleteTask(id int) (bool, error) {
	if m.returnErr {
		return false, Models.CustomError{Code: http.StatusBadRequest, Message: "DeleteTask Error"}
	}
	return true, nil
}

func TestAddTask(t *testing.T) {
	hand := New(&mockService{})

	body := `{"task":"Read","userID":1}`
	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(body))
	w := httptest.NewRecorder()

	hand.Addtask(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %v, got %v", http.StatusCreated, w.Code)
	}

	//JSON error check
	body = `{"task","userID":1}`
	req = httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(body))
	w = httptest.NewRecorder()

	hand.Addtask(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected %v, got %v", http.StatusInternalServerError, w.Code)
	}

	// error check
	body = `{"task":"Read","userID":1}`
	hand = New(&mockService{returnErr: true})
	req = httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(body))
	w = httptest.NewRecorder()
	hand.Addtask(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, w.Code)
	}
}

func TestViewTask(t *testing.T) {
	hand := New(&mockService{})

	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	w := httptest.NewRecorder()

	hand.Viewtask(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected %v, got %v", http.StatusOK, w.Code)
	}

	// error check
	hand = New(&mockService{returnErr: true})
	req = httptest.NewRequest(http.MethodGet, "/task", nil)
	w = httptest.NewRecorder()
	hand.Viewtask(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected %v, got %v", http.StatusNotFound, w.Code)
	}
}

func TestGetTask(t *testing.T) {
	hand := New(&mockService{})

	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	hand.Gettask(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected %v, got %v", http.StatusOK, w.Code)
	}

	// strconv error
	req = httptest.NewRequest(http.MethodGet, "/task/abc", nil)
	req.SetPathValue("id", "abc")
	w = httptest.NewRecorder()
	hand.Gettask(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, w.Code)
	}

	// service error
	hand = New(&mockService{returnErr: true})
	req = httptest.NewRequest(http.MethodGet, "/task/1", nil)
	req.SetPathValue("id", "1")
	w = httptest.NewRecorder()
	hand.Gettask(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected %v, got %v", http.StatusNotFound, w.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	hand := New(&mockService{})
	req := httptest.NewRequest(http.MethodPut, "/task/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	hand.Updatetask(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected %v, got %v", http.StatusOK, w.Code)
	}

	// invalid id
	req = httptest.NewRequest(http.MethodPut, "/task/abc", nil)
	req.SetPathValue("id", "abc")
	w = httptest.NewRecorder()
	hand.Updatetask(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, w.Code)
	}

	// service error
	hand = New(&mockService{returnErr: true})
	req = httptest.NewRequest(http.MethodPut, "/task/1", nil)
	req.SetPathValue("id", "1")
	w = httptest.NewRecorder()
	hand.Updatetask(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, w.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	hand := New(&mockService{})
	req := httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	hand.Deletetask(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected %v, got %v", http.StatusOK, w.Code)
	}

	// invalid id
	req = httptest.NewRequest(http.MethodDelete, "/task/abc", nil)
	req.SetPathValue("id", "abc")
	w = httptest.NewRecorder()
	hand.Deletetask(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, w.Code)
	}

	// service error
	hand = New(&mockService{returnErr: true})
	req = httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	req.SetPathValue("id", "1")
	w = httptest.NewRecorder()
	hand.Deletetask(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, w.Code)
	}
}
