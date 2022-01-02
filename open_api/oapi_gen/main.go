package main

import (
	"log"

	"github.com/DaisukeMatsumoto0925/oapi_gen/PetstoreExpanded"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	PetstoreExpanded.RegisterHandlersWithBaseURL(r, si, "api")
	log.Fatal(r.Start(":8080"))

}
