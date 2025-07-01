package models

import "fmt"

type User struct {
	UserID int
	Name   string
}

type Input struct {
	T string `json:"name"`
}

func (usr User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", usr.UserID, usr.Name)
}

type UserSlice []User
