package middleware

import (
	"jwt/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.JSON(401, gin.H{
				"error": "request does not contain token",
			})
			ctx.Abort()
			return
		}
		if err := auth.ValidateJWT(tokenStr); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
