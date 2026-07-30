package payment

import "renewit-go/internal/daraja"

func GetAllPayment() []Payment {
	return FetchPayment()
}

func CreatePaymentService(payment *Payment) error {

	payment.PaymentStatus = "PENDING"

	if err := SavePayment(payment); err != nil {
		return err
	}

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

	payment.MpesaCheckoutID = response.CheckoutRequestID
	payment.MpesaMerchantID = response.MerchantRequestID

	return UpdatePayment(payment)
}
