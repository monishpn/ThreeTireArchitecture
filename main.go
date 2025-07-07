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
	_ "awesomeProject/docs"
	Thandler "awesomeProject/handler/task"
	Uhandler "awesomeProject/handler/user"
	"awesomeProject/migrations/migrations"
	Tservice "awesomeProject/services/task"
	Uservice "awesomeProject/services/user"
	Tstore "awesomeProject/store/task"
	Ustore "awesomeProject/store/user"
	"gofr.dev/pkg/gofr"
)

func main() {
	//db, err := datasource.New("root:root123@tcp(localhost:3306)/test_db")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	app := gofr.New()
	app.Migrate(migrations.All())

	userStore := Ustore.New()
	userSvc := Uservice.New(userStore)
	userHandler := Uhandler.New(userSvc)

	taskStore := Tstore.New()
	taskSvc := Tservice.New(taskStore, userSvc)
	taskHandler := Thandler.New(taskSvc)

	app.GET("/health", func(ctx *gofr.Context) (any, error) {
		return map[string]string{"status": "UP"}, nil
	})
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
