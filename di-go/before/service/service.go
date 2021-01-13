package service

import "gowebapp/di-go/before/repository"

type UserService struct {
	ur repository.UserRepository
}

func NewUserService() *UserService {
	// リポジトリ作成
	ur := repository.NewUserRepository()
	// 作成したリポジトリを直接serviceに
	return &UserService{*ur}
}
