package routes
import (
	controller "jwt/controllers"
	mw "jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(mw.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}