package repository

import "database/sql"

type UserRepository struct {
	db sql.DB
}

func NewUserRepository() *UserRepository {
	// DB作成
	db, _ := sql.Open("mysql", "di-go:di-go@/di-go")
	defer db.Close()
	// 作成したDBを直接repositoryに
	return &UserRepository{*db}
}
