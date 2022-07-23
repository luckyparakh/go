package service

import (
	"mvc/domain"
	"mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userId)
}
