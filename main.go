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
	"awesomeProject/datasource"
	_ "awesomeProject/docs"
	Thandler "awesomeProject/handler/task"
	"awesomeProject/migrations/migrations"
	Tservice "awesomeProject/services/task"
	Tstore "awesomeProject/store/task"
	"gofr.dev/pkg/gofr"
	"log"

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

	app := gofr.New()
	app.Migrate(migrations.All())

	userStore := Ustore.New(db)
	userSvc := Uservice.New(userStore)
	userHandler := Uhandler.New(userSvc)

	taskStore := Tstore.New(db)
	taskSvc := Tservice.New(taskStore, userSvc)
	taskHandler := Thandler.New(taskSvc)
	
	app.GET("/task", taskHandler.Viewtask)
	app.GET("/task/{id}", taskHandler.Gettask)
	app.POST("/task", taskHandler.Addtask)
	app.PUT("/task/{id}", taskHandler.Updatetask)
	app.DELETE("/task/{id}", taskHandler.Deletetask)

	app.GET("/user", userHandler.Viewuser)
	app.GET("/user/{id}", userHandler.GetUserByID)
	app.POST("/user", userHandler.AddUser)

	app.Run()
}
