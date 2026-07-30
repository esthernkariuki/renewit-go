package daraja

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
