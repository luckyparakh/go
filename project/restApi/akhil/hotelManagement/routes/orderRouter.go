package routes

import (
	"hotelManagement/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/orders",controllers.GetOrders())
	incomingRouter.GET("/orders/:order_id",controllers.GetOrder())
	incomingRouter.POST("/orders",controllers.CreateOrder())
	incomingRouter.PATCH("/orders/:order_id",controllers.UpdateOrder())
}
