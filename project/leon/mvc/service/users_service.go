package service

import (
	"mvc/domain"
	"mvc/utils"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.UserDao.GetUser(userId)
}
