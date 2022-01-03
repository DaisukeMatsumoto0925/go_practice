package main

import (
	"log"

	openapi "github.com/DaisukeMatsumoto0925/oapi_gen/openApi"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/infrastructure/db"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	sqlhandler := db.NewSQLHandler()

	controller := controllers.NewPetController(sqlhandler)
	siw := &openapi.ServerInterfaceWrapper{
		Handler: controller,
	}
	openapi.RegisterHandlersWithBaseURL(r, siw.Handler, "api")
	log.Fatal(r.Start(":8080"))

}
