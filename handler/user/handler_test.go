package user

import (
	"go.uber.org/mock/gomock"
	"testing"
)

func TestAddUser(t *testing.T) {

	test_cases := []struct {
		desc  string
		input []byte

	}

	ctrl:=gomock.NewController(t)
	mockService:=NewMockUserService(ctrl)
	svc:=New(mockService)
}

func TestGetUserByID(t *testing.T) {}

func TestViewUser(t *testing.T) {}