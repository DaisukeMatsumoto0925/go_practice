package main

import (
	"fmt"
	"testing"

	"github.com/DaisukeMatsumoto0925/api/src/infra/rdb"
)

func Test_main(t *testing.T) {
	rdb := rdb.NewRDB()
	fmt.Println(rdb.Config)

	tx1 := rdb.Begin()
	t.Cleanup(func() {
		fmt.Println("Rollback!!")
		tx1.Rollback()
	})

	tx2 := rdb.Begin()
	if err := tx2.Commit().Error; err != nil {
		tx2.Rollback()
		fmt.Println(err)
	}

	fmt.Println("finish")

}
