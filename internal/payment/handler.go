package payment

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePaymentRequest struct {
	MaterialID  uint   `json:"material_id" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,min=1"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// GetPayment godoc
//
//	@Summary		Get all payments
//	@Description	Retrieve all payment records
//	@Tags			Payments
//	@Produce		json
//	@Success		200	{array}	payment.Payment
//	@Router			/payments [get]
func GetPayment(c *gin.Context) {
	payments := GetAllPayment()

	c.JSON(http.StatusOK, payments)
}

// CreatePayment godoc
//
//	@Summary		Create payment
//	@Description	Create a payment for a material
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			request	body	CreatePaymentRequest	true	"Payment details"
//	@Success		201		{object}	Payment
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/payments [post]
func CreatePayment(c *gin.Context) {

	// =========================================
	// READ REQUEST BODY
	// =========================================

	var request CreatePaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================================
	// GET AUTHENTICATED USER FROM JWT
	// =========================================

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found in token",
		})
		return
	}

	// JWT numeric claims are normally decoded as float64
	userIDFloat, ok := userIDValue.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID in token",
		})
		return
	}

	upcyclerID := uint(userIDFloat)

	fmt.Println("Authenticated Upcycler ID:", upcyclerID)

	// =========================================
	// BUILD PAYMENT
	// =========================================

	payment := Payment{
		MaterialID:  request.MaterialID,
		Quantity:    request.Quantity,
		PhoneNumber: request.PhoneNumber,
		UpcyclerID:  upcyclerID,
	}

	// =========================================
	// CREATE PAYMENT + STK PUSH
	// =========================================

	if err := CreatePaymentService(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================================
	// RESPONSE
	// =========================================

	c.JSON(http.StatusCreated, payment)
}
