package services

import (
	"auth/src/api/domian/oauth"
	"auth/src/api/domian/users"
	"auth/src/api/utils/apierror"
	"net/http"
	"time"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, *apierror.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, *apierror.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (os *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, *apierror.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, &apierror.ApiError{
			Status:  http.StatusBadRequest,
			Message: "invalid username or password",
		}
	}
	user, err := users.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	at := &oauth.AccessToken{
		User_Id: user.Id,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := at.Save(); err != nil {
		return nil, err
	}
	return at, nil
}

func (os *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, *apierror.ApiError) {
	token, err := oauth.GetAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, &apierror.ApiError{
			Message: "token exipred",
			Status:  http.StatusNotFound,
		}
	}
	return token, nil
}
