package main

import (
	"log"

	openapi "github.com/DaisukeMatsumoto0925/oapi_gen/openApi"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	// sqlhandler := db.NewSQLHandler()

	petController := controllers.NewPetController(nil)
	userController := controllers.NewUserController(nil)
	handler := &handler{
		*petController,
		*userController,
	}

	siw := &openapi.ServerInterfaceWrapper{
		Handler: handler,
	}
	openapi.RegisterHandlersWithBaseURL(r, siw.Handler, "api")
	log.Fatal(r.Start(":8080"))

}

type handler struct {
	controllers.PetController
	controllers.UserController
}
