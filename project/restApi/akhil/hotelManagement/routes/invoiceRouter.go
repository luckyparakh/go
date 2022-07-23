package routes

import (
	"hotelManagement/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRouter *gin.Engine){
	incomingRouter.GET("/invoice",controllers.GetInvoices())
	incomingRouter.GET("/invoice/:invoice_id",controllers.GetInvoice())
	incomingRouter.POST("/invoice",controllers.CreateInvoice())
	incomingRouter.PATCH("/invoice/:invoice_id",controllers.UpdateInvoice())
}