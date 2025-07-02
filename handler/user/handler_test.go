package user

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

func TestAddUser(t *testing.T) {
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
		user             string
		requestBody      string
		expectedResponse gofrResponse
		ifMock           bool
	}{
		{
			name:        "Successful add User",
			user:        "Tester",
			requestBody: `{"name":"Tester"}`,
			expectedResponse: gofrResponse{
				result: "User Created",
				err:    nil,
			},
			ifMock: true,
		},
		{
			name:        "Failed Binding",
			user:        "Testing",
			requestBody: `{"Tester"}`,
			expectedResponse: gofrResponse{
				result: nil,
				err:    gofrHttp.ErrorInvalidParam{Params: []string{"Give Correct Input"}},
			},
			ifMock: false,
		},
		{
			name:        "Check With Error",
			user:        "Tester",
			requestBody: `{"name":"Tester"}`,
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
			mockService := NewMockUserService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			if tt.ifMock {
				mockService.EXPECT().AddUser(ctx, tt.user).Return(tt.expectedResponse.err)
			}

			val, err := svc.AddUser(ctx)

			response := gofrResponse{val, err}

			assert.Equal(t, tt.expectedResponse, response, "TEST[%d], Failed.\n%s", i, tt.name)
		})
	}
}

func TestViewUser(t *testing.T) {
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
				result: models.UserSlice{
					models.User{UserID: 1, Name: "Tester-1"},
					models.User{UserID: 2, Name: "Tester-2"},
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
			mockService := NewMockUserService(ctrl)
			svc := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/user", http.NoBody)
			req.Header.Set("Content-Type", "application/json")

			request := gofrHttp.NewRequest(req)

			ctx.Request = request

			if tt.ifMock {
				mockService.EXPECT().ViewTask(ctx).Return(tt.expResp.result, tt.expResp.err)
			}

			val, err := svc.Viewuser(ctx)

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
				result: models.User{UserID: 1, Name: "Tester-1"},
				err:    nil,
			},
			ifMock: true,
		},
		{
			name:        "Error GetByID task",
			requestBody: "1",
			expResp: gofrResponse{
				result: models.User{},
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
			mockService := NewMockUserService(ctrl)
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
				mockService.EXPECT().GetUserId(ctx, id).Return(tt.expResp.result, tt.expResp.err)
			}

			val, err := svc.GetUserByID(ctx)

			response := gofrResponse{val, err}

			assert.Equal(t, tt.expResp, response, "TEST[%d], Failed.\n%s", i, tt.name)
		})
	}
}
