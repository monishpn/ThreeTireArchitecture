package task

import (
	models "awesomeProject/models"
	"errors"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"net/http"
	"reflect"
	"testing"
)

func TestAddTask(t *testing.T) {
	testCases := []struct {
		desc   string
		task   string
		uid    int
		ifUser bool
		expErr error
		ifMock bool
	}{
		{"Valid addition of task - If user exists", "Testing", 1, true, nil, true},
		{"Valid addition of task - If user doesn't exists", "Testing", 2, false, models.CustomError{http.StatusBadRequest, "No user found"}, true},

		{"Invalid addition of task", "", 3, false, models.CustomError{http.StatusBadRequest, "Task is Empty"}, false},
	}

	ctrl := gomock.NewController(t)
	mockTaskStore := NewMockTaskStore(ctrl)
	mockUserService := NewMockUserService(ctrl)
	svc := New(mockTaskStore, mockUserService)

	ctx := &gofr.Context{}

	for _, test := range testCases {
		if test.ifMock {
			mockUserService.EXPECT().CheckUserID(ctx, test.uid).Return(test.ifUser)

			if test.ifUser {
				mockTaskStore.EXPECT().AddTask(ctx, test.task, test.uid).Return(test.expErr)
			}
		}

		err := svc.AddTask(ctx, test.task, test.uid)

		if !errors.Is(test.expErr, err) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.expErr, err)
		}
	}
}

func TestView(t *testing.T) {
	exp := []models.Tasks{{
		1, "Testing-1", false, 1,
	}, {
		2, "testing-2", false, 1,
	},
	}

	ctrl := gomock.NewController(t)
	mockTaskStore := NewMockTaskStore(ctrl)
	mockUserService := NewMockUserService(ctrl)
	svc := New(mockTaskStore, mockUserService)

	ctx := &gofr.Context{}

	mockTaskStore.EXPECT().ViewTask(ctx).Return(exp, nil)
	op, err := svc.ViewTask(ctx)

	if err != nil {
		t.Errorf("%v.ViewTask(): expected no error, but got %v", t.Name(), err)
	}

	if !reflect.DeepEqual(op, exp) {
		t.Errorf("%v.ViewTask(): expected %v, but got %v", t.Name(), exp, op)
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		desc     string
		Tid      int
		ifExists bool
		exp      models.Tasks
		expErr   error
		ifMock   bool
	}{
		{"Valid retrieval of task - If it exists", 1, true, models.Tasks{1, "Testing", false, 1}, nil, true},
		{"Valid retrieval of task - If it does not exists", 5, false, models.Tasks{}, models.CustomError{http.StatusBadRequest, "No task found"}, false},
	}

	ctrl := gomock.NewController(t)
	mockTaskStore := NewMockTaskStore(ctrl)
	mockUserService := NewMockUserService(ctrl)
	svc := New(mockTaskStore, mockUserService)

	ctx := &gofr.Context{}

	for _, test := range testCases {
		mockTaskStore.EXPECT().CheckIfExists(ctx, test.Tid).Return(test.ifExists)

		if test.ifMock {
			mockTaskStore.EXPECT().GetByID(ctx, test.Tid).Return(test.exp, test.expErr)
		}

		op, err := svc.GetByID(ctx, test.Tid)

		if !errors.Is(test.expErr, err) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.expErr, err)
		}

		if !reflect.DeepEqual(test.exp, op) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.exp, op)
		}
	}
}

func TestUpdateTask(t *testing.T) {
	testCases := []struct {
		desc     string
		Tid      int
		ifExists bool
		exp      bool
		expErr   error
		ifMock   bool
	}{
		{"Valid Updating the task - If it exists", 1, true, true, nil, true},
		{"Valid Updating the task - If it does not exists", 5, false, false, models.CustomError{http.StatusBadRequest, "No task found"}, false},
	}

	ctrl := gomock.NewController(t)
	mockTaskStore := NewMockTaskStore(ctrl)
	mockUserService := NewMockUserService(ctrl)
	svc := New(mockTaskStore, mockUserService)

	ctx := &gofr.Context{}

	for _, test := range testCases {
		mockTaskStore.EXPECT().CheckIfExists(ctx, test.Tid).Return(test.ifExists)

		if test.ifMock {
			mockTaskStore.EXPECT().UpdateTask(ctx, test.Tid).Return(test.exp, test.expErr)
		}

		op, err := svc.UpdateTask(ctx, test.Tid)

		if !errors.Is(test.expErr, err) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.expErr, err)
		}

		if !reflect.DeepEqual(test.exp, op) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.exp, op)
		}
	}
}

func TestDeleteTask(t *testing.T) {
	testCases := []struct {
		desc     string
		Tid      int
		ifExists bool
		exp      bool
		expErr   error
		ifMock   bool
	}{
		{"Valid Deleting the task - If it exists", 1, true, true, nil, true},
		{"Valid Deleting the task - If it does not exists", 5, false, false, models.CustomError{http.StatusBadRequest, "No task found"}, false},
	}

	ctrl := gomock.NewController(t)
	mockTaskStore := NewMockTaskStore(ctrl)
	mockUserService := NewMockUserService(ctrl)
	svc := New(mockTaskStore, mockUserService)

	ctx := &gofr.Context{}

	for _, test := range testCases {
		mockTaskStore.EXPECT().CheckIfExists(ctx, test.Tid).Return(test.ifExists)

		if test.ifMock {
			mockTaskStore.EXPECT().DeleteTask(ctx, test.Tid).Return(test.exp, test.expErr)
		}

		op, err := svc.DeleteTask(ctx, test.Tid)

		if !errors.Is(test.expErr, err) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.expErr, err)
		}

		if !reflect.DeepEqual(test.exp, op) {
			t.Errorf("%v.AddTask(): expected %v, but got %v", test.desc, test.exp, op)
		}
	}
}
