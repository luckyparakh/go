package utils

import "github.com/gin-gonic/gin"

func Respond(ctx *gin.Context, statusCode int, body any) {
	//localhost:8090/users/123 -v -H "Accept:application/xml"
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(statusCode, body)
		return
	}
	ctx.JSON(statusCode, body)
}

func RespondErr(ctx *gin.Context, body *AppError) {
	ctx.JSON(body.StatusCode, body)
}
