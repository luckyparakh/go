package users

import (
	"auth/src/api/utils/apierror"
	"net/http"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id,username FROM users WHERE username=?"
)

var userData = map[string]*User{
	"fede": {
		Id:       1,
		Username: "fede",
	},
}

func GetUserByUsernameAndPassword(username, password string) (*User, *apierror.ApiError) {
	user := userData[username]
	if user.Username == "" {
		return nil, &apierror.ApiError{
			Message: "user not found",
			Status:  http.StatusNotFound,
		}
	}
	return user, nil
}
