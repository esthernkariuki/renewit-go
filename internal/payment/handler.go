package payment

import (
	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	payment := GetAllPayment()
	c.JSON(200, payment)
}

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
