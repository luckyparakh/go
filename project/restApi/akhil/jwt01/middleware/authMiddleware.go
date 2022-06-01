package middleware

import (
	"fmt"
	"jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprint("No auth header provided."),
			})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {

		}
	}
}
