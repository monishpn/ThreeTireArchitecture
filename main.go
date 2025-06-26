package main

import (
	"net/http"

	"awesomeProject/datasource"
	handler "awesomeProject/handler/task"
	store "awesomeProject/store/task"
)

func main() {
	db, err := datasource.New("")
	if err != nil {
		return
	}

	userStore := store.New(db)
	userSvc := userSvc.New(userStore)
	userHandler := userHndlr.New(userSvc)

	store := store.New(db)
	service := service.New(store, userSvc)
	handler := handler.New(service)

	http.HandleFunc("/task/create", handler.Create)

}
