package oauth

import (
	"auth/src/api/utils/apierror"
	"fmt"
	"net/http"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() *apierror.ApiError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.User_Id)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessToken(atoken string) (*AccessToken, *apierror.ApiError) {
	token := tokens[atoken]
	if token == nil {
		return nil, &apierror.ApiError{
			Message: "token not found",
			Status:  http.StatusNotFound,
		}
	}
	return token, nil
}
