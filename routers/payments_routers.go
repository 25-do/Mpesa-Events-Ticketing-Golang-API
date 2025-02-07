package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.RouterGroup, paymentController *controllers.PaymentController) {
	paymentRoutes := router.Group("/payment")
	{
		paymentRoutes.GET("/", paymentController.GetAllPayments)
		paymentRoutes.POST("/create", paymentController.CreatePayment)
		paymentRoutes.GET("/getSingle", paymentController.GetSinglePayment)
		paymentRoutes.PUT("/update", paymentController.UpdatePayment)
		paymentRoutes.DELETE("/delete", paymentController.DeletePayment)
	}
}
