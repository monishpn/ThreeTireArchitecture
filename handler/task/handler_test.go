package task

import (
	"awesomeProject/models"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	gofrHttp "gofr.dev/pkg/gofr/http"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestAddTask(t *testing.T) {
	type gofrResponse struct {
		result any
		err    error
	}

	mockContainer, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	tests := []struct {
		name             string
		task             string
		uID              int
		requestBody      string
		expectedResponse gofrResponse
		ifMock           bool
	}{
		{
			name:        "Successful add task",
			task:        "Testing",
			uID:         1,
			requestBody: `{"task":"Testing","userID":1}`,
			expectedResponse: gofrResponse{
				result: "Task added",
				err:    nil,
			},
			ifMock: true,
		},
		{
			name:        "Failed Binding",
			task:        "Testing",
			uID:         1,
			requestBody: `{"Testing","userID":1}`,
			expectedResponse: gofrResponse{
				result: nil,
				err:    gofrHttp.ErrorInvalidParam{Params: []string{"Give Correct Input"}},
			},
			ifMock: false,
		},
		{
			name:        "Check With Error",
			task:        "Testing",
			uID:         1,
			requestBody: `{"task":"Testing","userID":1}`,
			expectedResponse: gofrResponse{
				result: nil,
				err:    errors.New("testing error"),
			},
			ifMock: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			mockService := NewMockTaskService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			if tt.ifMock {
				mockService.EXPECT().AddTask(ctx, tt.task, tt.uID).Return(tt.expectedResponse.err)
			}

			val, err := svc.Addtask(ctx)

			response := gofrResponse{val, err}

			assert.Equal(t, tt.expectedResponse, response, "TEST[%d], Failed.\n%s", i, tt.name)
		})
	}
}

func TestViewTask(t *testing.T) {
	type gofrResponse struct {
		result any
		err    error
	}

	mockContainer, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	tests := []struct {
		name    string
		expResp gofrResponse
		ifMock  bool
	}{
		{
			name: "Successful view task",

			expResp: gofrResponse{
				result: []models.Tasks{
					{Tid: 1, Task: "A", Completed: false, UserID: 1},
					{Tid: 2, Task: "B", Completed: false, UserID: 2},
				},
				err: nil,
			},
			ifMock: true,
		},
		{
			name: "Error Testing",
			expResp: gofrResponse{
				result: nil,
				err:    errors.New("testing error"),
			},
			ifMock: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			mockService := NewMockTaskService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/task", http.NoBody)
			req.Header.Set("Content-Type", "application/json")

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			if tt.ifMock {
				mockService.EXPECT().ViewTask(ctx).Return(tt.expResp.result, tt.expResp.err)
			}

			val, err := svc.Viewtask(ctx)

			response := gofrResponse{val, err}

			assert.Equal(t, tt.expResp, response, "TEST[%d], Failed.\n%s", i, tt.name)
		})
	}
}

func TestGetByID(t *testing.T) {
	type gofrResponse struct {
		result any
		err    error
	}

	mockContainer, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	tests := []struct {
		name        string
		requestBody string
		expResp     gofrResponse
		ifMock      bool
	}{
		{
			name:        "Successful GetByID task",
			requestBody: "1",
			expResp: gofrResponse{
				result: models.Tasks{Tid: 1, Task: "A", Completed: false, UserID: 1},
				err:    nil,
			},
			ifMock: true,
		},
		{
			name:        "Error GetByID task",
			requestBody: "1",
			expResp: gofrResponse{
				result: models.Tasks{},
				err:    errors.New("testing error"),
			},
			ifMock: true,
		},
		{
			name:        "Testing strconv error",
			requestBody: "r",
			expResp: gofrResponse{
				result: nil,
				err:    gofrHttp.ErrorInvalidParam{Params: []string{"Invalid Param"}},
			},
			ifMock: false,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			mockService := NewMockTaskService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/task/{id}", http.NoBody)
			req.Header.Set("Content-Type", "application/json")
			req = mux.SetURLVars(req, map[string]string{ //working
				"id": tt.requestBody,
			})

			//req.SetPathValue("id", tt.requestBody) // not working

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			id, err := strconv.Atoi(tt.requestBody)

			if tt.ifMock {
				mockService.EXPECT().GetByID(ctx, id).Return(tt.expResp.result, tt.expResp.err)
			}

			val, err := svc.Gettask(ctx)

			response := gofrResponse{val, err}

			assert.Equal(t, tt.expResp, response, "TEST[%d], Failed.\n%s", i, tt.name)
		})
	}
}

