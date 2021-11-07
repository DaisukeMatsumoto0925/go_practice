package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/DaisukeMatsumoto0925/api/src/infra/rdb"
	"gorm.io/gorm"
)

func Test_main(t *testing.T) {
	rdb := rdb.NewRDB()
	rdb2 := newRdb()
	fmt.Println(rdb.Config)

	tx1 := rdb.Begin()
	t.Cleanup(func() {
		fmt.Println("Rollback!!")
		if err := tx1.Rollback().Error; err != nil {
			panic(err)
		}
	})

	ctx := context.Background()
	_, err := transaction(rdb2, ctx, func(ctx context.Context) (interface{}, error) {

		return nil, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// tx2 := rdb.Begin()
	// if err := tx2.Commit().Error; err != nil {
	// 	tx2.Rollback()
	// 	fmt.Println(err)
	// }

	fmt.Println("finish")

}

func newRdb() *gorm.DB {
	return rdb.NewRDB()
}
