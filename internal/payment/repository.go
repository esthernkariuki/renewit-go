package payment

import "renewit-go/database"

func FetchPayment() []Payment {
	var payments []Payment

	database.DB.
		Preload("Material").
		Preload("Trader").
		Preload("Upcycler").
		Find(&payments)

	return payments
}

func SavePayment(payment *Payment) error {
	result := database.DB.Create(payment)
	return result.Error
}

func UpdatePayment(payment *Payment) error {
	return database.DB.Save(payment).Error
}

func GetPaymentByCheckoutID(checkoutID string) (*Payment, error) {
	var payment Payment

	err := database.DB.
		Where("mpesa_checkout_id = ?", checkoutID).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
