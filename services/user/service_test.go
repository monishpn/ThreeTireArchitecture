package user

import (
	"go.uber.org/mock/gomock"
	"testing"
)

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockStore := NewMockUserStore(ctrl)

	svc := New(mockStore)

	mockStore.EXPECT().AddUser("Ram").Return(nil)

	err := svc.AddUser("Ram")
	if err != nil {
		t.Errorf("Error in Testing Add_User, Expected : %v, got : %v", nil, err)
	}
}
