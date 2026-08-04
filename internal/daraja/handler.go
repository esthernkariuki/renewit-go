package daraja

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// STKPush godoc
//
//	@Summary		Initiate M-Pesa STK Push
//	@Description	Send an STK Push request to the customer's phone using Safaricom Daraja API
//	@Tags			Daraja
//	@Accept			json
//	@Produce		json
//	@Param			request	body		daraja.STKPushPayload	true	"STK Push request"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/daraja/stkpush [post]
func STKPush(c *gin.Context) {

	var payload STKPushPayload

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	service := NewDarajaService()

	response, err := service.STKPush(
		payload.PhoneNumber,
		payload.Amount,
		payload.AccountReference,
		payload.TransactionDesc,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
