package service

import (
	"PlayTogether/model"
)

type UserServiceHandler struct {
	userRepo model.UserRepository
}

func NewUserService(userRepo model.UserRepository) model.UserService {
	return &UserServiceHandler{
		userRepo: userRepo,
	}
}

func (u *UserServiceHandler) GetByID(id string) (model.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserServiceHandler) CreateUser(user model.User) error {
	return u.userRepo.CreateUser(user)
}
