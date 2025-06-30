package user

import (
	"awesomeProject/models"
	"errors"
	"go.uber.org/mock/gomock"
	"net/http"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	testCases := []struct {
		id       int
		desc     string
		inp      string
		exp      error
		mockCall bool
	}{
		{1, "Testing for Valid Input", "Ram", nil, true},
		{id: 2, desc: "Testing for Empty String", inp: "", exp: models.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Empty String given as input"},
			mockCall: false},
	}

	ctrl := gomock.NewController(t)
	mockStore := NewMockUserStore(ctrl)
	svc := New(mockStore)

	for _, test := range testCases {
		if test.mockCall {
			mockStore.EXPECT().AddUser(test.inp).Return(test.exp)
		}

		err := svc.AddUser(test.inp)

		if !errors.Is(err, test.exp) {
			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.exp, nil, err)
		}
	}
}

func TestViewTask(t *testing.T) {
	testCases := []struct {
		id       int
		desc     string
		ifRow    bool
		exp      models.UserSlice
		expErr   error
		mockCall bool
	}{
		{1, "Testing for Valid Input", true,
			models.UserSlice{
				{1, "Ram"},
				{2, "Shyam"},
			},
			nil, true,
		},
		{
			2, "Testing for no user", false,
			models.UserSlice{},
			models.CustomError{http.StatusNoContent, "No user Found"},
			false,
		},
	}

	ctrl := gomock.NewController(t)
	mockStore := NewMockUserStore(ctrl)
	svc := New(mockStore)

	for _, test := range testCases {
		mockStore.EXPECT().CheckIfRowsExists().Return(test.ifRow)

		if test.mockCall {
			mockStore.EXPECT().ViewUser().Return(test.exp, test.expErr)
		}

		op, err := svc.ViewTask()

		if !errors.Is(err, test.expErr) {
			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.desc, test.expErr, err)
		}

		if !reflect.DeepEqual(op, test.exp) {
			t.Errorf("Error in Testing: %v\n, Expected : %v\n, got : %v", test.exp, test.exp, op)
		}
	}
}

func TestGetUserId(t *testing.T) {
	testCases := []struct {
		id       int
		desc     string
		ifUser   bool
		input    int
		exp      models.User
		expErr   error
		mockCall bool
	}{
		{1, "Testing while user exists", true, 1, models.User{1, "Ram"}, nil, true},
		{2, "Testing while user doesn't exists", false, 5, models.User{}, models.CustomError{
			Code:    http.StatusNotFound,
			Message: "user does not exists",
		}, false},
	}

	ctrl := gomock.NewController(t)
	mockStore := NewMockUserStore(ctrl)
	svc := New(mockStore)

	for _, test := range testCases {
		mockStore.EXPECT().CheckUserID(test.input).Return(test.ifUser)

		if test.mockCall {
			mockStore.EXPECT().GetUserByID(test.input).Return(test.exp, test.expErr)
		}

		op, err := svc.GetUserId(test.input)

		if !errors.Is(err, test.expErr) {
			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.desc, test.expErr, err)
		}

		if !reflect.DeepEqual(op, test.exp) {
			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.desc, test.exp, op)
		}
	}
}
