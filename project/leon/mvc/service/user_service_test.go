package service

import (
	"mvc/domain"
	"mvc/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type userDaoMock struct{}

var (
	UserDaoMock domain.UserDaoInterface
	getUserFunc func(userId int64) (*domain.User, *utils.AppError)
)

func init() {
	domain.UserDao = &userDaoMock{}
}
func (u *userDaoMock) GetUser(userId int64) (*domain.User, *utils.AppError) {
	return getUserFunc(userId)
}
func TestGetUserNotFound(t *testing.T) {
	getUserFunc=func(userId int64) (*domain.User, *utils.AppError){
		return nil,&utils.AppError{
			Message: "user not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestGetUserFound(t *testing.T) {
	getUserFunc=func(userId int64) (*domain.User, *utils.AppError){
		return &domain.User{
			Id: 123,
		},nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
