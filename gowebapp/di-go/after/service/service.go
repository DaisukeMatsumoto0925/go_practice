package service

import "context"

type UserService interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) {
	us.ur.CreateUser(ctx, user)
}
