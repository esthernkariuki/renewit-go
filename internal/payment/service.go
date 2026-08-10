package payment

import (
	"errors"
	"fmt"
	"strings"

	"renewit-go/database"
	"renewit-go/internal/daraja"
	"renewit-go/internal/materials"
)

func GetAllPayment() []Payment {
	return FetchPayment()
}

func CreatePaymentService(payment *Payment) error {

	// =========================================
	// GET MATERIAL
	// =========================================

	var material materials.Material

	result := database.DB.First(
		&material,
		payment.MaterialID,
	)

	if result.Error != nil {
		return fmt.Errorf("material lookup failed: %w", result.Error)
	}

	fmt.Println("Material found:", material)

	// =========================================
	// VALIDATE QUANTITY
	// =========================================

	if payment.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}

	if payment.Quantity > material.Quantity {
		return fmt.Errorf(
			"only %d units are available",
			material.Quantity,
		)
	}

	// =========================================
	// CALCULATE PAYMENT
	// =========================================

	payment.Amount =
		material.Price * float64(payment.Quantity)

	// =========================================
	// GET TRADER FROM MATERIAL
	// =========================================

	payment.TraderID = material.TraderID

	// =========================================
	// PAYMENT DETAILS
	// =========================================

	payment.PaymentStatus = "PENDING"

	payment.AccountReference =
		fmt.Sprintf(
			"Material-%d",
			material.ID,
		)

	payment.TransactionDesc =
		fmt.Sprintf(
			"Purchase of %s",
			material.Type,
		)

	// =========================================
	// SAVE PAYMENT BEFORE STK PUSH
	// =========================================

	if err := SavePayment(payment); err != nil {
		return err
	}

	phoneNumber := strings.TrimSpace(payment.PhoneNumber)

	if strings.HasPrefix(phoneNumber, "0") {
		phoneNumber = "254" + phoneNumber[1:]
	}

	if strings.HasPrefix(phoneNumber, "+254") {
		phoneNumber = phoneNumber[1:]
	}

	payment.PhoneNumber = phoneNumber
	// =========================================
	// SEND M-PESA STK PUSH
	// =========================================

	service := daraja.NewDarajaService()

	response, err := service.STKPush(
		payment.PhoneNumber,
		int(payment.Amount),
		payment.AccountReference,
		payment.TransactionDesc,
	)

	if err != nil {
		return err
	}

	// =========================================
	// SAVE DARaja IDs
	// =========================================

	payment.MpesaCheckoutID =
		response.CheckoutRequestID

	payment.MpesaMerchantID =
		response.MerchantRequestID

	return UpdatePayment(payment)
}
