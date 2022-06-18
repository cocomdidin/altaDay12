package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app := echo.New()
	app.GET("/users", userController.GetAllUsers(db))
	app.POST("/users", userController.CreateUser(db))

	// app.GET("/users", controller.GetAllUsers(db))
	// app.POST("/users", controller.CreateUser(db))
	app.Start(":8080")
}
