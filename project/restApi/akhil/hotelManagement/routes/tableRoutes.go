package routes

import (
	"hotelManagement/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/tables",controllers.GetTables())
	incomingRouter.GET("/tables/:table_id",controllers.GetTable())
	incomingRouter.POST("/tables",controllers.CreateTable())
	incomingRouter.PATCH("/tables/:table_id",controllers.UpdateTable())
}
