package main

import (
	infrastructure "clean_arc/infrastructure"
	"clean_arc/interfaces/controllers"
)

func main() {
	userController := controllers.NewUserController(infrastructure.NewSqlHandler())
	r := infrastructure.NewRouter(userController)
	r.Run()
}
