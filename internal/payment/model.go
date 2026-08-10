package payment

import (
	"renewit-go/internal/materials"
	"renewit-go/internal/users"
	"time"
)

type Payment struct {
	ID                      uint               `json:"id" gorm:"primaryKey"`
	MaterialID              uint               `json:"material_id" binding:"required"`
	Material                materials.Material `json:"material" gorm:"foreignKey:MaterialID"`
	Quantity                int                `json:"quantity" binding:"required,min=1"`
	PhoneNumber             string             `json:"phone_number" binding:"required"`
	Amount                  float64            `json:"amount"`
	AccountReference        string             `json:"account_reference"`
	TransactionDesc         string             `json:"transaction_desc"`
	PaymentStatus           string             `json:"payment_status"`
	MpesaMerchantID         string             `json:"mpesa_merchant_id"`
	MpesaCheckoutID         string             `json:"mpesa_checkout_id"`
	MpesaReceiptNumber      string             `json:"mpesa_receipt_number"`
	TransactionDate         time.Time          `json:"transaction_date"`
	ResultCode              string             `json:"result_code"`
	ResultDescription       string             `json:"result_description"`
	PhoneNumberFromCallback string             `json:"phone_number_from_callback"`
	AmountFromCallback      float64            `json:"amount_from_callback"`
	TraderID                uint               `json:"trader_id"`
	Trader                  users.User         `json:"trader" gorm:"foreignKey:TraderID"`
	UpcyclerID              uint               `json:"upcycler_id"`
	Upcycler                users.User         `json:"upcycler" gorm:"foreignKey:UpcyclerID"`
	CreatedAt               time.Time          `json:"created_at"`
	UpdatedAt               time.Time          `json:"updated_at"`
}
