package app

import "auth/src/api/controller/oauth"

func mapping() {
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
