// @title Task API
// @version 1.0
// @description This is a sample task management API.
// @termsOfService http://swagger.io/terms/

// @contact.name Monish
// @contact.email you@example.com

// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"net/http"
	"time"

	_ "awesomeProject/docs"
	httpSwagger "github.com/swaggo/http-swagger"

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

	http.Handle("/swagger/", httpSwagger.WrapHandler)

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
