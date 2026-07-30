package catalogue

import (
	"time"
)

type Catalogue struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	MaterialType        string    `json:"material_type"`
	PricePerKg          float64   `json:"price_per_kg"`
	LastUpdateDate      time.Time `json:"last_update_date"`
	MaterialDescription string    `json:"material_description"`
	CreatedAt           time.Time `json:"created_at"`
}
