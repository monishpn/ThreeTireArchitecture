package user

import (
	models "awesomeProject/models"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	"net/http"
	"testing"
)

func TestAddUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		user     string
		mockfunc func()
		err      error
	}{
		{
			name: "Successful add user",
			user: "Tester",
			mockfunc: func() {
				mock.SQL.ExpectExec("insert into USERS (name) values (?)").
					WithArgs("Tester").
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Failed add user",
			user: "Tester",
			mockfunc: func() {
				mock.SQL.ExpectExec("insert into Users (name) values (?)").
					WithArgs("Tester").
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			err: models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Adding the Data to the Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			err := svc.AddUser(ctx, tt.user)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}
		})
	}
}

func TestViewUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id", "user"}).
		AddRow(1, "Tester-1").
		AddRow(2, "Tester-2")

	tests := []struct {
		name     string
		mockfunc func()
		expAns   []models.User
		err      error
	}{
		{
			name: "Successful View user",
			mockfunc: func() {
				mock.SQL.ExpectQuery("Select * from USERS").
					WillReturnRows(rows)
			},
			expAns: []models.User{
				models.User{UserID: 1, Name: "Tester-1"},
				models.User{UserID: 2, Name: "Tester-2"},
			},
			err: nil,
		},
		{
			name: "Failed View user",
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from Tasks").
					WithArgs(1).
					WillReturnRows(rows)
			},
			expAns: []models.User{},
			err:    models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans, err := svc.ViewUser(ctx)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}
		})
	}
}

func TestGetByIDUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id", "user"}).
		AddRow(1, "Tester-1")

	tests := []struct {
		name     string
		UID      int
		mockfunc func()
		expAns   models.User
		err      error
	}{
		{
			name: "Successful GetByID user",
			UID:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from USERS where uid=?").
					WithArgs(1).
					WillReturnRows(rows)
			},
			expAns: models.User{UserID: 1, Name: "Tester-1"},
			err:    nil,
		},
		{
			name: "Failed add task",
			UID:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from USERS").
					WithArgs(1).
					WillReturnRows(rows)
			},
			expAns: models.User{},
			err:    models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans, err := svc.GetUserByID(ctx, tt.UID)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}
		})
	}
}

func TestCheckIDUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id"}).AddRow(1)
	empRow := mock.SQL.NewRows([]string{"id"})

	tests := []struct {
		name     string
		UID      int
		mockfunc func()
		expAns   bool
	}{
		{
			name: "Successful Check user",
			UID:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select uid from USERS where uid=?").
					WithArgs(1).WillReturnRows(rows)
			},
			expAns: true,
		},
		{
			name: "EmptyRowCheck",
			UID:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select uid from USERS where uid=?").
					WithArgs(1).WillReturnRows(rows)
			},
			expAns: false,
		},
		{
			name: "Failed Check task",
			UID:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select uid from USERS").
					WithArgs(1).WillReturnRows(rows)
			},
			expAns: false,
		}, {
			name: "Check Empty",
			UID:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select uid from USERS where uid=?").
					WithArgs(1).
					WillReturnRows(empRow)
			},
			expAns: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans := svc.CheckUserID(ctx, tt.UID)

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}
		})
	}
}

func TestCheckRowUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"COUNT"}).
		AddRow(1)

	tests := []struct {
		name     string
		mockfunc func()
		expAns   bool
	}{
		{
			name: "Successful CheckRow user",
			mockfunc: func() {
				mock.SQL.ExpectQuery("Select COUNT(*) from USERS").
					WillReturnRows(rows)
			},
			expAns: true,
		},
		{
			name: "Empty Row Check",
			mockfunc: func() {
				mock.SQL.ExpectQuery("Select COUNT(*) from USERS").
					WillReturnRows(rows)
			},
			expAns: false,
		},
		{
			name: "Failed Check Row user",
			mockfunc: func() {
				mock.SQL.ExpectQuery("Select COUNT(*) from Users").
					WillReturnRows(rows)
			},
			expAns: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans := svc.CheckIfRowsExists(ctx)

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}
		})
	}
}
