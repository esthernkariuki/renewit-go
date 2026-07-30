package payment

import (
	"renewit-go/internal/users"
	"time"
)

type Payment struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Type      string    `json:"type" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
	Condition string    `json:"condition" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	ListedAt  time.Time `json:"listed_at"`

	PhoneNumber      string  `json:"phone_number"`
	Amount           float64 `json:"amount"`
	AccountReference string  `json:"account_reference"`
	TransactionDesc  string  `json:"transaction_desc"`

	PaymentStatus      string    `json:"payment_status"`
	MpesaMerchantID    string    `json:"mpesa_merchant_id"`
	MpesaCheckoutID    string    `json:"mpesa_checkout_id"`
	MpesaReceiptNumber string    `json:"mpesa_receipt_number"`
	TransactionDate    time.Time `json:"transaction_date"`
	ResultCode         string    `json:"result_code"`
	ResultDescription  string    `json:"result_description"`

	PhoneNumberFromCallback string  `json:"phone_number_from_callback"`
	AmountFromCallback      float64 `json:"amount_from_callback"`

	TraderID uint       `json:"trader_id" binding:"required"`
	Trader   users.User `json:"trader" gorm:"foreignKey:TraderID"`

	UpcyclerID uint       `json:"upcycler_id" binding:"required"`
	Upcycler   users.User `json:"upcycler" gorm:"foreignKey:UpcyclerID"`
}
