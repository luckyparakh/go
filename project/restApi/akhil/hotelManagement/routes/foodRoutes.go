package routes

import (
	"hotelManagement/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/foods",controllers.GetFoods())
	incomingRouter.GET("/foods/:food_id",controllers.GetFood())
	incomingRouter.POST("/foods",controllers.CreateFood())
	incomingRouter.PATCH("/foods/:food_id",controllers.UpdateFood())
}
