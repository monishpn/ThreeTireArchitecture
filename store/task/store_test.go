package task

import (
	Models "awesomeProject/models"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	"net/http"
	"testing"
)

func TestViewTask(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id", "task", "completed", "uid"}).
		AddRow(1, "Testing-1", false, 1).
		AddRow(2, "Testing-2", false, 2)

	tests := []struct {
		name     string
		Tid      int
		mockfunc func()
		expAns   []Models.Tasks
		err      error
	}{
		{
			name: "Successful GetByID task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from TASKS").
					WillReturnRows(rows)

			},
			expAns: []Models.Tasks{
				Models.Tasks{Tid: 1, Task: "Testing-1", Completed: false, UserID: 1},
				Models.Tasks{Tid: 2, Task: "Testing-2", Completed: false, UserID: 2},
			},
			err: nil,
		},
		{
			name: "Failed add task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from Tasks").
					WithArgs(1).
					WillReturnRows(rows)
			},
			expAns: []Models.Tasks{},
			err:    Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans, err := svc.ViewTask(ctx)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}

		})
	}

}

func TestAddTask(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		task     string
		uid      int
		mockfunc func()
		err      error
	}{
		{
			name: "Successful add task",
			task: "Testing",
			uid:  1,
			mockfunc: func() {
				mock.SQL.ExpectExec("Insert into TASKS (task,completed,uid) values (?,?,?)").
					WithArgs("Testing", false, 1).
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Failed add task",
			task: "Testing",
			uid:  1,
			mockfunc: func() {
				mock.SQL.ExpectExec("Insert into TASKS (task,completed,uid) values (?,?)").
					WithArgs("Testing", false, 1).
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			err: Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Adding the Data to the Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			err := svc.AddTask(ctx, tt.task, tt.uid)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

		})
	}

}

func TestGetByIDTask(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id", "task", "completed", "uid"}).
		AddRow(1, "Testing", false, 1)

	tests := []struct {
		name     string
		Tid      int
		mockfunc func()
		expAns   Models.Tasks
		err      error
	}{
		{
			name: "Successful GetByID task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from TASKS where id=?").
					WithArgs(1).
					WillReturnRows(rows)

			},
			expAns: Models.Tasks{Tid: 1, Task: "Testing", UserID: 1},
			err:    nil,
		},
		{
			name: "Failed add task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from TASKS where id").
					WithArgs(1).
					WillReturnRows(rows)
			},
			expAns: Models.Tasks{},
			err:    Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While retrieving the Data from the Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans, err := svc.GetByID(ctx, tt.Tid)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}

		})
	}

}

func TestUpdateTask(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		Tid      int
		mockfunc func()
		expAns   bool
		err      error
	}{
		{
			name: "Successful Update task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectExec("UPDATE TASKS SET completed= true WHERE id=?").
					WithArgs(1).
					WillReturnResult(mock.SQL.NewResult(0, 1))

			},
			expAns: true,
			err:    nil,
		},
		{
			name: "Failed Update task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectExec("UPDATE TASKS SET completed= true WHERE id").
					WithArgs(1).
					WillReturnResult(mock.SQL.NewResult(0, 1))
			},
			expAns: false,
			err:    Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While Updating the database "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans, err := svc.UpdateTask(ctx, tt.Tid)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}

		})
	}

}

func TestDeleteTask(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		Tid      int
		mockfunc func()
		expAns   bool
		err      error
	}{
		{
			name: "Successful Update task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectExec("delete from TASKS where id=?").
					WithArgs(1).
					WillReturnResult(mock.SQL.NewResult(0, 1))

			},
			expAns: true,
			err:    nil,
		},
		{
			name: "Failed Update task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectExec("delete from TASKS where id").
					WithArgs(1).
					WillReturnResult(mock.SQL.NewResult(0, 1))
			},
			expAns: false,
			err:    Models.CustomError{Code: http.StatusInternalServerError, Message: "Error While deleting data in Database"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			var db *sql.DB
			svc := New(db)

			ans, err := svc.DeleteTask(ctx, tt.Tid)
			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v :  error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}

		})
	}

}

func TestCheckTask(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id"}).
		AddRow(1)

	tests := []struct {
		name     string
		Tid      int
		mockfunc func()
		expAns   bool
	}{
		{
			name: "Successful Checking task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select id from TASKS where id=?").
					WithArgs(1).
					WillReturnRows(rows)

			},
			expAns: true,
		},
		{
			name: "Empty row - Checking task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select id from TASKS where id=?").
					WithArgs(1).
					WillReturnRows(rows)

			},
			expAns: false,
		},
		{
			name: "Failed add task",
			Tid:  1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("select * from TASKS where id").
					WithArgs(1).
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

			ans := svc.CheckIfExists(ctx, tt.Tid)

			if !assert.Equal(t, tt.expAns, ans) {
				t.Errorf("%v :  \nExpected = %v\n got = %v", tt.name, tt.expAns, ans)
			}

		})
	}

}
