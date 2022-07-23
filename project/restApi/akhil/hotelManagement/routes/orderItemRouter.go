package routes

import (
	"hotelManagement/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/orderItems",controllers.GetOrderItems())
	incomingRouter.GET("/orderItems/:orderItem_id",controllers.GetOrderItem())
	incomingRouter.GET("/orderItems-order/:order_id",controllers.GetOrderItemsbyOrder())
	incomingRouter.POST("/orderItems",controllers.CreateOrderItem())
	incomingRouter.PATCH("/orderItems/:orderItem_id",controllers.UpdateOrderItem())
}
