package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing-system/models"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	PaymentService services.PaymentServiceInterface
}

func NewPaymentController(paymentService services.PaymentServiceInterface) *PaymentController {
	return &PaymentController{PaymentService: paymentService}
}
func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
	var venue models.Payment
	if err := c.ShouldBindJSON(&venue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.PaymentService.MpesaOnlinePayment(venue.Amount, venue.PhoneNumber, venue.OrganizerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newPayment, err := ctrl.PaymentService.CreatePayment(&venue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPayment)
}
func (uc *PaymentController) GetAllPayments(c *gin.Context) {
	payment, err := uc.PaymentService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
func (ctrl *PaymentController) GetSinglePayment(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := ctrl.PaymentService.GetSinglePayment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive payment"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (ctrl *PaymentController) UpdatePayment(c *gin.Context) {
	var ven models.Payment
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "an error occured check your Query parameters"})
		return
	}
	if err := c.ShouldBindJSON(&ven); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := ctrl.PaymentService.UpdatePayment(uint(id), ven)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while updating the payment"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (ctrl *PaymentController) DeletePayment(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := ctrl.PaymentService.DeletePayment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