func TestUpdate(t *testing.T) {
	type gofrResponse struct {
		result any
		err    error
	}

	mockContainer, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	tests := []struct {
		name        string
		requestBody string
		expResp     gofrResponse
		ifMock      bool
	}{
		{
			name:        "Successful Update task",
			requestBody: "1",
			expResp: gofrResponse{
				result: "Task updated",
				err:    nil,
			},
			ifMock: true,
		},
		{
			name:        "Error Update task",
			requestBody: "1",
			expResp: gofrResponse{
				result: nil,
				err:    errors.New("testing error"),
			},
			ifMock: true,
		},
		{
			name:        "Testing strconv error",
			requestBody: "r",
			expResp: gofrResponse{
				result: nil,
				err:    gofrHttp.ErrorInvalidParam{Params: []string{"Invalid Param"}},
			},
			ifMock: false,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			mockService := NewMockTaskService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/task/{id}", http.NoBody)
			req.Header.Set("Content-Type", "application/json")
			req = mux.SetURLVars(req, map[string]string{ //working
				"id": tt.requestBody,
			})

			//req.SetPathValue("id", tt.requestBody) // not working

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			id, err := strconv.Atoi(tt.requestBody)

			if tt.ifMock {
				mockService.EXPECT().UpdateTask(ctx, id).Return(true, tt.expResp.err)
			}

			val, err := svc.Updatetask(ctx)

			response := gofrResponse{val, err}

			assert.Equal(t, tt.expResp, response, "TEST[%d], Failed.\n%s", i, tt.name)
		})
	}
}

func TestDelete(t *testing.T) {
	type gofrResponse struct {
		result any
		err    error
	}

	mockContainer, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	tests := []struct {
		name        string
		requestBody string
		expResp     gofrResponse
		ifMock      bool
		ifResp      bool
	}{
		{
			name:        "Successful Delete task",
			requestBody: "1",
			expResp:     gofrResponse{},
			ifMock:      true,
			ifResp:      false,
		},
		{
			name:        "Error Delete task",
			requestBody: "1",
			expResp: gofrResponse{
				result: nil,
				err:    errors.New("testing error"),
			},
			ifMock: true,
			ifResp: true,
		},
		{
			name:        "Testing strconv error",
			requestBody: "r",
			expResp: gofrResponse{
				result: nil,
				err:    gofrHttp.ErrorInvalidParam{Params: []string{"Invalid Param"}},
			},
			ifMock: false,
			ifResp: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			mockService := NewMockTaskService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/task/{id}", http.NoBody)
			req.Header.Set("Content-Type", "application/json")
			req = mux.SetURLVars(req, map[string]string{ //working
				"id": tt.requestBody,
			})

			//req.SetPathValue("id", tt.requestBody) // not working

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			id, err := strconv.Atoi(tt.requestBody)

			if tt.ifMock {
				mockService.EXPECT().DeleteTask(ctx, id).Return(true, tt.expResp.err)
			}

			val, err := svc.Deletetask(ctx)

			response := gofrResponse{val, err}

			if tt.ifResp {
				assert.Equal(t, tt.expResp, response, "TEST[%d], Failed.\n%s", i, tt.name)
			}
		})
	}
}
