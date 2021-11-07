package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type key string

var (
	transactionKey key = "transaction"
)

func main() {
	fmt.Println("hello")
}

func transaction(
	db *gorm.DB, ctx context.Context,
	f func(context.Context) (interface{}, error),
) (interface{}, error) {
	tx := db.Begin()

	ctx = setTx(ctx, tx)

	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}

func setTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, transactionKey, tx)
}
