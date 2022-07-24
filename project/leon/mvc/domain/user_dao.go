package domain

import (
	"fmt"
	"log"
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "rishi", LastName: "parakh", Email: "my@email.com"},
	}
	UserDao UserDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type UserDaoInterface interface {
	GetUser(int64) (*User, *utils.AppError)
}
type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.AppError) {
	log.Println("Access DB")
	user := users[userId]
	if user == nil {
		return nil, &utils.AppError{
			Message:    fmt.Sprintf("user %v not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not found"}
	}
	return user, nil
}
