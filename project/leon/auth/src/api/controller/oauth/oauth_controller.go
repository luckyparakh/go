package oauth

import (
	"auth/src/api/domian/oauth"
	"auth/src/api/services"
	"auth/src/api/utils/apierror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var accessTokenReq *oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&accessTokenReq); err != nil {
		apiError := apierror.ApiError{
			Message: "bad request",
			Status:  http.StatusBadRequest,
			Error:   err,
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	accessToken, err := services.OauthService.CreateAccessToken(*accessTokenReq)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
