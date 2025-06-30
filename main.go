package main

import (
	"log"
	"net/http"
	"time"

	"awesomeProject/datasource"
	Thandler "awesomeProject/handler/task"
	Tservice "awesomeProject/services/task"
	Tstore "awesomeProject/store/task"

	Uhandler "awesomeProject/handler/user"
	Uservice "awesomeProject/services/user"
	Ustore "awesomeProject/store/user"
)

func main() {
	db, err := datasource.New("root:root123@tcp(localhost:3306)/test_db")
	if err != nil {
		log.Println(err)
		return
	}

	userStore := Ustore.New(db)
	userSvc := Uservice.New(userStore)
	userHandler := Uhandler.New(userSvc)

	taskStore := Tstore.New(db)
	taskSvc := Tservice.New(taskStore, userSvc)
	taskHandler := Thandler.New(taskSvc)

	http.HandleFunc("GET /task", taskHandler.Viewtask)
	http.HandleFunc("GET /task/{id}", taskHandler.Gettask)
	http.HandleFunc("POST /task", taskHandler.Addtask)
	http.HandleFunc("PUT /task/{id}", taskHandler.Updatetask)
	http.HandleFunc("DELETE /task/{id}", taskHandler.Deletetask)

	http.HandleFunc("GET /user", userHandler.Viewuser)
	http.HandleFunc("GET /user/{id}", userHandler.GetUserByID)
	http.HandleFunc("POST /user", userHandler.AddUser)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
