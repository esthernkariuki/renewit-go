package payment

import (
	"github.com/gin-gonic/gin"
)

// GetPayment godoc
//
//	@Summary		Get all payments
//	@Description	Retrieve all payment records
//	@Tags			Payments
//	@Produce		json
//	@Success		200	{array}	payment.Payment
//	@Router			/payments [get]
func GetPayment(c *gin.Context) {
	payment := GetAllPayment()
	c.JSON(200, payment)
}

// CreatePayment godoc
//
//	@Summary		Create payment
//	@Description	Create a new payment record
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			request	body		payment.Payment	true	"Payment details"
//	@Success		201		{object}	payment.Payment
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/payments [post]
func CreatePayment(c *gin.Context) {
	var payment Payment
	err := c.ShouldBindJSON(&payment)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := CreatePaymentService(&payment); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to updated payment",
		})
		return
	}
	c.JSON(201, payment)

}
