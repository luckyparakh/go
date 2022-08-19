package oauth

import (
	"auth/src/api/utils/apierror"
	"net/http"
	"strings"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() *apierror.ApiError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return &apierror.ApiError{
			Message: "blank username",
			Status:  http.StatusBadRequest,
		}
	}
	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return &apierror.ApiError{
			Message: "blank password",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
