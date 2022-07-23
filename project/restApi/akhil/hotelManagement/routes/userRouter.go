package routes

import (
	"github.com/gin-gonic/gin"
	"hotelManagement/controllers"
)

func UserRoutes(incomingRouter *gin.Engine){
	incomingRouter.GET("/users",controllers.GetUsers())
	incomingRouter.GET("/users/:user_id",controllers.GetUsers())
	incomingRouter.POST("/users/signup",controllers.SignUp())
	incomingRouter.POST("/users/login",controllers.Login())
}