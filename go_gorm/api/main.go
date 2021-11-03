package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DaisukeMatsumoto0925/api/src/infra/rdb"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	rdb := rdb.NewRDB()
	fmt.Println(rdb.Config)

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
