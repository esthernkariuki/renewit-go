package payment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {

	var callback CallbackRequest

	err := c.ShouldBindJSON(&callback)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var amount float64
	var receiptNumber string
	var phoneNumber string
	var transactionDate string

	for _, item := range callback.Body.STKCallback.CallbackMetadata.Item {

		switch item.Name {

		case "Amount":
			amount = item.Value.(float64)

		case "MpesaReceiptNumber":
			receiptNumber = item.Value.(string)

		case "PhoneNumber":
			phoneNumber = fmt.Sprintf("%.0f", item.Value.(float64))

		case "TransactionDate":
			transactionDate = fmt.Sprintf("%.0f", item.Value.(float64))
		}
	}

	paymentRecord, err := GetPaymentByCheckoutID(
		callback.Body.STKCallback.CheckoutRequestID,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Payment not found",
		})
		return
	}

	paymentRecord.PaymentStatus = "SUCCESS"
	paymentRecord.MpesaReceiptNumber = receiptNumber
	paymentRecord.PhoneNumberFromCallback = phoneNumber
	paymentRecord.AmountFromCallback = amount
	paymentRecord.ResultCode = fmt.Sprintf("%d", callback.Body.STKCallback.ResultCode)
	paymentRecord.ResultDescription = callback.Body.STKCallback.ResultDesc

	parsedTime, err := time.Parse("20060102150405", transactionDate)

	if err == nil {
		paymentRecord.TransactionDate = parsedTime
	}
	err = UpdatePayment(paymentRecord)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update payment",
		})
		return
	}

	fmt.Println("Amount:", amount)
	fmt.Println("Receipt Number:", receiptNumber)
	fmt.Println("Phone Number:", phoneNumber)
	fmt.Println("Transaction Date:", transactionDate)

	c.JSON(http.StatusOK, gin.H{
		"message": "Callback received successfully",
	})
}
