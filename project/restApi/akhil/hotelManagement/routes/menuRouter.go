package routes

import (
	"hotelManagement/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRouter *gin.Engine){
	incomingRouter.GET("/menus",controllers.GetMenus())
	incomingRouter.GET("/menus/:menu_id",controllers.GetMenu())
	incomingRouter.POST("/menus",controllers.CreateMenu())
	incomingRouter.PATCH("/menus/:menu_id",controllers.UpdateMenu())
}