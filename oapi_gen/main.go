package main

import (
	"log"

	"github.com/DaisukeMatsumoto0925/oapi_gen/openapi"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	si := &openapi.Server{}
	openapi.RegisterHandlersWithBaseURL(r, si, "api")
	log.Fatal(r.Start(":8080"))

}
