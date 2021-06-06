package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:@/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)

	fmt.Println(us)
}
