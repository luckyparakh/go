package domain

import (
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "rishi", LastName: "parakh", Email: "my@email.com"},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.AppError{
			Message:    "user not found",
			StatusCode: http.StatusNotFound,
			Code:       "not found"}
	}
	return user, nil
}
