package main

import (
	"fmt"
	"gowebapp/di-go/before/service"
)

func main() {

	us := service.NewUserService()
	fmt.Println(us)
}
